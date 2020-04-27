package types

import (
	"errors"
	"strings"
)

type Ini struct {
	Sections []Element
}

func (obj *Ini) Init() {
	obj.Sections = append(obj.Sections, &Section{}) // global section
}

func (obj *Ini) Full() string {
	var body strings.Builder
	for _, line := range obj.Sections {
		body.WriteString(line.Fulln())
	}
	return body.String()
}

func (obj *Ini) FindSection(name string) (*Section, error) {
	for i, sect := range obj.Sections {
		if sect.(*Section).Name == name {
			return obj.Sections[i].(*Section), nil
		}
	}
	return nil, errors.New("Section not found: " + name)
}

func (obj *Ini) DelSection(name string) error {
	gotIt := false
	for i, sect := range obj.Sections {
		if sect.Type() == TSection &&
			sect.(*Section).Name == name {
			obj.Sections[i] = &Deleted{}
			gotIt = true
		}
	}
	if gotIt {
		return nil
	} else if failIfNotExist() {
		return errors.New("Section not found: " + name)
	} else {
		return nil
	}
}

func (obj *Ini) Get() string {
	var sectList strings.Builder
	for _, sect := range obj.Sections {
		name := sect.(*Section).Name
		if name != "" || len(sect.(*Section).Lines) > 0 {
			sectList.WriteString(name + endOfLine)
		}
	}
	return strings.TrimSuffix(sectList.String(), endOfLine)
}

func (obj *Ini) GetSection(section string) (string, error) {
	sect, err := obj.FindSection(section)
	if err != nil {
		return "", err
	} else {
		return sect.Full(), nil
	}
}

func (obj *Ini) GetKey(section, key string) (string, error) {
	sect, err := obj.FindSection(section)
	if err != nil {
		return "", err
	} else {
		return sect.GetKey(key)
	}
}

func (obj *Ini) GetKeyVal(section, key, value string) error {
	sect, err := obj.FindSection(section)
	if err != nil {
		return err
	} else {
		return sect.GetKeyVal(key, value)
	}
}

func (obj *Ini) AddSection(section string) *Section {
	sect, err := obj.FindSection(section)
	if err != nil {
		var newSection Section
		newSection.Name = section
		newSection.Prefix = obj.Sections[len(obj.Sections)-1].Indent()
		sect = &newSection
		obj.Sections = append(obj.Sections, sect)
	}
	return sect
}

func (obj *Ini) SetSection(section string) *Section {
	return obj.AddSection(section)
}

func (obj *Ini) AddKey(section, key, value string, reverse bool) error {
	sect, err := obj.FindSection(section)
	if err != nil {
		if createIfNotExist() {
			sect = obj.AddSection(section)
		} else {
			return err
		}
	}
	sect.AddKey(key, value, reverse)
	return nil
}

func (obj *Ini) SetKey(section, key, value string) error {
	sect, err := obj.FindSection(section)
	if err != nil {
		if createIfNotExist() {
			sect = obj.AddSection(section)
		} else {
			return err
		}
	}
	return sect.SetKey(key, value)
}

func (obj *Ini) DelKey(section, key string) error {
	sect, err := obj.FindSection(section)
	if err != nil {
		if failIfNotExist() {
			return err
		} else {
			return nil
		}
	}
	return sect.DelKey(key)
}

func (obj *Ini) DelKeyVal(section, key, value string) error {
	sect, err := obj.FindSection(section)
	if err != nil {
		if failIfNotExist() {
			return err
		} else {
			return nil
		}
	}
	return sect.DelKeyVal(key, value)
}
