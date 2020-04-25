package main

import (
	"fmt"
	"os"

	"multini/output"
	"multini/types"
)

const (
	EXIT_SUCCESS           int = 0
	EXIT_ARG_ERR           int = 1
	EXIT_FILE_READ_ERR     int = 2
	EXIT_BAD_SYNTAX_ERR    int = 3
	EXIT_ELEMENT_NOT_FOUND int = 4
	EXIT_FILE_WRITE_ERR    int = 5
	EXIT_ACTION_ERR        int = 6
)

var (
	Version string = "development"
)

func main() {
	var err error
	var ini types.Ini

	if err = parseArgs(); err != nil {
		output.Errorln(err)
		os.Exit(EXIT_ARG_ERR)
	}

	if ArgChk {
		os.Exit(chk())
	}

	ini, err = iniRead(ArgFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(EXIT_FILE_READ_ERR)
	}

	switch {
	case ArgGet:
		err = get(&ini)
	case ArgAdd:
		err = add(&ini)
	case ArgDel:
		err = del(&ini)
	case ArgSet:
		err = set(&ini)
	}

	if err != nil {
		output.Errorln(err)
		os.Exit(EXIT_ACTION_ERR)
	}

	if ArgOutput == "-" {
		output.Println(ini.Full())
		os.Exit(EXIT_SUCCESS)
	} else if ArgOutput != "" {
		ArgFile = ArgOutput
	}

	if ArgAdd || ArgSet || ArgDel {
		if ArgInplace || ArgOutput != "" {
			err = iniWriteInplace(ArgFile, &ini)
		} else {
			err = iniWrite(ArgFile, &ini)
		}
		if err != nil {
			output.Errorln(err)
			os.Exit(EXIT_FILE_WRITE_ERR)
		}
	}

	os.Exit(EXIT_SUCCESS)
}
