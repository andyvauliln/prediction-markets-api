package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/andyvauliln/prediction-markets-api/interfaces"
	"github.com/andyvauliln/prediction-markets-api/services"
	"github.com/andyvauliln/prediction-markets-api/types"
	"github.com/andyvauliln/prediction-markets-api/utils/httputils"
	"github.com/gorilla/mux"
)

type marketEndpoint struct {
	marketService interfaces.MarketService
}

// ServeMarketResource sets up the routing of markets endpoints and the corresponding handlers.
func ServeMarketResource(
	r *mux.Router,
	p interfaces.MarketService,
) {
	e := &marketEndpoint{p}
	r.HandleFunc("/market", e.HandleCreateMarket).Methods("POST")
	r.HandleFunc("/market/{id}", e.HandleGetMarket).Methods("GET")
	r.HandleFunc("/markets", e.HandleGetAllMarkets).Methods("GET")
}

func (e *marketEndpoint) HandleCreateMarket(w http.ResponseWriter, r *http.Request) {
	p := &types.Market{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(p)
	if err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	defer r.Body.Close()

	err = p.Validate()
	if err != nil {
		httputils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = e.marketService.Create(p)
	if err != nil {
		switch err {
		case services.ErrMarketExists:
			httputils.WriteError(w, http.StatusBadRequest, "Market exists")
			return
		case services.ErrMarketNotFound:
			httputils.WriteError(w, http.StatusBadRequest, "Market not found")
			return
		case services.ErrmarketInvalid:
			httputils.WriteError(w, http.StatusBadRequest, "Market invalid")
			return
		default:
			logger.Error(err)
			httputils.WriteError(w, http.StatusInternalServerError, "")
			return
		}
	}

	httputils.WriteJSON(w, http.StatusCreated, p)
}

func (e *marketEndpoint) HandleGetAllMarkets(w http.ResponseWriter, r *http.Request) {
	res, err := e.marketService.GetAll()
	if err != nil {
		logger.Error(err)
		httputils.WriteError(w, http.StatusInternalServerError, "")
		return
	}

	httputils.WriteJSON(w, http.StatusOK, res)
}

func (e *marketEndpoint) HandleGetMarket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	MarketID := vars["id"]

	if !common.IsHexAddress(MarketID) {
		httputils.WriteError(w, http.StatusBadRequest, "Invalid Address")
	}

	MarketIDHex := common.HexToAddress(MarketID)
	quoteTokenAddress := common.HexToAddress(quoteToken)
	res, err := e.marketService.GetByMarketAddress(MarketIDHex)
	if err != nil {
		logger.Error(err)
		httputils.WriteError(w, http.StatusInternalServerError, "")
		return
	}

	httputils.WriteJSON(w, http.StatusOK, res)
}
