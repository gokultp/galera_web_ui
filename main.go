package main

import (
	"fmt"

	"github.com/gokultp/galera_web_ui/galera"
)

func main() {
	cluster, err := galera.NewCluster()
	fmt.Println(err)

	err = cluster.AddNode("testnode")

	err = cluster.GetCluster()
	fmt.Println(err)

	fmt.Println(cluster)
	for _, node := range cluster.Nodes {
		fmt.Println(node.ContainerID, node.Name, node.IP, node.Port)
	}

	fmt.Print(cluster.Query("create table nums3(a integer, b integer, c integer)"))
}
