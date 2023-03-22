package md

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func genTempFileName(path string) string {
	basePath := filepath.Base(path)
	basePath = fmt.Sprintf("%v.%d", basePath, time.Now().Unix())
	return filepath.Join(os.TempDir(), basePath)
}

// backup md file to a temporary file in default os temp directory
func bak(op, np string) error {
	opf, err := os.Open(op)
	if err != nil {
		return err
	}
	defer opf.Close()

	npf, err := os.Create(np)
	if err != nil {
		return err
	}
	defer npf.Close()

	if _, err := io.Copy(npf, opf); err != nil {
		return err
	}
	return nil
}

func genNewMD(path, data string) error {
	bakPath := genTempFileName(path)
	if err := bak(path, bakPath); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// write data to md file and if something went wrong
	if _, err := fmt.Fprintf(f, "%v", data); err != nil {
		// first open the backup file
		bf, err := os.Open(bakPath)
		if err != nil {
			panic(err)
		}
		defer bf.Close()
		// close original md file
		f.Close()
		// remove the original md file
		if err := os.Remove(path); err != nil {
			panic(err)
		}
		// create a new md file
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
		// copy content of backup file to the new file
		if _, err := io.Copy(f, bf); err != nil {
			panic(err)
		}
		return fmt.Errorf("something went wrong while writing new version of md file. err=%v", err)

	} else {
		if err := f.Sync(); err != nil {
			return err
			// should we restore the backup file if something went wrong with sync?
		}
		f.Close()
		return nil
	}
}
