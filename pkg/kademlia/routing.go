package kademlia

// K represents the maximum number of nodes we want to track in a particular
// bucket. This has an effect on memory consumption.
const K int = 20

// NodeBucket allows us to keep track of a number of known nodes in the network
// whose node IDs are a specific distance from our current node's ID.
type NodeBucket struct {
	Distance  int    // The distance between the nodes in this list and our node.
	NodeCount int    // The number of nodes in this bucket.
	Nodes     []Node // The nodes in this bucket (dynamically allocated for reduced memory consumption).
}
