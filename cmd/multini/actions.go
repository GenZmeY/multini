package main

import (
	"multini/internal/output"
	"multini/internal/types"
)

func chk() int {
	var ok bool
	var err error
	ok, err = iniCheck(ArgFile)
	if err != nil {
		output.Errorln(err)
		return EXIT_FILE_READ_ERR
	}
	if ok {
		return EXIT_SUCCESS
	} else {
		return EXIT_BAD_SYNTAX_ERR
	}
}

func add(ini *types.Ini) error {
	if ArgKeyIsSet {
		return ini.AddKey(ArgSection, ArgKey, ArgValue, ArgReverse)
	} else {
		ini.AddSection(ArgSection)
		return nil
	}
}

func set(ini *types.Ini) error {
	if ArgKeyIsSet {
		return ini.SetKey(ArgSection, ArgKey, ArgValue)
	} else {
		ini.SetSection(ArgSection)
		return nil
	}
}

func get(ini *types.Ini) error {
	var err error = nil
	var res string
	if ArgValueIsSet {
		return ini.GetKeyVal(ArgSection, ArgKey, ArgValue)
	} else if ArgKeyIsSet {
		res, err = ini.GetKey(ArgSection, ArgKey)
	} else if ArgSectionIsSet {
		res, err = ini.GetSection(ArgSection)
	} else {
		res = ini.Get()
	}

	if err == nil {
		output.Println(res)
	}
	return err
}

func del(ini *types.Ini) error {
	if ArgValueIsSet {
		return ini.DelKeyVal(ArgSection, ArgKey, ArgValue)
	} else if ArgKeyIsSet {
		return ini.DelKey(ArgSection, ArgKey)
	} else {
		return ini.DelSection(ArgSection)
	}
}
