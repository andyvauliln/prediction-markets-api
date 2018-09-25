package main

import (
	"fmt"
	"net/http"

	"github.com/andyvauliln/prediction-markets-api/app"
	"github.com/andyvauliln/prediction-markets-api/cmd"
	"github.com/andyvauliln/prediction-markets-api/daos"
	"github.com/andyvauliln/prediction-markets-api/endpoints"
	"github.com/andyvauliln/prediction-markets-api/ethereum"
	"github.com/andyvauliln/prediction-markets-api/services"
	"github.com/andyvauliln/prediction-markets-api/ws"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Get application up and running",
	Long:  `Get application up and running`,
	Run:   run,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func main() {
	cmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	// connect to the database
	_, err := daos.InitSession(nil)
	if err != nil {
		panic(err)
	}

	provider := ethereum.NewWebsocketProvider()

	router := NewRouter(provider)
	http.Handle("/", router)
	http.HandleFunc("/socket", ws.ConnectionEndpoint)

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	log.Info("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func NewRouter(provider *ethereum.EthereumProvider) *mux.Router {

	r := mux.NewRouter()
	log.Info(provider)
	// get daos for dependency injection
	marketDao := daos.NewMarketDao()

	// get services for injection
	marketService := services.NewMarketService(marketDao)

	// deploy http and ws endpoints
	endpoints.ServeMarketResource(r, marketService)

	return r
}
