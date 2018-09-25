package services

import (
	"github.com/Proofsuite/amp-matching-engine/interfaces"

	"gopkg.in/mgo.v2/bson"

	"github.com/Proofsuite/amp-matching-engine/types"
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
	p, err := s.marketDao.GetByID(market.Id)
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
