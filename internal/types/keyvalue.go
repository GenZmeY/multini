package types

type KeyValue struct {
	PrefixKey    string
	Key          string
	PostfixKey   string
	PrefixValue  string
	Value        string
	PostfixValue string
	Comment      Comment
}

func (obj *KeyValue) Full() string {
	return obj.PrefixKey +
		obj.Key +
		obj.PostfixKey +
		"=" +
		obj.PrefixValue +
		obj.Value +
		obj.PostfixValue +
		obj.Comment.Full()
}

func (obj *KeyValue) Line() string {
	return obj.Full()
}

func (obj *KeyValue) Fulln() string {
	return obj.Full() + endOfLine
}

func (obj *KeyValue) Type() TElement {
	return TKeyValue
}

func (obj *KeyValue) Indent() string {
	return obj.PrefixKey
}
