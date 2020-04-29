package main

import (
	"errors"
	"os"

	"multini/output"
	"multini/types"

	"github.com/juju/gnuflag"
)

var (
	ArgInplace  bool
	ArgGet      bool
	ArgSet      bool
	ArgAdd      bool
	ArgDel      bool
	ArgChk      bool
	ArgVerbose  bool
	ArgVersion  bool
	ArgWindows  bool
	ArgUnix     bool
	ArgHelp     bool
	ArgExisting bool
	ArgReverse  bool
	ArgQuiet    bool
	ArgOutput   string

	ArgFile         string
	ArgFileIsSet    bool = false
	ArgSection      string
	ArgSectionIsSet bool = false
	ArgKey          string
	ArgKeyIsSet     bool = false
	ArgValue        string
	ArgValueIsSet   bool = false
)

func printHelp() {
	output.Println("A utility for manipulating ini files with duplicate keys")
	output.Println("")
	output.Println("Usage: multini [OPTION]... [ACTION] config_file [section] [param] [value]")
	output.Println("Actions:")
	output.Println("  -g, --get          Get values for a given combination of parameters.")
	output.Println("  -s, --set          Set values for a given combination of parameters.")
	output.Println("  -a, --add          Add values for a given combination of parameters.")
	output.Println("  -d, --del          Delete the given combination of parameters.")
	output.Println("  -c, --chk          Display parsing errors for the specified file.")
	output.Println("")
	output.Println("Options:")
	output.Println("  -e, --existing     For --set and --del, fail if item is missing.")
	output.Println("  -r, --reverse      For --add, adds an item to the top of the section")
	output.Println("  -i, --inplace      Lock and write files in place.")
	output.Println("                     This is not atomic but has less restrictions")
	output.Println("                     than the default replacement method.")
	output.Println("  -o, --output FILE  Write output to FILE instead. '-' means stdout")
	//	output.Println("  -v, --verbose      Indicate on stderr if changes were made")
	output.Println("  -u, --unix         Use LF as end of line")
	output.Println("  -w, --windows      Use CRLF as end of line")
	output.Println("  -q, --quiet        Suppress all normal output")
	output.Println("  -h, --help         Write this help to stdout")
	output.Println("      --version      Write version to stdout")
}

func printVersion() {
	output.Println("multini ", Version)
}

func init() {
	gnuflag.BoolVar(&ArgGet, "get", false, "")
	gnuflag.BoolVar(&ArgGet, "g", false, "")
	gnuflag.BoolVar(&ArgAdd, "add", false, "")
	gnuflag.BoolVar(&ArgAdd, "a", false, "")
	gnuflag.BoolVar(&ArgSet, "set", false, "")
	gnuflag.BoolVar(&ArgSet, "s", false, "")
	gnuflag.BoolVar(&ArgDel, "del", false, "")
	gnuflag.BoolVar(&ArgDel, "d", false, "")
	gnuflag.BoolVar(&ArgChk, "chk", false, "")
	gnuflag.BoolVar(&ArgChk, "c", false, "")
	gnuflag.BoolVar(&ArgInplace, "inplace", false, "")
	gnuflag.BoolVar(&ArgInplace, "i", false, "")
	gnuflag.BoolVar(&ArgUnix, "unix", false, "")
	gnuflag.BoolVar(&ArgUnix, "u", false, "")
	gnuflag.BoolVar(&ArgWindows, "windows", false, "")
	gnuflag.BoolVar(&ArgWindows, "w", false, "")
	gnuflag.BoolVar(&ArgReverse, "reverse", false, "")
	gnuflag.BoolVar(&ArgReverse, "r", false, "")
	gnuflag.BoolVar(&ArgExisting, "existing", false, "")
	gnuflag.BoolVar(&ArgExisting, "e", false, "")
	gnuflag.BoolVar(&ArgQuiet, "quiet", false, "")
	gnuflag.BoolVar(&ArgQuiet, "q", false, "")
	gnuflag.BoolVar(&ArgVerbose, "verbose", false, "")
	gnuflag.BoolVar(&ArgVerbose, "v", false, "")
	gnuflag.StringVar(&ArgOutput, "output", "", "")
	gnuflag.StringVar(&ArgOutput, "o", "", "")
	gnuflag.BoolVar(&ArgVersion, "version", false, "")
	gnuflag.BoolVar(&ArgHelp, "help", false, "")
	gnuflag.BoolVar(&ArgHelp, "h", false, "")
}

func parseArgs() error {
	gnuflag.Parse(false)

	// info
	switch {
	case ArgHelp:
		printHelp()
		os.Exit(EXIT_SUCCESS)
	case ArgVersion:
		printVersion()
		os.Exit(EXIT_SUCCESS)
	}

	// File EOF
	switch {
	case ArgWindows:
		types.SetEndOfLineWindows()
	case ArgUnix:
		types.SetEndOfLineUnix()
	default:
		types.SetEndOfLineNative()
	}

	// Output settings
	output.SetEndOfLineNative()
	output.SetVerbose(ArgVerbose)
	output.SetQuiet(ArgQuiet)

	// Positional Args
	for i := 0; i < 4 && i < gnuflag.NArg(); i++ {
		switch i {
		case 0:
			ArgFile = gnuflag.Arg(0)
			ArgFileIsSet = true
		case 1:
			ArgSection = gnuflag.Arg(1)
			ArgSectionIsSet = true
		case 2:
			ArgKey = gnuflag.Arg(2)
			ArgKeyIsSet = true
		case 3:
			ArgValue = gnuflag.Arg(3)
			ArgValueIsSet = true
		}
	}

	if !ArgFileIsSet {
		return errors.New("Config_file not specified")
	}

	// Mode
	actionCounter := 0
	if ArgChk {
		actionCounter++
	}
	if ArgDel {
		actionCounter++
		if !ArgSectionIsSet {
			return errors.New("Section not specified")
		}
	}
	if ArgGet {
		actionCounter++
	}
	if ArgSet {
		actionCounter++
		if !ArgSectionIsSet {
			return errors.New("Section not specified")
		}
	}
	if ArgAdd {
		actionCounter++
		if !ArgSectionIsSet {
			return errors.New("Section not specified")
		}
	}
	switch actionCounter {
	case 0:
		return errors.New("Action not set")
	case 1:
		return nil
	default:
		return errors.New("Only one action can be used at the same time")
	}
}
