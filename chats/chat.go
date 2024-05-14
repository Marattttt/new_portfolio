package chats

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"sync"

	"github.com/Marattttt/newportfolio/models"
	"github.com/docker/distribution/uuid"
	"nhooyr.io/websocket"
)

type IncomingMessage struct {
	Text string  `json:"text"`
	To   *string `json:"to"`
}

func (im IncomingMessage) ToMsg(from string) models.Message {
	return models.Message{
		From: from,
		Text: im.Text,
		To:   im.To,
	}
}

type Client struct {
	// Client id in a room
	id uuid.UUID

	// The room the client is attached to
	room *Room

	// Client name
	name string

	// The coonnection the client is attached to
	conn *websocket.Conn

	close chan struct{}

	sendMu sync.Mutex
}

func NewClient(req models.RoomJoinRequest, room *Room, conn *websocket.Conn) Client {
	return Client{
		id:   uuid.Generate(),
		room: room,
		name: req.Name,
		conn: conn,
	}
}

func (c *Client) HandleSend(ctx context.Context) error {
	if err := c.conn.Write(ctx, websocket.MessageText, []byte("Being handlesend")); err != nil {
		c.room.logger.Warn("Could not send message", slog.String("err", err.Error()))
	}

	for {
		select {
		case <-c.close:
			c.room.logger.Debug("Close client", slog.Int("eheh", 1))
			if err := c.conn.CloseNow(); err != nil {
				return fmt.Errorf("closing conn: %w", err)
			}
			return nil

		case err := <-c.receiveChan(ctx):
			// Unexpected error
			if err != nil && !errors.Is(err, io.EOF) {
				return fmt.Errorf("waiting for a new message: %w", err)
			}

			// Simple error
			if err != nil {
				c.room.logger.Info("stopped receiving messages", slog.String("clientId", c.id.String()))
				return nil
			}
		}
	}
}

// Wraps the receive message method method in a channel
func (c *Client) receiveChan(ctx context.Context) chan error {
	done := make(chan error)
	go func() {
		done <- c.receive(ctx)
	}()
	return done
}

// Waits for a new message and passes it to the assigned sender
func (c *Client) receive(ctx context.Context) error {
	msgType, msgReader, err := c.conn.Reader(ctx)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return fmt.Errorf("reading new message: %w", err)
		}
		return nil
	}

	if msgType == websocket.MessageBinary {
		c.conn.Write(ctx, websocket.MessageText, []byte("Binary not supported"))
	}

	var msg IncomingMessage
	if err := json.NewDecoder(msgReader).Decode(&msg); err != nil {
		c.conn.Write(ctx, websocket.MessageText, []byte("Could not unmarshall message"))
	}

	c.room.Sender.Send(ctx, c, msg.ToMsg(c.name))

	return nil
}

func (c *Client) Receive(ctx context.Context, msg models.Message) error {
	c.sendMu.Lock()
	defer c.sendMu.Unlock()

	wswriter, err := c.conn.Writer(ctx, websocket.MessageText)
	if err != nil {
		return fmt.Errorf("creating a writer: %w", err)
	}
	defer wswriter.Close()

	if err := json.NewEncoder(wswriter).Encode(msg); err != nil {
		return fmt.Errorf("encoding a message: %w", err)
	}

	return nil
}
