package types

type Trash struct {
	Value string
}

func (obj *Trash) Full() string {
	return obj.Value
}

func (obj *Trash) Line() string {
	return obj.Value
}

func (obj *Trash) Fulln() string {
	return obj.Full() + endOfLine
}

func (obj *Trash) Type() TElement {
	return TTrash
}

func (obj *Trash) Indent() string {
	return ""
}
