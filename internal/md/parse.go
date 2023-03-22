package md

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/yaskinny/md-table-gen/internal/utils"
)

func WriteNewMD(data, path, headerName string) error {
	headerRe := regexp.MustCompile(fmt.Sprintf(`^##\s%v$`, headerName))
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	var (
		ns          = bufio.NewScanner(f)
		buf         = make([]byte, 4096)
		before      = make([]byte, 4096)
		beforeBuf   = bytes.NewBuffer(before)
		beforeFound = false
	)

	for ns.Scan() {
		buf = ns.Bytes()
		if !headerRe.Match(buf) {
			if _, err := beforeBuf.Write(buf); err != nil {
				return err
			}
			if err := beforeBuf.WriteByte('\n'); err != nil {
				return err
			}
		} else {
			beforeFound = true
			break
		}
	}

	if !beforeFound {
		d := fmt.Sprintf(utils.LeftNullTrimer(beforeBuf.String()) + data)
		if err := genNewMD(path, d); err != nil {
			return err
		}
		return nil
	}

	var (
		after          = make([]byte, 4096)
		afterBuf       = bytes.NewBuffer(after)
		normalHeaderRe = regexp.MustCompile(`^##\s`)
	)
	if beforeFound {
		extra := false

		for ns.Scan() {
			buf = ns.Bytes()
			if !extra {
				if normalHeaderRe.Match(buf) {
					extra = true
				}
			}
			if extra {
				if _, err := afterBuf.Write(buf); err != nil {
					return err
				}
				if err := afterBuf.WriteByte('\n'); err != nil {
					return err
				}
			}
		}
	}
	d := fmt.Sprintf(utils.LeftNullTrimer(beforeBuf.String()) + data + "\n" + utils.LeftNullTrimer(afterBuf.String()))
	if err := genNewMD(path, d); err != nil {
		return err
	}
	return nil
}
