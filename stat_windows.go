// +build windows

package main

import (
	"os"
)

func GetUidGid(info os.FileInfo) (int, int) {
	return -1, -1
}
