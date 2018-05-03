package galera

import (
	"github.com/docker/docker/api/types"
)

// Node encapsulates details of a galera node
type Node struct {
	ContainerID string
	Name        string
	Port        types.Port
	Status      string
}
