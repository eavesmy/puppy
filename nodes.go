package puppy

type Node struct {
	Id      string
	State   int
	Heath   int
	Self    bool
	Name    string
	Group   int32
	Version int32
	Type    int32
	Routes  map[string]rpcMethod
}

type NodeMgr struct {
	nodes map[string]*Node
}

var nodeMgr = &NodeMgr{nodes: make(map[string]*Node)}

func (n *NodeMgr) AddSelfNode(node *Node) {
	// node.Id =
}

func (n *NodeMgr) AddNode(root string) {
	node := &Node{
		Id:     "",
		Routes: map[string]rpcMethod{},
	}
	n.nodes[node.Id] = node
}

func (n *NodeMgr) RmNode(id string) {
	n.nodes[id] = nil
}

func genId(group, version int32) {

}
