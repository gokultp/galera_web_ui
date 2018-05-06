package apis

import (
	"log"
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
	if err != nil {
		return nil, err
	}

	err = api.Cluster.GetCluster()
	if err != nil {
		return nil, err
	}

	api.Router = mux.NewRouter()

	rest := api.Router.PathPrefix("/api").Subrouter()
	rest.HandleFunc("/cluster", api.GetClusters).Methods(http.MethodGet)
	rest.HandleFunc("/node/add", api.AddNode).Methods(http.MethodPost)
	rest.HandleFunc("/node/start", api.StartNode).Methods(http.MethodPost)
	rest.HandleFunc("/node/stop", api.StopNode).Methods(http.MethodPost)
	rest.HandleFunc("/node/switch", api.SwitchNode).Methods(http.MethodPost)
	rest.HandleFunc("/status", api.GetReplicaStatus).Methods(http.MethodPost)
	rest.HandleFunc("/query", api.MakeQuery).Methods(http.MethodPost)

	api.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./client/build/static"))))

	api.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/build/index.html")
	})

	log.Println("Listening")
	return api, nil
}

func (api *API) Listen(port string) error {
	return http.ListenAndServe(port, api.Router)
}
