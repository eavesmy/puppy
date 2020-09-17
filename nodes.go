package puppy

type Node struct {
	id     string
	routes map[string]rpcMethod
}
