package packages

func InitGraph() *Graph {
	return &Graph{
		nodes: []*Node{},
	}
}

func InitNode(id int) *Node {
	return &Node{
		id:    id,
		edges: make(map[int]int),
	}
}

type Graph struct {
	nodes []*Node
}

type Node struct {
	id    int
	edges map[int]int
}

//GetId return node's id
func (n *Node) GetId() int {
	return n.id
}

//GetEdges return node's edges
func (n *Node) GetEdges() map[int]int {
	return n.edges
}

//AddNode for add new node
//parameter : none
// return : id of new node
func (graph *Graph) AddNode() (id int) {
	id = len(graph.nodes)

	graph.nodes = append(graph.nodes, InitNode(id))

	return id
}

//AddEdges for add edges,
//parameter : node id (int), node destinaton id (int), weight of the edge (int)
//return : none
func (graph *Graph) AddEdges(nodeId, dstId, wight int) {

	//the edges will hold destination id as a key and weight as a value

	graph.nodes[nodeId].edges[dstId] = wight

}

//NeighBors
//parameter : node's id (int)
//returns  : a list of node that are linked to this node ([]*Node)
func (graph *Graph) NeighBors(id int) (result []*Node) {

	//iterate all nodes
	for _, node := range graph.nodes {

		for edge := range node.edges {

			//if the node's id is equal parameter id, then take the destination id from edge's key
			if node.id == id {
				result = append(result, graph.nodes[edge])
			}

			//if the edge key is equal id then take the node
			if edge == id {
				result = append(result, node)
			}

		}

	}

	return result
}

//GetNodes return all node
func (graph *Graph) GetNodes() (result []*Node) {
	return graph.nodes
}

//GetEdges return all edges as a slice wich contain [start node's id, end node's id, weight]
func (graph *Graph) GetEdges() (result [][3]int) {

	for _, n := range graph.nodes {

		for k, v := range n.edges {
			result = append(result, [3]int{n.id, k, v})
		}

	}

	return result
}
