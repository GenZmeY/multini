package main

import (
	"bufio"
	"io/ioutil"
	"os"

	"multini/types"
)

func replaceOriginal(oldFile, newFile string) error {
	err := os.Remove(oldFile)
	if err == nil {
		err = os.Rename(newFile, oldFile)
	}
	return err
}

func iniWrite(filename string, ini *types.Ini) error {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "multini")
	if err == nil {
		datawriter := bufio.NewWriter(tmpFile)
		_, err = datawriter.WriteString(ini.Full())
		if err == nil {
			err = datawriter.Flush()
			tmpFile.Close()
			if err == nil {
				err = replaceOriginal(filename, tmpFile.Name())
			}
		}
	}
	return err
}

func iniWriteInplace(filename string, ini *types.Ini) error {
	targetFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err == nil {
		datawriter := bufio.NewWriter(targetFile)
		_, err = datawriter.WriteString(ini.Full())
		if err == nil {
			err = datawriter.Flush()
			targetFile.Close()
		}
	}
	return err
}
