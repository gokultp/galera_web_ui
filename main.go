package main

import (
	"fmt"

	"github.com/gokultp/galera_web_ui/galera"
)

func main() {
	cluster, err := galera.NewCluster()
	fmt.Println(err)
	err = cluster.GetCluster()
	fmt.Println(err)

	for _, node := range cluster.Nodes {
		fmt.Println(node.ContainerID, node.Name)
	}
}
