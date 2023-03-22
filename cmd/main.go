package main

import (
	"fmt"
	"os"
)

func main() {
	CF, err := parseFlagsAndValidate(os.Args)
	if err != nil {
		os.Exit(64)
	}

	if err := CF.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
