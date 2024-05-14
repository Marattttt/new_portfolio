package chats

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Marattttt/newportfolio/contracts"
	"github.com/Marattttt/newportfolio/models"
	"github.com/docker/distribution/uuid"
	"nhooyr.io/websocket"
)

type Hub struct {
	logger      *slog.Logger
	Rooms       map[int]Room
	BaseContext func() context.Context
}

func (h *Hub) Close() error {
	for _, r := range h.Rooms {
		fmt.Println("closing roomm", r.Id)
		r.Close()
		fmt.Println("closedroom", r.Id)
	}
	fmt.Println("doe closeeing roooms")
	return nil
}

func NewHub() Hub {
	return Hub{
		Rooms: make(map[int]Room),
	}
}

func (h *Hub) Join(ctx context.Context, req models.RoomJoinRequest, conn *websocket.Conn, logger *slog.Logger) {
	_, ok := h.Rooms[req.RoomId]
	if !ok {
		if h.Rooms == nil {
			h.Rooms = make(map[int]Room)
		}

		room := Room{
			Id:      req.RoomId,
			hub:     h,
			logger:  slog.With(logger, slog.Int("roomId", req.RoomId)),
			Clients: make(map[uuid.UUID]*Client),
		}
		room.Sender = SimpleSender{
			logger: room.logger,
			room:   &room,
		}

		logger.Debug("Created new room", slog.Int("roomId", room.Id))
		h.Rooms[req.RoomId] = room
	}

	r := h.Rooms[req.RoomId]

	r.Join(ctx, req, conn)
}

type Room struct {
	Id      int
	hub     *Hub
	logger  *slog.Logger
	Clients map[uuid.UUID]*Client
	Sender  MessageSender
}

func (r *Room) Close() error {
	for _, c := range r.Clients {
		fmt.Println("closign clinet", c.name)
		c.close <- struct{}{}
		fmt.Println("done closign clinet", c.name)
	}
	fmt.Println("done closing room ", r.Id)
	return nil
}

func NewRoom(hub *Hub, sender MessageSender) Room {
	return Room{
		Id:      len(hub.Rooms),
		hub:     hub,
		logger:  slog.With(hub.logger, slog.Int("roomId", len(hub.Rooms))),
		Clients: make(map[uuid.UUID]*Client),
		Sender:  sender,
	}
}

func (r *Room) Join(ctx context.Context, req models.RoomJoinRequest, conn *websocket.Conn) {
	if r.Clients == nil {
		r.logger.Debug("Creating new clients slice", slog.Int("roomId", r.Id))
		r.Clients = make(map[uuid.UUID]*Client)
	}

	client := NewClient(req, r, conn)
	r.Clients[client.id] = &client

	r.logger.Info("Begin receiving messages from client", slog.Any("clientName", client.name))
	if err := client.HandleSend(ctx); err != nil {
		r.logger.Warn("Unexpected error from chat client", slog.String("err", err.Error()))
	}
}

type MessageSender interface {
	Send(ctx context.Context, from *Client, msg models.Message)
}

type EchoSender struct {
	room   *Room
	logger *slog.Logger
}

func (es EchoSender) Send(ctx context.Context, from *Client, msg models.Message) {
	es.room.logger.Debug("New message", slog.Any("msg", msg))
	for _, c := range es.room.Clients {
		c.Receive(ctx, msg)
	}
}

type SimpleSender struct {
	room   *Room
	logger *slog.Logger
}

func (ss SimpleSender) Send(ctx context.Context, from *Client, msg models.Message) {
	ss.room.logger.Debug("New message", slog.Any("msg", msg))
	for _, c := range ss.room.Clients {
		if msg.To != nil {
			if c.name != *msg.To {
				continue
			}
		}
		if msg.From == c.name {
			continue
		}

		c.Receive(ctx, msg)
	}
}

var GlobalHub Hub

func HandleRoomJoin(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	ctx := r.Context()

	wsconn, err := websocket.Accept(w, r, nil)
	if err != nil {
		logger.Debug("Could not accespt websocket connection", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(contracts.ErrorResponse{"Something went wrong"})
		return
	}

	if err := wsconn.Write(ctx, websocket.MessageText, []byte("start connection")); err != nil {
		logger.Error("ehe", slog.String("err", err.Error()))
	}

	defer func() {
		logger.Info("Closing ws connection")

		if err := wsconn.Write(ctx, websocket.MessageText, []byte("closing connection")); err != nil {
			logger.Error("ehe", slog.String("err", err.Error()))
		}
		wsconn.CloseNow()
	}()

	joinReq := ReceiveJoinRequest(ctx, wsconn, logger)

	if joinReq == nil {
		wsconn.CloseNow()
		return
	}

	go GlobalHub.Join(r.Context(), *joinReq, wsconn, logger)
}

func ReceiveJoinRequest(ctx context.Context, conn *websocket.Conn, logger *slog.Logger) *models.RoomJoinRequest {
	msgType, reader, err := conn.Reader(ctx)
	if err != nil {
		logger.Warn("Error creating message reader for ws connction", slog.String("err", err.Error()))
		return nil
	}

	if msgType == websocket.MessageBinary {
		logger.Debug("Unsupported message type", slog.String("msgType", msgType.String()))
		return nil
	}

	if err := conn.Write(ctx, websocket.MessageText, []byte("please, send a message")); err != nil {
		logger.Error("ehe", slog.String("err", err.Error()))
	}

	var joinReq models.RoomJoinRequest
	if err := json.NewDecoder(reader).Decode(&joinReq); err != nil {
		slog.Debug("could not decode join request")
		return nil
	}

	logger.Info("Parsed join request", slog.Any("joinReq", joinReq))

	if err := conn.Write(ctx, websocket.MessageText, []byte("parsed and welcome")); err != nil {
		logger.Error("ehe", slog.String("err", err.Error()))
	}

	return &joinReq
}

func HandleEcho(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	wsconn, err := websocket.Accept(w, r, nil)
	if err != nil {
		logger.Error("Accepting websocket connection", slog.String("err", err.Error()))
	}
	defer wsconn.CloseNow()

	logger.Info("Opened ws connection")

	for {
		msgType, msg, err := wsconn.Read(context.TODO())
		if err != nil {
			if !errors.Is(err, io.EOF) {
				logger.Error("Reading message", slog.String("err", err.Error()))
			} else {
				logger.Info("Released ws connection")
			}
			break
		}

		wsconn.Write(context.TODO(), msgType, msg)
	}
}
