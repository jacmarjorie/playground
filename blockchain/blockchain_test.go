package blockchain

import (
	"testing"
)

type blockdata struct {
	from   string
	to     string
	amount float64
}

func Test_Blockchain(t *testing.T) {
	tests := []struct {
		name     string
		data     []*blockdata
		expected bool
	}{
		{
			name:     "Test 1",
			data:     []*blockdata{{"Alice", "Bob", 5}, {"John", "Bob", 2}},
			expected: true,
		},
	}

	for _, tt := range tests {
		bc := New(2)
		for _, d := range tt.data {
			if err := bc.addBlock(d.from, d.to, d.amount); err != nil {
				t.Fatalf("test %v failed: err = %v", tt.name, err)
			}
		}
		if bc.isValid() != tt.expected {
			t.Fatalf("test %v: block chain invalid", tt.name)
		}
	}

}
