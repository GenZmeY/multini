package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"

	"multini/types"
)

func replaceOriginal(oldFile, newFile string) error {
	realOldFile, err := filepath.EvalSymlinks(oldFile)
	if err != nil {
		return err
	}

	infoOldFile, err := os.Stat(realOldFile)
	if err != nil {
		return err
	}
	mode := infoOldFile.Mode()

	var uid, gid int = GetUidGid(infoOldFile)

	err = os.Remove(realOldFile)
	if err != nil {
		return err
	}

	err = os.Rename(newFile, realOldFile)
	if err != nil {
		return err
	}

	err = os.Chmod(realOldFile, mode)
	if err != nil {
		return err
	}

	// try to restore original uid/gid
	// don't worry if we can't
	os.Chown(realOldFile, uid, gid)

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
	realfilename, err := filepath.EvalSymlinks(filename)
	mode := os.FileMode(int(0644))
	if os.IsNotExist(err) {
		realfilename = filename
	} else if err != nil {
		return err
	} else {
		info, err := os.Stat(realfilename)
		if err != nil {
			return err
		}
		mode = info.Mode()
	}
	targetFile, err := os.OpenFile(realfilename, os.O_WRONLY|os.O_CREATE, mode)
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
