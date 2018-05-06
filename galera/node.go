package galera

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Node encapsulates details of a galera node
type Node struct {
	ContainerID string
	Name        string
	Port        uint16
	Status      string
	IP          string
}

const (
	// ErrNoClient is thrown while there is no docker cli client is provided.
	ErrNoClient string = "No docker client provided"
	// ImageName is name of galera docker image.
	ImageName string = "erkules/galera"
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
		fmt.Println(container.Labels)
		if container.Image == ImageName {
			node := Node{
				ContainerID: container.ID,
				Name:        container.Names[0],
				Status:      container.Status,
				IP:          container.NetworkSettings.Networks["bridge"].IPAddress,
			}
			if len(container.Ports) > 0 {
				node.Port = container.Ports[0].PublicPort
			}
			nodes = append(nodes, node)
		}

	}
	return nodes, nil

}

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

func (node *Node) StartNode(cli *client.Client, imageName string) error {
	ctx := context.Background()

	// "bfirsh/reticulate-splines"
	out, err := cli.ImagePull(ctx, ImageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, out)

	config := &container.Config{
		Image: ImageName,
	}

	resp, err := cli.ContainerCreate(ctx, config, nil, nil, node.Name)
	if err != nil {
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	node.ContainerID = resp.ID
	return nil
}

func (node *Node) StopNode(cli *client.Client) error {
	ctx := context.Background()
	return cli.ContainerStop(ctx, node.ContainerID, nil)
}
