package node

type Operator interface {
	Expression
	Operator()
}
