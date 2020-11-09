package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"

	"multini/internal/output"
	"multini/internal/types"
)

var (
	// Ng - Named Group
	NgPrefix       string = `prefix`
	NgPostifx      string = `postfix`
	NgSection      string = `section`
	NgKey          string = `key`
	NgKeyPostfix   string = `key_postfix`
	NgValue        string = `value`
	NgValuePrefix  string = `value_prefix`
	NgValuePostfix string = `value_postfix`
	NgComment      string = `comment`
	NgData         string = `data`

	RxEmpty   string = `^(?P<` + NgPrefix + `>\s+)?$`
	RxSection string = `^(?P<` + NgPrefix + `>\s+)?\[(?P<` + NgSection + `>[^\]]+)\](?P<` + NgPostifx + `>\s+)?$`
	RxKey     string = `^(?P<` + NgPrefix + `>\s+)?(?P<` + NgKey + `>.*[^\s]+)(?P<` + NgKeyPostfix + `>\s+)?$`
	RxValue   string = `^(?P<` + NgValuePrefix + `>\s+)?(?P<` + NgValue + `>.*[^\s])(?P<` + NgValuePostfix + `>\s+)?$`

	RxEmptyCompile   *regexp.Regexp = regexp.MustCompile(RxEmpty)
	RxSectionCompile *regexp.Regexp = regexp.MustCompile(RxSection)
	RxKeyCompile     *regexp.Regexp = regexp.MustCompile(RxKey)
	RxValueCompile   *regexp.Regexp = regexp.MustCompile(RxValue)
)

func parse(str string) map[string]string {
	var result map[string]string = make(map[string]string)
	var data string

	data, result[NgComment] = getDataComment(str)

	if data != "" {
		findNamedGroups(data, RxEmptyCompile, &result)
	}

	if result[NgPrefix] != "" {
		return result
	}

	findNamedGroups(data, RxSectionCompile, &result)

	if result[NgSection] == "" && data != "" {
		keyPart, valPart := getKeyValue(data)
		findNamedGroups(keyPart, RxKeyCompile, &result)
		findNamedGroups(valPart, RxValueCompile, &result)
	}

	return result
}

func findNamedGroups(str string, Rx *regexp.Regexp, result *map[string]string) {
	match := Rx.FindStringSubmatch(str)
	for i, name := range Rx.SubexpNames() {
		if i != 0 && name != "" && i <= len(match) {
			(*result)[name] = match[i]
		}
	}
}

func getDataComment(str string) (string, string) {
	var indexes []int
	var commentIndex int = -1

	indexes = append(indexes, strings.Index(str, "//"))
	indexes = append(indexes, strings.Index(str, "#"))
	indexes = append(indexes, strings.Index(str, ";"))

	for _, index := range indexes {
		if commentIndex == -1 {
			if index != -1 {
				commentIndex = index
			}
		} else {
			if index != -1 {
				if commentIndex > index {
					commentIndex = index
				}
			}
		}
	}

	if commentIndex == -1 {
		return str, ""
	} else {
		return str[:commentIndex], str[commentIndex:]
	}
}

func getKeyValue(data string) (string, string) {
	index := strings.Index(data, "=")
	if index != -1 {
		return data[:index], data[index+1:]
	}
	return "", ""
}

func debugMap(el map[string]string) string {
	var dbgMap strings.Builder
	for key, val := range el {
		dbgMap.WriteString("  " + key + ": \"" + val + "\"" + output.EOL())
	}
	return dbgMap.String()
}

func appendLine(ini *types.Ini, line string) error {
	// elements := rxParse(line)
	elements := parse(line)
	switch {
	case elements[NgSection] != "":
		var newSection types.Section
		newSection.Name = elements[NgSection]
		newSection.Prefix = elements[NgPrefix]
		newSection.Postfix = elements[NgPostifx]
		newSection.Comment.Value = elements[NgComment]
		if newSection.Line() == line {
			ini.Sections = append(ini.Sections, &newSection)
			return nil
		} else {
			output.Verboseln("Got:", newSection.Line())
			var newTrash types.Trash = types.Trash{Value: line}
			ini.Sections[len(ini.Sections)-1].(*types.Section).Lines = append(ini.Sections[len(ini.Sections)-1].(*types.Section).Lines, &newTrash)
		}
	case elements[NgKey] != "":
		var newKeyValue types.KeyValue
		newKeyValue.Key = elements[NgKey]
		newKeyValue.PostfixKey = elements[NgKeyPostfix]
		newKeyValue.PrefixKey = elements[NgPrefix]
		newKeyValue.Value = elements[NgValue]
		newKeyValue.PrefixValue = elements[NgValuePrefix]
		newKeyValue.PostfixValue = elements[NgValuePostfix]
		newKeyValue.Comment.Value = elements[NgComment]
		if newKeyValue.Line() == line {
			ini.Sections[len(ini.Sections)-1].(*types.Section).Lines = append(ini.Sections[len(ini.Sections)-1].(*types.Section).Lines, &newKeyValue)
			return nil
		} else {
			output.Verboseln("Got:", newKeyValue.Line())
			var newTrash types.Trash = types.Trash{Value: line}
			ini.Sections[len(ini.Sections)-1].(*types.Section).Lines = append(ini.Sections[len(ini.Sections)-1].(*types.Section).Lines, &newTrash)
		}
	case elements[NgComment] != "":
		var newComment types.Comment
		newComment.Value = elements[NgComment]
		newComment.Prefix = elements[NgPrefix]
		if newComment.Line() == line {
			ini.Sections[len(ini.Sections)-1].(*types.Section).Lines = append(ini.Sections[len(ini.Sections)-1].(*types.Section).Lines, &newComment)
			return nil
		} else {
			output.Verboseln("Got:", newComment.Line())
			var newTrash types.Trash = types.Trash{Value: line}
			ini.Sections[len(ini.Sections)-1].(*types.Section).Lines = append(ini.Sections[len(ini.Sections)-1].(*types.Section).Lines, &newTrash)
		}
	case elements[NgPrefix] != "" || line == "":
		var newEmptyLine types.EmptyLine
		newEmptyLine.Prefix = elements[NgPrefix]
		if newEmptyLine.Line() == line {
			ini.Sections[len(ini.Sections)-1].(*types.Section).Lines = append(ini.Sections[len(ini.Sections)-1].(*types.Section).Lines, &newEmptyLine)
			return nil
		} else {
			output.Verboseln("Got:", newEmptyLine.Line())
			var newTrash types.Trash = types.Trash{Value: line}
			ini.Sections[len(ini.Sections)-1].(*types.Section).Lines = append(ini.Sections[len(ini.Sections)-1].(*types.Section).Lines, &newTrash)
		}
	default:
		var newTrash types.Trash = types.Trash{Value: line}
		ini.Sections[len(ini.Sections)-1].(*types.Section).Lines = append(ini.Sections[len(ini.Sections)-1].(*types.Section).Lines, &newTrash)
	}
	return errors.New(debugMap(elements))
}

func iniRead(filename string) (types.Ini, error) {
	var (
		err     error
		iniFile *os.File
		ini     types.Ini
	)
	iniFile, err = os.Open(filename)

	if err != nil {
		return ini, err
	}

	fileScanner := bufio.NewScanner(iniFile)
	fileScanner.Split(bufio.ScanLines)
	ini.Init()

	for i := 1; fileScanner.Scan(); i++ {
		appendLine(&ini, fileScanner.Text())
	}

	iniFile.Close()
	return ini, nil
}

func iniCheck(filename string) (bool, error) {
	var (
		err     error
		iniFile *os.File
		ini     types.Ini
		ok      bool = true
	)
	iniFile, err = os.Open(filename)

	if err != nil {
		return ok, err
	}

	fileScanner := bufio.NewScanner(iniFile)
	fileScanner.Split(bufio.ScanLines)
	ini.Init()

	for i := 1; fileScanner.Scan(); i++ {
		line := fileScanner.Text()
		err = appendLine(&ini, line)
		if err != nil {
			output.Errorln(i, ":", line)
			output.Verboseln(err)
			ok = false
		}
	}

	iniFile.Close()
	return ok, nil
}
