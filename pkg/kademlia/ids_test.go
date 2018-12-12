package kademlia

import "testing"

type nodeIDTest struct {
	nodeID    NodeID
	prefixLen int
}

func TestGetPrefixLen(t *testing.T) {
	testValues := []nodeIDTest{
		{NodeID{0x00}, 0},              // 0000 0000
		{NodeID{0x80}, 1},              // 1000 0000
		{NodeID{0xC0}, 2},              // 1100 0000
		{NodeID{0xE0}, 3},              // 1110 0000
		{NodeID{0xF0, 0x00}, 4},        // 1111 0000 0000 0000
		{NodeID{0xFF, 0x00}, 8},        // 1111 1111 0000 0000
		{NodeID{0xFF, 0xF0, 0x00}, 12}, // 1111 1111 1111 0000
	}
	for _, test := range testValues {
		prefixLen, err := test.nodeID.GetPrefixLen()
		if err != nil {
			t.Errorf("Expected non-nil error, but got: %s", err)
		}
		if prefixLen != test.prefixLen {
			t.Errorf(
				"Expected prefix length of %d, but got %d for test value: %s",
				test.prefixLen,
				prefixLen,
				test.nodeID.ToHex(),
			)
		}
	}
}
