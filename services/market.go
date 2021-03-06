package services

import (
	"github.com/andyvauliln/prediction-markets-api/interfaces"
	"github.com/andyvauliln/prediction-markets-api/utils"
	"github.com/andyvauliln/prediction-markets-api/ws"
	"github.com/ethereum/go-ethereum/common"

	"gopkg.in/mgo.v2/bson"

	"github.com/andyvauliln/prediction-markets-api/types"
)

// MarketService struct with daos required, responsible for communicating with daos.
// MarketService functions are responsible for interacting with daos and implements business logics.
type MarketService struct {
	marketDao interfaces.MarketDao
}

// NewMarketService returns a new instance of balance service
func NewMarketService(
	marketDao interfaces.MarketDao,

) *MarketService {

	return &MarketService{marketDao}
}

// Create function is responsible for inserting new market in DB.
// It checks for existence of market in DB first
func (s *MarketService) Create(market *types.Market) error {
	p, err := s.marketDao.GetByID(market.ID)
	if err != nil {
		return err
	}

	if p != nil {
		return ErrMarketExists
	}

	err = s.marketDao.Create(market)
	if err != nil {
		return err
	}

	return nil
}

// GetByID fetches details of a market using its mongo ID
func (s *MarketService) GetByID(id bson.ObjectId) (*types.Market, error) {
	return s.marketDao.GetByID(id)
}

// GetAll is reponsible for fetching all the markets in the DB
func (s *MarketService) GetAll() ([]types.Market, error) {
	return s.marketDao.GetAll()
}

// GetTrades is currently not implemented correctly
func (s *MarketService) GetMarkets(bt, qt common.Address) ([]types.Market, error) {
	return s.marketDao.GetAll()
}

// GetByAddress fetches the detailed document of a token using its contract address
func (s *MarketService) GetByMarketAddress(marketID common.Address) (*types.Market, error) {
	return s.marketDao.GetByMarketAddress(marketID)
}

// Subscribe
func (s *MarketService) Subscribe(conn *ws.Conn, bt, qt common.Address) {
	socket := ws.GetMarketSocket()

	markets, err := s.GetMarkets(bt, qt)
	if err != nil {
		socket.SendErrorMessage(conn, err.Error())
		return
	}

	id := utils.GetMarketChannelID(bt, qt)
	err = socket.Subscribe(id, conn)
	if err != nil {
		message := map[string]string{
			"Code":    "UNABLE_TO_REGISTER",
			"Message": "UNABLE_TO_REGISTER " + err.Error(),
		}

		socket.SendErrorMessage(conn, message)
		return
	}

	ws.RegisterConnectionUnsubscribeHandler(conn, socket.UnsubscribeHandler(id))
	socket.SendInitMessage(conn, markets)
}

// Unsubscribe
func (s *MarketService) Unsubscribe(conn *ws.Conn, bt, qt common.Address) {
	socket := ws.GetMarketSocket()

	id := utils.GetMarketChannelID(bt, qt)
	socket.Unsubscribe(id, conn)
}
