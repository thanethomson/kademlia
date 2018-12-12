package kademlia

import (
	"encoding/json"
	"fmt"
	"time"
)

// NodeRequestType allows us to represent the different request types as
// constants.
type NodeRequestType string

// The various request types our nodes can handle.
const (
	ReqPing      NodeRequestType = "ping"
	ReqStore     NodeRequestType = "store"
	ReqFindNode  NodeRequestType = "findNode"
	ReqFindValue NodeRequestType = "findValue"
)

// NodeRequest is a dynamic message type that our nodes can handle.
type NodeRequest struct {
	ID     string          `json:"id"`               // A request correlation ID (in case of UDP transport, to match incoming responses).
	Type   NodeRequestType `json:"type"`             // What type of request is this?
	Params interface{}     `json:"params,omitempty"` // The dynamically deserialised message body.

	// The timestamp for when this request was sent.
	sent time.Time
}

// NodeResponseType allows us to represent our different responses as constants.
type NodeResponseType string

// The various response types our nodes can return.
const (
	ResPong  NodeResponseType = "pong"  // Response to a "ping" request.
	ResStore NodeResponseType = "store" // Response to when storage is requested.
	ResNode  NodeResponseType = "node"  // Provide a node's details.
	ResValue NodeResponseType = "value" // Provides the value of a key/value pair.
)

type NodeResponseStatus string

// The possible statuses of a node response.
const (
	StatusOK            NodeResponseStatus = "ok"
	StatusFailed        NodeResponseStatus = "failed"        // For example, if storage failed.
	StatusNotFound      NodeResponseStatus = "notFound"      // For when a node or key/value pair cannot be found.
	StatusInternalError NodeResponseStatus = "internalError" // "Internal server error".
)

// NodeResponse models the responses our nodes can send back to other nodes.
type NodeResponse struct {
	RequestID string             `json:"requestId"`        // The ID of the request for which this response is relevant.
	Status    NodeResponseStatus `json:"status"`           // What is the status of this response?
	Type      NodeResponseType   `json:"type"`             // What type of message is this? (for deserialisation)
	Params    interface{}        `json:"params,omitempty"` // The deserialised body of the message.
}

type NodeRequestStoreParams struct {
	Key   string  `json:"key"`
	Value *string `json:"value"` // Can be null if we want to delete the key.
}

type NodeRequestFindNodeParams struct {
	ID string `json:"id"` // The ID of the node we're trying to find.
}

type NodeRequestFindValueParams struct {
	Key string `json:"key"` // The key for which we need to look up its value.
}

type NodeResponseStoreParams struct {
	PreviousValue *string `json:"previousValue"` // The previous value, if any, that was replaced by the store request.
}

type NodeResponseFindNodeParams struct {
	ID        string `json:"id"`        // The ID of the node.
	IPAddress string `json:"ipAddress"` // The IP address of the node we've found.
	Port      uint16 `json:"port"`
}

type NodeResponseFindValueParams struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
}

// ParseNodeRequest attempts to parse the given JSON string into a NodeRequest
// object.
func ParseNodeRequest(s string) (*NodeRequest, error) {
	var rawParams json.RawMessage
	req := NodeRequest{
		Params: &rawParams,
	}
	if err := json.Unmarshal([]byte(s), &req); err != nil {
		return nil, err
	}
	// now check which type of request this is
	switch req.Type {
	case ReqPing:
		// no parameters for ping
		req.Params = nil
	case ReqStore:
		var p NodeRequestStoreParams
		if err := json.Unmarshal(rawParams, &p); err != nil {
			return nil, err
		}
		req.Params = p
	case ReqFindNode:
		var p NodeRequestFindNodeParams
		if err := json.Unmarshal(rawParams, &p); err != nil {
			return nil, err
		}
		req.Params = p
	case ReqFindValue:
		var p NodeRequestFindValueParams
		if err := json.Unmarshal(rawParams, &p); err != nil {
			return nil, err
		}
		req.Params = p
	default:
		return nil, fmt.Errorf("Unsupported request type: %s", req.Type)
	}
	return &req, nil
}

func (r *NodeRequest) ToJSON() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ParseNodeResponse attempts to parse the given JSON string into a NodeResponse
// object.
func ParseNodeResponse(s string) (*NodeResponse, error) {
	var rawParams json.RawMessage
	res := NodeResponse{
		Params: &rawParams,
	}
	if err := json.Unmarshal([]byte(s), &res); err != nil {
		return nil, err
	}
	switch res.Type {
	case ResPong:
		res.Params = nil
	case ResStore:
		var p NodeResponseStoreParams
		if err := json.Unmarshal(rawParams, &p); err != nil {
			return nil, err
		}
		res.Params = p
	case ResNode:
		var p NodeResponseFindNodeParams
		if err := json.Unmarshal(rawParams, &p); err != nil {
			return nil, err
		}
		res.Params = p
	case ResValue:
		var p NodeResponseFindValueParams
		if err := json.Unmarshal(rawParams, &p); err != nil {
			return nil, err
		}
		res.Params = p
	default:
		return nil, fmt.Errorf("Unsupported response type: %s", res.Type)
	}
	return &res, nil
}

func (r *NodeResponse) ToJSON() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
