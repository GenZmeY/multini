package types

type Comment struct {
	Prefix string
	Value  string
}

func (obj *Comment) Full() string {
	return obj.Prefix + obj.Value
}

func (obj *Comment) Line() string {
	return obj.Prefix + obj.Value
}

func (obj *Comment) Fulln() string {
	return obj.Full() + endOfLine
}

func (obj *Comment) Type() TElement {
	return TComment
}

func (obj *Comment) Indent() string {
	return obj.Prefix
}
