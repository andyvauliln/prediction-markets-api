package app

import (
	"encoding/json"
	"fmt"

	"github.com/kr/pretty"
)

type MessageType string

const (
	GET_MARKETS = "GET_MARKETS"
	DONE        = "DONE"
)

type Message struct {
	MessageType MessageType `json:"messageType"`
	Payload     Payload     `json:"payload"`
}

type Payload interface{}

type GetMarktesPayload struct {
	Market *Market `json:"market"`
}

func (m *Message) UnmarshalJSON(b []byte) error {
	message := map[string]interface{}{}

	err := json.Unmarshal(b, &message)
	if err != nil {
		return err
	}

	m.MessageType = MessageType(message["messageType"].(string))
	m.Payload = message["payload"]

	return nil
}

func (m *Message) String() string {
	return fmt.Sprintf("\nMessage:\nMessageType: %v\nPayload:\n%v\n", pretty.Formatter(m.MessageType), pretty.Formatter(m.Payload))
}

type GetMarketsMessage struct {
	MessageType MessageType       `json:"messageType"`
	Payload     GetMarktesPayload `json:"payload"`
}
