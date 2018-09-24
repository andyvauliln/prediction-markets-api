package app

import (
	"flag"
	"log"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/posener/wstest"
)

var wg = &sync.WaitGroup{}
var addr = flag.String("addr", "localhost:8080", "http service address")

type Client struct {
	connection   *websocket.Conn
	requests     chan *Message
	responses    chan *Message
	requestLogs  []*Message
	responseLogs []*Message
}

func NewClient(s *Server) *Client {
	flag.Parse()
	uri := url.URL{Scheme: "ws", Host: *addr, Path: "/api"}
	log.Printf("Connecting to %s", uri.String())

	d := wstest.NewDialer(s)
	c, _, err := d.Dial(uri.String(), nil)
	if err != nil {
		panic(err)
	}

	requests := make(chan *Message)
	responses := make(chan *Message)
	requestLogs := make([]*Message, 0)
	responseLogs := make([]*Message, 0)

	return &Client{connection: c,
		requests:     requests,
		responses:    responses,
		requestLogs:  requestLogs,
		responseLogs: responseLogs,
	}
}

func (c *Client) start() {
	c.handleMessages()
	c.handleIncomingMessages()
}

func (c *Client) handleMessages() {
	go func() {
		for {
			select {
			case request := <-c.requests:
				c.requestLogs = append(c.requestLogs, request)
				switch request.MessageType {
				case DONE:
					go c.done()
				}
			case response := <-c.responses:
				c.responseLogs = append(c.responseLogs, response)
				switch response.MessageType {
				case DONE:
					go c.done()
				}
			}
		}
	}()
}

func (c *Client) handleIncomingMessages() {
	message := new(Message)
	go func() {
		for {
			err := c.connection.ReadJSON(&message)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Error: %#v", err)
				}
				break
			}

			c.responses <- message
		}
	}()
}

func (c *Client) done() {

}
