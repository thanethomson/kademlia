package kademlia

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Node encapsulates all of the relevant information for a particular node in
// the Kademlia network.
type Node struct {
	ID        NodeID    // The (hexadecimal) ID of this node.
	IPAddress string    // The IP address at which to contact this node.
	Port      uint16    // The port on which to contact this node.
	FirstSeen time.Time // When did we first see this node?
	LastSeen  time.Time // The system time at which this node was last seen.
	Uptime    uint64    // For how many seconds have we known this node to be up and running?

	client *http.Client // The client we're using to communicate with this node.
}

// NodeResponseHandler allows us to define callbacks to handle responses to very
// specific requests.
type NodeResponseHandler func(req *NodeRequest, res *NodeResponse) error

// NewNode builds a new Kademlia node representation with the given parameters.
func NewNode(id NodeID, ipAddress string, port uint16) *Node {
	now := time.Now()
	return &Node{
		ID:        id,
		IPAddress: ipAddress,
		Port:      port,
		FirstSeen: now,
		LastSeen:  now,
		Uptime:    0,
		client:    &http.Client{},
	}
}

// SendRequest is responsible for wrapping the low-level sending of a request
// for a particular node. It also tracks requests such that we can map future
// responses back to requests.
func (n *Node) SendRequest(req *NodeRequest, handler NodeResponseHandler) error {
	nodeURL, err := url.Parse(fmt.Sprintf("http://%s:%d", n.IPAddress, n.Port))
	if err != nil {
		return err
	}
	var httpReq *http.Request

	// construct the request
	switch req.Type {
	case ReqPing:
		nodeURL.Path += "/ping/" + req.ID
		httpReq, err := http.NewRequest("GET", nodeURL.String(), nil)
	case ReqStore:
		params, ok := req.Params.(NodeRequestStoreParams)
		if !ok {
			return fmt.Errorf("Missing request parameters for storing value")
		}
		nodeURL.Path += "/store/" + params.Key
	default:
		return fmt.Errorf("Unsupported request type: %s", req.Type)
	}

	go func() {
		resp, err := n.client.Do(httpReq)
		handler(req, nil)
	}()
	return nil
}
