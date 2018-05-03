package galera

import "github.com/docker/docker/client"

// Cluster struct encapsulates informations about cluster like nodes in cluster
type Cluster struct {
	Nodes  []Node
	IP     string
	Client *client.Client
}

// GetCluster gets all cluster details
func (c *Cluster) GetCluster() error {
	nodes, err := GetNodes(c.Client)
	if err != nil {
		return err
	}
	c.Nodes = nodes
	return nil
}
