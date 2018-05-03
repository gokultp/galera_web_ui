package galera

import "github.com/docker/docker/client"

// Cluster struct encapsulates informations about cluster like nodes in cluster
type Cluster struct {
	Nodes  []Node
	IP     string
	Client *client.Client
}
