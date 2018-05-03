package galera

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Node encapsulates details of a galera node
type Node struct {
	ContainerID string
	Name        string
	Port        types.Port
	Status      string
}

// GetNodes returns existing nodes on the system
func GetNodes(cli *client.Client) ([]Node, error) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	nodes := []Node{}
	for _, container := range containers {
		nodes = append(nodes, Node{
			ContainerID: container.ID,
			Name:        container.Names[0],
			Port:        container.Ports[0],
			Status:      container.Status,
		})
	}

	return nodes, nil

}
