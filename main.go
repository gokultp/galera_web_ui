package main

import (
	"fmt"

	"github.com/gokultp/galera_web_ui/galera"
)

func main() {
	cluster, err := galera.NewCluster()
	fmt.Println(err)

	// node := galera.NewNode("galera_test3", 3000)
	// err = node.CreateNode(cluster.Client, "bfirsh/reticulate-splines")
	// fmt.Println(err)

	err = cluster.GetCluster()
	fmt.Println(err)

	fmt.Println(cluster)
	for _, node := range cluster.Nodes {
		fmt.Println(node.ContainerID, node.Name, node.IP, node.Port)
	}

	cluster.Nodes[0].RunQuery()
}
