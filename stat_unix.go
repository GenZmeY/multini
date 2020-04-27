// +build !windows

package main

import (
	"os"
	"syscall"
)

func GetUidGid(info os.FileInfo) (int, int) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if ok {
		return int(stat.Uid), int(stat.Gid)
	} else {
		return -1, -1
	}
}
