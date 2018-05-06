package galera

import (
	"context"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Node encapsulates details of a galera node
type Node struct {
	ContainerID string `json:"id"`
	Name        string `json:"name"`
	Port        string `josn:"port"`
	Status      string `josn:"status"`
	IP          string `josn:"ip"`
	Active      bool   `josn:"active"`
}

const (
	// ErrNoClient is thrown while there is no docker cli client is provided.
	ErrNoClient string = "No docker client provided"
	// ErrNodeNotIntialised is thrown while staring node before init
	ErrNodeNotIntialised string = "Node is not created, create node before starting"
	// ImageName is name of galera docker image.
	ImageName string = "erkules/galera"
)

// GetNodes returns existing nodes on the system
func GetNodes(cli *client.Client) ([]Node, error) {
	if cli == nil {
		return nil, errors.New(ErrNoClient)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	nodes := []Node{}
	for _, container := range containers {
		if container.Image == ImageName {
			node := Node{
				ContainerID: container.ID,
				Name:        container.Names[0],
				Status:      container.Status,
				IP:          container.NetworkSettings.Networks["bridge"].IPAddress,
				Port:        "3306",
				Active:      strings.Contains(container.Status, "Up"),
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

func (node *Node) CreateNode(cli *client.Client, clusterIP string) error {
	ctx := context.Background()

	out, err := cli.ImagePull(ctx, ImageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, out)

	config := &container.Config{
		Image: ImageName,
		Cmd:   []string{"--wsrep-cluster-name=local-test", "--wsrep-cluster-address=gcomm://" + clusterIP},
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

// StartNode starts a node
func (node *Node) StartNode(cli *client.Client) error {
	ctx := context.Background()

	if node.ContainerID == "" {
		return errors.New(ErrNodeNotIntialised)
	}
	return cli.ContainerStart(ctx, node.ContainerID, types.ContainerStartOptions{})
}

// StopNode stops a node
func (node *Node) StopNode(cli *client.Client) error {
	ctx := context.Background()
	return cli.ContainerStop(ctx, node.ContainerID, nil)
}

// GetDBConnectionString returns the connection string config for DB in node
func (node *Node) GetDBConnectionString() string {
	return "root@tcp(" + node.IP + ":" + node.Port + ")/"
}
