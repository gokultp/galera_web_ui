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
