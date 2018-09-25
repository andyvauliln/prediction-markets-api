package services

import (
	"errors"

	"github.com/andyvauliln/prediction-markets-api/utils"
)

var logger = utils.Logger

var ErrMarketNotFound = errors.New("Market not found")
