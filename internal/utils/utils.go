package utils

import (
	"path/filepath"
	"strings"
)

func LeftNullTrimer(s string) string {
	return strings.TrimLeft(s, "\x00")
}

func AppendDirName(p string) string {
	dn := filepath.Dir(p)
	if dn == "." {
		return dn
	}
	dns := strings.Split(dn, "/")

	// if it is inside of a directory, for instance dirName/fileName return 'dirName'
	// if it is inside of multiple directories just return last two directory names
	// for instance dir1/dir2/dir3/fileName, return 'dir2/dir3'
	if len(dns) >= 2 {
		dn = filepath.Join(dns[len(dns)-2:]...)
	}
	return dn
}
