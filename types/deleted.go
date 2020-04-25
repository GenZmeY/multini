package types

type Deleted struct {
}

func (obj *Deleted) Full() string {
	return ""
}

func (obj *Deleted) Line() string {
	return ""
}

func (obj *Deleted) Fulln() string {
	return ""
}

func (obj *Deleted) Type() TElement {
	return TDeleted
}

func (obj *Deleted) Indent() string {
	return ""
}
