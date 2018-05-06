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

	// err = cluster.AddNode("testnode")
	// fmt.Println(err)

	// err = cluster.GetCluster()
	// fmt.Println(err)
	fmt.Println(cluster.StartNode(cluster.Nodes[0].ContainerID))

	for _, node := range cluster.Nodes {
		fmt.Println(node.ContainerID, node.Name, node.IP, node.Port, node.Status)
	}

	fmt.Print(cluster.Query("show status like 'wsrep_cluster_size'"))
}
