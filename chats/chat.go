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
	"nhooyr.io/websocket"
)

type Message struct {
	FromId int     `json:"from"`
	From   *Client `json:"-"`
	Toid   int     `json:"to"`
	To     *Client `json:"-"`
	Text   *string `json:"text,omitempty"`
}

type Client struct {
	// Client id in a room
	id int

	// The room the client is attached to
	room *Room

	// Client name
	name string

	// The coonnection the client is attached to
	conn *websocket.Conn

	sendMu *sync.Mutex
}

func NewClient(req models.RoomJoinRequest, room *Room, conn *websocket.Conn) Client {
	return Client{
		id:   len(room.Clients),
		room: room,
		name: req.Name,
		conn: conn,
	}
}

func (c *Client) HandleSend(ctx context.Context) error {
	if err := c.conn.Write(ctx, websocket.MessageText, []byte("connected!")); err != nil {
		c.room.logger.Warn("Could not send message", slog.String("err", err.Error()))
	}

	for {
		msgType, msgReader, err := c.conn.Reader(ctx)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return fmt.Errorf("reading new message: %w", err)
			}
			break
		}

		if msgType == websocket.MessageBinary {
			c.conn.Write(ctx, websocket.MessageText, []byte("Binary not supported"))
		}

		var msg Message
		if err := json.NewDecoder(msgReader).Decode(&msg); err != nil {
			c.conn.Write(ctx, websocket.MessageText, []byte("Could not unmarshall message"))
		}

		c.room.Sender.Send(ctx, c, &msg)
	}

	return nil
}

func (c *Client) Receive(ctx context.Context, msg *Message) error {
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
