package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Server type handles a mapping of socket structs and the trading engine
type Server struct {
	clients map[*Socket]bool
	engine  *TradingEngine
}

// NewServer returns a a new empty Server instance
func NewServer() *Server {
	return &Server{
		clients: make(map[*Socket]bool),
		engine:  nil,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/api" {
		w.WriteHeader(http.StatusNotFound)
	}

	s.OpenWebsocketConnection(w, r)
}

func (s *Server) Start() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		s.OpenWebsocketConnection(w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Setup registers a list of quote tokens and token pairs
func (s *Server) Setup(quoteTokens Tokens, pairs TokenPairs, done chan bool) {
	fmt.Printf("Starting server ....\n\n\n")
	s.engine = NewTradingEngine()

	for _, val := range quoteTokens {
		s.engine.RegisterNewQuoteToken(val)
	}

	for _, p := range pairs {
		err := s.engine.RegisterNewPair(p, done)
		if err != nil {
			fmt.Printf("Error registering token pair: %v", err)
		}

	}
}

// OpenWebsocketConnection opens a new websocket connection
func (s *Server) OpenWebsocketConnection(w http.ResponseWriter, r *http.Request) {
	messageOut := make(chan *Message)
	messagesIn := make(chan *Message)
	events := make(chan *Event)
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// defer connection.Close()
	fmt.Printf("Opening new connection ...\n\n")
	socket := &Socket{server: s, connection: connection, messagesOut: messageOut, messagesIn: messagesIn, events: events}

	go socket.handleMessagesOut() //Handle messages from server socket to client
	go socket.handleMessagesIn()  //Handle messages from client to server socket
	socket.listenToMessagesIn()

}
