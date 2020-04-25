package types

type EmptyLine struct {
	Prefix string
}

func (obj *EmptyLine) Full() string {
	return obj.Prefix
}

func (obj *EmptyLine) Line() string {
	return obj.Prefix
}

func (obj *EmptyLine) Fulln() string {
	return obj.Full() + endOfLine
}

func (obj *EmptyLine) Type() TElement {
	return TEmptyLine
}

func (obj *EmptyLine) Indent() string {
	return obj.Prefix
}
