package types

import (
	"encoding/json"
)

// SubscriptionEvent is an enum signifies whether the incoming message is of type Subscribe or unsubscribe
type SubscriptionEvent string

// Enum members for SubscriptionEvent
const (
	SUBSCRIBE   SubscriptionEvent = "subscribe"
	UNSUBSCRIBE SubscriptionEvent = "unsubscribe"
	Fetch       SubscriptionEvent = "fetch"
)

const MarketChannel = "markets"

type WebSocketMessage struct {
	Channel string           `json:"channel"`
	Payload WebSocketPayload `json:"payload"`
}

type WebSocketPayload struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type WebSocketSubscription struct {
	Event  SubscriptionEvent `json:"event"`
	Market Market            `json:"pair"`
	Params `json:"params"`
}

// Params is a sub document used to pass parameters in Subscription messages
type Params struct {
	From     int64  `json:"from"`
	To       int64  `json:"to"`
	Duration int64  `json:"duration"`
	Units    string `json:"units"`
	TickID   string `json:"tickID"`
}

func NewMarketWebsocketMessage(m *Market) *WebSocketMessage {
	return &WebSocketMessage{
		Channel: "markets",
		Payload: WebSocketPayload{
			Type: "NEW_MARKET",
			Data: m,
		},
	}
}

func (w *WebSocketMessage) Print() {
	b, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		logger.Error(err)
	}

	logger.Info(string(b))
}

func (w *WebSocketPayload) Print() {
	b, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		logger.Error(err)
	}

	logger.Info(string(b))
}
