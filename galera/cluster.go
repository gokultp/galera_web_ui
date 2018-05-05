package galera

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/docker/docker/client"
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

func (c *Cluster) AddNode() error {
	return nil
}

// func (c *Cluster) RunQuery(query string) (interface{}, error) {
func (c *Cluster) RunQuery(query string) {

	rows, err := c.DB.Query(query)
	if err != nil {
		log.Fatal(err)

	}

	columns, err := rows.Columns()
	valp := make([]interface{}, len(columns))

	vals := make([]interface{}, len(columns))

	for i := range valp {
		vals[i] = &valp[i]
	}

	fmt.Println(columns)

	for rows.Next() {
		err := rows.Scan(vals...)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(valp[0].([]byte)), string(valp[0].([]byte)))
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

}

// func rowsToJSON(rows *sql.Rows) (result map[string]interface{}, err error) {
// 	columns, err := rows.Columns()

// 	defer rows.Close()
// }
