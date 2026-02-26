package cluster

type Cluster struct {
	Name  string    `json:"name"`
	Nodes []*string `json:"nodes"`
	Self  string    `json:"self"`
}
