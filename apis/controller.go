package apis

import (
	"encoding/json"
	"net/http"
)

func (api *API) GetClusters(w http.ResponseWriter, r *http.Request) {
	err := api.Cluster.Refresh()

	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	jsonResponse(nil, api.Cluster, 200, w)

}

func (api *API) AddNode(w http.ResponseWriter, r *http.Request) {
	payload, err := decodePayload(r)
	if err != nil {
		jsonResponse(err, nil, 400, w)
		return
	}

	err = api.Cluster.AddNode(payload["name"])
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	err = api.Cluster.Refresh()
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	jsonResponse(nil, api.Cluster, 200, w)

}

func (api *API) StartNode(w http.ResponseWriter, r *http.Request) {
	payload, err := decodePayload(r)
	if err != nil {
		jsonResponse(err, nil, 400, w)
		return
	}
	err = api.Cluster.StartNode(payload["id"])
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	err = api.Cluster.Refresh()
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	jsonResponse(nil, api.Cluster, 200, w)

}

func (api *API) StopNode(w http.ResponseWriter, r *http.Request) {
	payload, err := decodePayload(r)
	if err != nil {
		jsonResponse(err, nil, 400, w)
		return
	}
	err = api.Cluster.StopNode(payload["id"])
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	err = api.Cluster.Refresh()
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	jsonResponse(nil, api.Cluster, 200, w)

}

func (api *API) SwitchNode(w http.ResponseWriter, r *http.Request) {
	payload, err := decodePayload(r)
	if err != nil {
		jsonResponse(err, nil, 400, w)
		return
	}
	err = api.Cluster.SwitchDBConnection(payload["id"])
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	err = api.Cluster.Refresh()
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	jsonResponse(nil, api.Cluster, 200, w)
}

func (api *API) MakeQuery(w http.ResponseWriter, r *http.Request) {
	payload, err := decodePayload(r)
	if err != nil {
		jsonResponse(err, nil, 400, w)
		return
	}
	res, err := api.Cluster.Query(payload["query"])
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	jsonResponse(nil, res, 200, w)
}

func (api *API) GetReplicaStatus(w http.ResponseWriter, r *http.Request) {
	res, err := api.Cluster.Query("show status like 'wsrep_cluster%'")
	if err != nil {
		jsonResponse(err, nil, 500, w)
		return
	}
	jsonResponse(nil, res, 200, w)
}

func jsonResponse(err error, data interface{}, status int, w http.ResponseWriter) {
	resp := make(map[string]interface{})
	if err != nil {
		resp["status"] = false
		resp["error"] = err.Error()
	} else {
		resp["status"] = true
		resp["data"] = data
	}
	jsonData, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(jsonData))
}

func decodePayload(r *http.Request) (map[string]string, error) {
	payload := make(map[string]string)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	return payload, err
}
