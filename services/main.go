package services

import (
	"errors"

	"github.com/andyvauliln/amp-matching-engine/utils"
)

var logger = utils.Logger

var ErrMarketNotFound = errors.New("Market not found")
