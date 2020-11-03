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
	NgPrefix        string = `prefix`
	NgPostifx       string = `postfix`
	NgSection       string = `section`
	NgKey           string = `key`
	NgKeyPostfix    string = `key_postfix`
	NgValue         string = `value`
	NgValuePrefix   string = `value_prefix`
	NgValuePostfix  string = `value_postfix`
	NgComment       string = `comment`
	NgCommentPrefix string = `comment_prefix`

	RxBodyPrefix    string         = `(?P<` + NgPrefix + `>\s+)?`
	RxSectionName   string         = `\[(?P<` + NgSection + `>.+)\]`
	RxKey           string         = `(?P<` + NgKey + `>(?:[^;#/=]+[^\s=;#/]|[^;#/=]))?`
	RxKeyPostfix    string         = `(?P<` + NgKeyPostfix + `>\s+)?`
	RxValuePrefix   string         = `(?P<` + NgValuePrefix + `>\s+)?`
	RxValue         string         = `(?P<` + NgValue + `>(?:[^;#/]+[^\s;#/]|[^;#/]))?`
	RxValuePostfix  string         = `(?P<` + NgValuePostfix + `>\s+)?`
	RxKeyVal        string         = RxKey + RxKeyPostfix + `=` + RxValuePrefix + RxValue + RxValuePostfix
	RxBody          string         = `(?:` + RxSectionName + `|` + RxKeyVal + `)?`
	RxBodyPostfix   string         = `(?P<` + NgPostifx + `>\s+)?`
	RxCommentPrefix string         = `(?P<` + NgCommentPrefix + `>([#;]|//)\s*)`
	RxCommentText   string         = `(?P<` + NgComment + `>.+)?`
	RxComment       string         = `(?:` + RxCommentPrefix + RxCommentText + `)?`
	Rx              string         = RxBodyPrefix + RxBody + RxBodyPostfix + RxComment
	RxCompiled      *regexp.Regexp = regexp.MustCompile(Rx)
)

func rxParse(rx *regexp.Regexp, str string) map[string]string {
	match := rx.FindStringSubmatch(str)
	result := make(map[string]string)
	for i, name := range rx.SubexpNames() {
		if i != 0 && name != "" && i <= len(match) {
			result[name] = match[i]
		}
	}
	return result
}

func debugMap(el map[string]string) string {
	var dbgMap strings.Builder
	for key, val := range el {
		dbgMap.WriteString("  " + key + ": \"" + val + "\"" + output.EOL())
	}
	return dbgMap.String()
}

func appendLine(ini *types.Ini, line string) error {
	elements := rxParse(RxCompiled, line)
	switch {
	case elements[NgSection] != "":
		var newSection types.Section
		newSection.Name = elements[NgSection]
		newSection.Prefix = elements[NgPrefix]
		newSection.Postfix = elements[NgPostifx]
		newSection.Comment.Prefix = elements[NgCommentPrefix]
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
		newKeyValue.Comment.Prefix = elements[NgCommentPrefix]
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
		newComment.Prefix = elements[NgPrefix] + elements[NgCommentPrefix]
		newComment.Value = elements[NgComment]
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
