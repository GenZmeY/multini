package types

import (
	"errors"
	"strings"
)

type Section struct {
	Prefix  string
	Name    string
	Postfix string
	Comment Comment
	Lines   []Element
}

func (obj *Section) Headern() string {
	if obj.Name == "" {
		return ""
	} else {
		return obj.Prefix + "[" + obj.Name + "]" + obj.Postfix + obj.Comment.Full() + endOfLine
	}
}

func (obj *Section) Line() string {
	return obj.Header()
}

func (obj *Section) Header() string {
	return strings.TrimSuffix(obj.Headern(), endOfLine)
}

func (obj *Section) Fulln() string {
	var body strings.Builder
	for _, line := range obj.Lines {
		body.WriteString(line.Fulln())
	}
	return obj.Headern() + body.String()
}

func (obj *Section) Full() string {
	return strings.TrimSuffix(obj.Fulln(), endOfLine)
}

func (obj *Section) Type() TElement {
	return TSection
}

func (obj *Section) Indent() string {
	return obj.Prefix
}

func (obj *Section) DelKey(name string) error {
	gotIt := false
	for i, keyVal := range obj.Lines {
		if keyVal.Type() == TKeyValue &&
			keyVal.(*KeyValue).Key == name {
			obj.Lines[i] = &Deleted{}
			gotIt = true
		}
	}
	if gotIt {
		return nil
	} else if failIfNotExist() {
		return errors.New("Parameter not found: " + name)
	} else {
		return nil
	}
}

func (obj *Section) DelKeyVal(name, value string) error {
	gotIt := false
	for i, keyVal := range obj.Lines {
		if keyVal.Type() == TKeyValue &&
			keyVal.(*KeyValue).Key == name &&
			keyVal.(*KeyValue).Value == value {
			obj.Lines[i] = &Deleted{}
			gotIt = true
		}
	}
	if gotIt {
		return nil
	} else if failIfNotExist() {
		return errors.New("Parameter:value pair not found: " + name + ":" + value)
	} else {
		return nil
	}
}

func (obj *Section) GetKey(name string) (string, error) {
	var err error = nil
	var result strings.Builder
	for _, keyVal := range obj.Lines {
		if keyVal.Type() == TKeyValue && keyVal.(*KeyValue).Key == name {
			result.WriteString(keyVal.(*KeyValue).Value + endOfLine)
		}
	}
	if result.String() == "" {
		err = errors.New("Parameter not found: " + name)
	}
	return strings.TrimSuffix(result.String(), endOfLine), err
}

func (obj *Section) GetKeyVal(name, value string) error {
	for _, keyVal := range obj.Lines {
		if keyVal.Type() == TKeyValue &&
			keyVal.(*KeyValue).Key == name &&
			keyVal.(*KeyValue).Value == value {
			return nil
		}
	}
	return errors.New("Parameter:Value not found: " + name + ":" + value)
}

func (obj *Section) appendKey(name, value string) {
	var newKeyValue KeyValue
	var replaceIndex int = -1
	newKeyValue.Key = name
	newKeyValue.Value = value
	// replace first emptyline
	for i := len(obj.Lines) - 1; i >= 0; i-- {
		if obj.Lines[i].Type() == TEmptyLine {
			replaceIndex = i
		} else {
			break
		}
	}
	// for right indent and tabs
	for i := len(obj.Lines) - 1; i >= 0; i-- {
		if obj.Lines[i].Type() == TKeyValue {
			template := obj.Lines[i].(*KeyValue)
			newKeyValue.PrefixKey = template.PrefixKey
			newKeyValue.PostfixKey = template.PostfixKey
			newKeyValue.PrefixValue = template.PrefixValue
			newKeyValue.PostfixValue = template.PostfixValue
			break
		}
	}
	if replaceIndex == -1 {
		obj.Lines = append(obj.Lines, &newKeyValue)
	} else {
		obj.Lines = append(obj.Lines, obj.Lines[replaceIndex])
		obj.Lines[replaceIndex] = &newKeyValue
	}
}

func (obj *Section) AddKey(name, value string) {
	gotIt := false
	for i, keyVal := range obj.Lines {
		if keyVal.Type() == TKeyValue &&
			keyVal.(*KeyValue).Key == name &&
			keyVal.(*KeyValue).Value == value {
			if !gotIt {
				gotIt = true
			} else {
				obj.Lines[i] = &Deleted{}
			}
		}
	}
	if !gotIt {
		obj.appendKey(name, value)
	}
}

func (obj *Section) SetKey(name, value string) error {
	gotIt := false
	for i, keyVal := range obj.Lines {
		if keyVal.Type() == TKeyValue &&
			keyVal.(*KeyValue).Key == name {
			if !gotIt {
				keyVal.(*KeyValue).Value = value
				gotIt = true
			} else {
				obj.Lines[i] = &Deleted{}
			}
		}
	}
	if !gotIt {
		if createIfNotExist() {
			obj.appendKey(name, value)
		} else {
			return errors.New("Parameter not found: " + name)
		}
	}
	return nil
}
