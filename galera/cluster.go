package galera

import (
	"database/sql"

	"github.com/docker/docker/client"
	_ "github.com/go-sql-driver/mysql"
)

// Cluster struct encapsulates informations about cluster like nodes in cluster
type Cluster struct {
	Name          string
	Nodes         []Node
	IP            string
	Client        *client.Client
	ConnectedNode int
	DB            *sql.DB
}

// NewCluster creates a new clusten instance (Constructor like function)
func NewCluster() (*Cluster, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Cluster{
		Client: cli,
	}, nil

}

// GetCluster gets all cluster details
func (c *Cluster) GetCluster() error {
	nodes, err := GetNodes(c.Client)
	if err != nil {
		return err
	}
	c.Nodes = nodes
	if len(nodes) == 0 {
		return nil
	}
	c.DB, err = sql.Open("mysql", "root@tcp("+nodes[0].IP+":3306)/test")

	if err != nil {
		return err
	}
	return nil

}

// AddNode adds node to the cluster
func (c *Cluster) AddNode() error {
	return nil
}

// Query will run query on selected cluster
func (c *Cluster) Query(query string) ([]map[string]string, error) {

	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToMap(rows)

}

// SwitchDBConnection will switch db connection to given node
func (c *Cluster) SwitchDBConnection(nodeIndex int) error {
	err := c.DB.Close()
	if err != nil {
		return err
	}
	c.DB, err = sql.Open("mysql", "root@tcp("+c.Nodes[nodeIndex].IP+":3306)/test")
	return err
}

// rowsToMap converts SQL row to string map, which is easier to convert to JSON
func rowsToMap(rows *sql.Rows) (results []map[string]string, err error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	valPointers := make([][]byte, len(columns))
	vals := make([]interface{}, len(columns))

	for rowIndex := range valPointers {
		vals[rowIndex] = &valPointers[rowIndex]
	}

	for rows.Next() {
		err := rows.Scan(vals...)
		if err != nil {
			return nil, err
		}
		res := make(map[string]string)

		for colIndex := range columns {
			res[columns[colIndex]] = string(valPointers[colIndex])
		}
		results = append(results, res)
	}
	return results, nil

}
