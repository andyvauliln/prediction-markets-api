package app

import (
	"log"

	"github.com/gorilla/websocket"
)

// Socket acts as a hub that handles messages from the client application and responses
// from the server matching engine.
type Socket struct {
	server      *Server
	connection  *websocket.Conn
	messagesIn  chan *Message
	messagesOut chan *Message
	events      chan *Event
}

// listenToMessagesIn reads incoming messages from the websocket connection
// and sends these messages into the messageIn channel
func (s *Socket) listenToMessagesIn() {
	for {
		message := new(Message)
		err := s.connection.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}
		s.messagesIn <- message
	}
}

// handleMessagesIN listens on the messageIn channel and routes them to the appropriate
// handler based on their MessageType
func (s *Socket) handleMessagesIn() {
	for {
		m := <-s.messagesIn
		switch m.MessageType {
		case DONE:
		default:
			panic("Unknown message type")
		}
	}
}

// handleMessagesOut listens on the event channel (events sent from the matching engine) and routes them
// to the appropriate handler based on their event type
func (s *Socket) handleMessagesOut() {
	for {
		e := <-s.events
		switch e.eventType {
		case DONE:
		default:
			panic("Unknown action type")
		}
	}
}
