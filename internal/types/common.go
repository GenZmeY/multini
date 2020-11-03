package types

import (
	"runtime"
)

type TElement int

const (
	TDeleted   TElement = 0
	TComment   TElement = 1
	TEmptyLine TElement = 2
	TKeyValue  TElement = 3
	TSection   TElement = 4
	TTrash     TElement = 5
)

var (
	endOfLine string = "\n"
	existing  bool   = false
)

func SetEndOfLineNative() {
	switch os := runtime.GOOS; os {
	case "windows":
		SetEndOfLineWindows()
	default:
		SetEndOfLineUnix()
	}
}

func SetEndOfLineUnix() {
	endOfLine = "\n"
}

func SetEndOfLineWindows() {
	endOfLine = "\r\n"
}

func SetExistingMode(value bool) {
	existing = value
}

func createIfNotExist() bool {
	return !existing
}

func failIfNotExist() bool {
	return existing
}
