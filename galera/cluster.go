package galera

import "github.com/docker/docker/client"

// Cluster struct encapsulates informations about cluster like nodes in cluster
type Cluster struct {
	Nodes  []Node
	IP     string
	Client *client.Client
}

// NewCluster creates a new clusten instance (Constructor like function)
func NewCluster() (*Cluster, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Cluster{
		Client: cli,
	}, nil

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
