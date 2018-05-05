package galera

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"database/sql"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
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
		// if strings.HasPrefix(container.Names[0], "/galera_") {
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
		// }

	}
	return nodes, nil

}

func NewNode(name string, port uint16) *Node {
	return &Node{
		Name: name,
		Port: port,
	}
}

func (node *Node) CreateNode(cli *client.Client, imageName string) error {
	ctx := context.Background()

	// "bfirsh/reticulate-splines"
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, out)

	strPort := strconv.Itoa(int(node.Port))
	config := &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			nat.Port(strPort): struct{}{},
		},
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(strPort): []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: strPort,
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, node.Name)
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

func (node *Node) RunQuery() {
	// var id int
	// var name string

	db, err := sql.Open("mysql",
		"tcp(127.0.0.1:3306)")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("show status like ?", "wsrep_cluster_size")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()

	fmt.Println(columns)

	// for rows.Next() {
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(id, name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	defer db.Close()
}
