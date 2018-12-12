package kademlia

// Node encapsulates all of the relevant information for a particular node in
// the Kademlia network.
type Node struct {
	ID        string // The (hexadecimal) ID of this node.
	IPAddress string // The IP address at which to contact this node.
	Port      uint16 // The port on which to contact this node.

	// All of the requests this node has sent, for which we haven't yet gotten a
	// response. This allows for managing timeouts and mapping incoming
	// responses to specific requests.
	outstandingRequests map[string]*NodeRequest
}

// NodeResponseHandler allows us to define callbacks to handle responses to very
// specific requests.
type NodeResponseHandler func(req *NodeRequest, res *NodeResponse) error

// NewNode builds a new Kademlia node with the given parameters.
func NewNode(id, ipAddress string, port uint16) *Node {
	return &Node{
		ID:        id,
		IPAddress: ipAddress,
		Port:      port,
	}
}

func (n *Node) SendRequest(req *NodeRequest, handler NodeResponseHandler) {

}
