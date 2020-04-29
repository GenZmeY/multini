package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"multini/types"
)

// Source: https://gist.github.com/var23rav/23ae5d0d4d830aff886c3c970b8f6c6b
/*
   GoLang: os.Rename() give error "invalid cross-device link" for Docker container with Volumes.
   MoveFile(source, destination) will work moving file between folders
*/
func tryMoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return err
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return err
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}
	return nil
}

func tryRemoveRenameFile(sourcePath, destPath string) bool {
	err := os.Remove(destPath)
	if err != nil {
		return false
	}
	err = os.Rename(sourcePath, destPath)
	if err != nil {
		return false
	}
	return true
}

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

	if !tryRemoveRenameFile(newFile, realOldFile) {
		err = tryMoveFile(newFile, realOldFile)
		if err != nil {
			return err
		}
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
