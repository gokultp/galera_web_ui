package apis

import (
	"net/http"

	"github.com/gokultp/galera_web_ui/galera"
	"github.com/gorilla/mux"
)

type API struct {
	Cluster *galera.Cluster
	Router  *mux.Router
}

func NewAPI() (*API, error) {
	var err error
	api := &API{}
	api.Cluster, err = galera.NewCluster()
	api.Router = mux.NewRouter()
	if err != nil {
		return nil, err
	}

	api.Router.HandleFunc("/cluster", api.GetClusters).Methods(http.MethodGet)
	api.Router.HandleFunc("/node", api.AddNode).Methods(http.MethodPost)
	api.Router.HandleFunc("/node/start", api.StartNode).Methods(http.MethodPost)
	api.Router.HandleFunc("/node/stop", api.StopNode).Methods(http.MethodPost)
	api.Router.HandleFunc("/node/switch", api.SwitchNode).Methods(http.MethodPost)
	api.Router.HandleFunc("/status", api.GetReplicaStatus).Methods(http.MethodPost)
	api.Router.HandleFunc("/query", api.MakeQuery).Methods(http.MethodPost)

	return api, nil
}

func (api *API) Listen(port string) error {
	return http.ListenAndServe(port, api.Router)
}
