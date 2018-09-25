package ws

var marketSocket *MarketSocket

// MarketSocket holds the map of connections subscribed to pair channels
// corresponding to the key/event they have subscribed to.
type MarketSocket struct {
	subscriptions map[string]map[*Conn]bool
}

func GetMarketSocket() *MarketSocket {
	if marketSocket == nil {
		marketSocket = &MarketSocket{make(map[string]map[*Conn]bool)}
	}

	return marketSocket
}

// Subscribe registers a new websocket connections to the Market channel updates
func (s *MarketSocket) Subscribe(channelID string, conn *Conn) error {
	if s.subscriptions[channelID] == nil {
		s.subscriptions[channelID] = make(map[*Conn]bool)
	}

	s.subscriptions[channelID][conn] = true
	return nil
}

// Unsubscribe removes a websocket connection from the Market channel updates
func (s *MarketSocket) Unsubscribe(channelID string, conn *Conn) {
	if s.subscriptions[channelID][conn] {
		s.subscriptions[channelID][conn] = false
		delete(s.subscriptions[channelID], conn)
	}
}

// UnsubscribeHandler unsubscribes a connection from a certain Market channel id
func (s *MarketSocket) UnsubscribeHandler(channelID string) func(conn *Conn) {
	return func(conn *Conn) {
		s.Unsubscribe(channelID, conn)
	}
}

// BroadcastMessage broadcasts Market message to all subscribed sockets
func (s *MarketSocket) BroadcastMessage(channelID string, p interface{}) {
	go func() {
		for conn, active := range MarketSocket.subscriptions[channelID] {
			if active {
				s.SendUpdateMessage(conn, p)
			}
		}
	}()
}

// SendMessage sends a websocket message on the Market channel
func (s *MarketSocket) SendMessage(conn *Conn, msgType string, p interface{}) {
	SendMessage(conn, MarketChannel, msgType, p)
}

// SendErrorMessage sends an error message on the Market channel
func (s *MarketSocket) SendErrorMessage(conn *Conn, p interface{}) {
	s.SendMessage(conn, "ERROR", p)
}

// SendInitMessage is responsible for sending message on Market ohlcv channel at subscription
func (s *MarketSocket) SendInitMessage(conn *Conn, p interface{}) {
	s.SendMessage(conn, "INIT", p)
}

// SendUpdateMessage is responsible for sending message on Market ohlcv channel at subscription
func (s *MarketSocket) SendUpdateMessage(conn *Conn, p interface{}) {
	s.SendMessage(conn, "UPDATE", p)
}
