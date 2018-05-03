package galera

import (
	"context"
	"errors"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Node encapsulates details of a galera node
type Node struct {
	ContainerID string
	Name        string
	Ports       []types.Port
	Status      string
}

const (
	// ErrNoClient is thrown while there is no docker cli client is provided
	ErrNoClient string = "No docker client provided"
)

// GetNodes returns existing nodes on the system
func GetNodes(cli *client.Client) ([]Node, error) {
	if cli == nil {
		return nil, errors.New(ErrNoClient)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	nodes := []Node{}
	for _, container := range containers {

		if strings.HasPrefix(container.Names[0], "galera_") {

			nodes = append(nodes, Node{
				ContainerID: container.ID,
				Name:        container.Names[0],
				Ports:       container.Ports,
				Status:      container.Status,
			})
		}

	}
	return nodes, nil

}
