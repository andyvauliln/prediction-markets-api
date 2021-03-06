package daos

import (
	"errors"
	"time"

	"github.com/andyvauliln/prediction-markets-api/app"
	"github.com/andyvauliln/prediction-markets-api/types"
	"github.com/ethereum/go-ethereum/common"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MarketDao contains:
// collectionName: MongoDB collection name
// dbName: name of mongodb to interact with
type MarketDao struct {
	collectionName string
	dbName         string
}

// Market dao options
type MarketDaoOption = func(*MarketDao) error

// Market Dao pre-settings
func MarketDaoDBOption(dbName string) func(dao *MarketDao) error {
	return func(dao *MarketDao) error {
		dao.dbName = dbName
		return nil
	}
}

// NewMarketDao returns a new instance of AddressDao
func NewMarketDao(options ...MarketDaoOption) *MarketDao {
	dao := &MarketDao{}
	dao.collectionName = "markets"
	dao.dbName = app.Config.DBName

	for _, op := range options {
		err := op(dao)
		if err != nil {
			panic(err)
		}
	}

	index := mgo.Index{
		Key:    []string{"MarketID"},
		Unique: true,
	}

	err := db.Session.DB(dao.dbName).C(dao.collectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	return dao
}

// Create function performs the DB insertion task for Market collection
func (dao *MarketDao) Create(market *types.Market) error {
	market.ID = bson.NewObjectId()
	market.CreatedAt = time.Now()
	market.UpdatedAt = time.Now()

	err := db.Create(dao.dbName, dao.collectionName, market)
	return err
}

// GetAll function fetches all the markets in the market collection of mongodb.
func (dao *MarketDao) GetAll() ([]types.Market, error) {
	var response []types.Market
	err := db.Get(dao.dbName, dao.collectionName, bson.M{}, 0, 0, &response)
	return response, err
}

// GetByID function fetches details of a market using market's mongo ID.
func (dao *MarketDao) GetByID(id bson.ObjectId) (*types.Market, error) {
	var response *types.Market
	err := db.GetByID(dao.dbName, dao.collectionName, id, &response)
	return response, err
}

// GetByTokenAddress function fetches Market based on
// CONTRACT ADDRESS of base token and quote token
func (dao *MarketDao) GetByMarketAddress(marketID common.Address) (*types.Market, error) {
	var res []*types.Market

	q := bson.M{
		"marketID": marketID.Hex(),
	}

	err := db.Get(dao.dbName, dao.collectionName, q, 0, 1, &res)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("Market not found")
	}

	return res[0], nil
}
