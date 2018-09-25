package services

import (
	"errors"

	"github.com/Proofsuite/amp-matching-engine/utils"
)

var logger = utils.Logger

var ErrMarketNotFound = errors.New("Market not found")
