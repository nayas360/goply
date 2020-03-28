package types

type NodeType string

// casts NodeType as a string and returns value
func (n NodeType) String() string {
	return string(n)
}

// compares NodeType with a string
func (n NodeType) Equals(o string) bool {
	return string(n) == o
}

const (
	NumberT NodeType = "<number>"
	AddOpT  NodeType = "<add_op>"
	SubOpT  NodeType = "<sub_op>"
	MulOpT  NodeType = "<mul_op>"
	DivOpT  NodeType = "<div_op>"
	LParenT NodeType = "<l_paren>"
	RParenT NodeType = "<r_paren>"
)
