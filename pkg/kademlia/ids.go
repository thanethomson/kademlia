package kademlia

import (
	"encoding/hex"
	"fmt"
)

// MaxPrefixLen represents the maximum number of bits in a prefix.
const MaxPrefixLen int = 256

// MaxPrefixLenBytes represents the maximum number of bytes in a prefix.
const MaxPrefixLenBytes int = MaxPrefixLen / 8

// NodeID represents a node's ID in big endian notation.
type NodeID []byte

// ToHex converts the node ID to its hexadecimal string representation.
func (id *NodeID) ToHex() string {
	return hex.EncodeToString(*id)
}

// GetPrefixLen returns a string containing the number of 1's in this node ID's
// prefix.
func (id *NodeID) GetPrefixLen() (int, error) {
	idLen := len(*id)
	if idLen < 1 || idLen > MaxPrefixLenBytes {
		return 0, fmt.Errorf("Invalid length for NodeID: %d", idLen)
	}
	for i, b := range *id {
		for j := 7; j >= 0; j-- {
			bit := (b >> uint(j)) & 0x01
			if bit == 0 {
				return (i * 8) + (7 - j), nil
			}
		}
	}
	return idLen * 8, nil
}
