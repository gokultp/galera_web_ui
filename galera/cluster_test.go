package galera_test

import (
	"testing"

	"github.com/gokultp/galera_web_ui/galera"
)

func TestNewCluster(t *testing.T) {
	cluster, err := galera.NewCluster()

	if err != nil {
		t.Error("unexpected error")
	}

	if cluster == nil {
		t.Error("cluster should not be nil")
	}

	if cluster.Client == nil {
		t.Errorf("cluster client should not be nil")
	}
}

func TestGetCluster(t *testing.T) {
	cluster, err := galera.NewCluster()

	if err != nil {
		t.Error("unexpected error")
	}

	err = cluster.GetCluster()

	if err != nil {
		t.Error("unexpected error on getting cluster")
	}
}

func TestRefreshCluster(t *testing.T) {
	cluster, err := galera.NewCluster()

	if err != nil {
		t.Error("unexpected error")
	}
	err = cluster.GetCluster()

	if err != nil {
		t.Error("unexpected error on getting cluster")
	}

	err = cluster.Refresh()

	if err != nil {
		t.Error("unexpected error on refreshing cluster")
	}
}

func TestAddNode(t *testing.T) {
	cluster, err := galera.NewCluster()

	if err != nil {
		t.Error("unexpected error")
	}
	err = cluster.GetCluster()

	if err != nil {
		t.Error("unexpected error on getting cluster")
	}

	err = cluster.AddNode("test_node")
	if err != nil {
		t.Error("unexpected error on adding node to cluster")
	}
	err = cluster.Refresh()

	if err != nil {
		t.Error("unexpected error on refreshing cluster")
	}

	if len(cluster.Nodes) == 0 {
		t.Error("Node length should not be zero")
	}
	var sel_node *galera.Node
	for _, node := range cluster.Nodes {
		if node.Name == "test_node" {
			sel_node = &node
			break
		}
	}
	if sel_node == nil {
		t.Error("Node is not added to cluster")
	}
}

func TestStopNode(t *testing.T) {
	cluster, err := galera.NewCluster()

	if err != nil {
		t.Error("unexpected error")
	}
	err = cluster.GetCluster()

	if err != nil {
		t.Error("unexpected error on getting cluster")
	}

	var sel_node, sel_node1 *galera.Node
	for _, node := range cluster.Nodes {
		if node.Active == true {
			sel_node = &node
			break
		}
	}

	err = cluster.StopNode(sel_node.ContainerID)
	if err != nil {
		t.Error("unexpected error on stopping node to cluster")
	}
	err = cluster.Refresh()

	if err != nil {
		t.Error("unexpected error on refreshing cluster")
	}

	for _, node := range cluster.Nodes {
		if node.ContainerID == sel_node1.ContainerID {
			sel_node1 = &node
			break
		}
	}

	if sel_node1.Active == true {
		t.Error("Node didn't stopped")
	}

}

func TestStartNode(t *testing.T) {
	cluster, err := galera.NewCluster()

	if err != nil {
		t.Error("unexpected error")
	}
	err = cluster.GetCluster()

	if err != nil {
		t.Error("unexpected error on getting cluster")
	}

	var sel_node, sel_node1 *galera.Node
	for _, node := range cluster.Nodes {
		if node.Active == false {
			sel_node = &node
			break
		}
	}

	err = cluster.StartNode(sel_node.ContainerID)
	if err != nil {
		t.Error("unexpected error on starting node to cluster")
	}
	err = cluster.Refresh()

	if err != nil {
		t.Error("unexpected error on refreshing cluster")
	}

	for _, node := range cluster.Nodes {
		if node.ContainerID == sel_node1.ContainerID {
			sel_node1 = &node
			break
		}
	}

	if sel_node1.Active == false {
		t.Error("Node didn't started")
	}

}
