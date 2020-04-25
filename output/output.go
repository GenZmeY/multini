package output

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

var (
	endOfLine string      = "\n"
	devNull   *log.Logger = log.New(ioutil.Discard, "", 0)
	stdout    *log.Logger = log.New(os.Stdout, "", 0)
	stderr    *log.Logger = log.New(os.Stderr, "", 0)
	verbose   *log.Logger = devNull
)

func SetVerbose(enabled bool) {
	if enabled {
		verbose = stderr
	} else {
		verbose = devNull
	}
}

func SetEndOfLineNative() {
	switch os := runtime.GOOS; os {
	case "windows":
		setEndOfLineWindows()
	default:
		setEndOfLineUnix()
	}
}

func EOL() string {
	return endOfLine
}

func setEndOfLineUnix() {
	endOfLine = "\n"
}

func setEndOfLineWindows() {
	endOfLine = "\r\n"
}

func Print(v ...interface{}) {
	stdout.Print(v...)
}

func Printf(format string, v ...interface{}) {
	stdout.Printf(format, v...)
}

func Println(v ...interface{}) {
	stdout.Print(fmt.Sprint(v...) + endOfLine)
}

func Error(v ...interface{}) {
	stderr.Print(v...)
}

func Errorf(format string, v ...interface{}) {
	stderr.Printf(format, v...)
}

func Errorln(v ...interface{}) {
	stderr.Print(fmt.Sprint(v...) + endOfLine)
}

func Verbose(v ...interface{}) {
	verbose.Print(v...)
}

func Verbosef(format string, v ...interface{}) {
	verbose.Printf(format, v...)
}

func Verboseln(v ...interface{}) {
	verbose.Print(fmt.Sprint(v...) + endOfLine)
}
