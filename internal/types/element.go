package types

type Element interface {
	Full() string
	Fulln() string
	Line() string
	Indent() string
	Type() TElement
}
