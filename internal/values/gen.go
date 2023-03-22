package values

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yaskinny/md-table-gen/internal/utils"
)

func RenderValueFile(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()
	ns := bufio.NewScanner(file)
	output := strings.Builder{}
	if _, err := output.WriteString(fmt.Sprintf("### %v\n%v", filepath.Base(path), tableHeader)); err != nil {
		return "", err
	}

	objs := make(map[string]string)

	for ns.Scan() {
		data := ns.Bytes()
		switch {
		case isSection(data):
			if _, err := output.WriteString(utils.LeftNullTrimer(string(renderSection(data)))); err != nil {
				return "", err
			}
			if _, err := output.WriteString(tableHeader); err != nil {
				return "", err
			}

		case isOpt(data):
			if _, err := output.WriteString(utils.LeftNullTrimer(string(renderOpt(data)))); err != nil {
				return "", err
			}

		case isObjOpt(data):
			// [0] objName
			// [1] objOpt
			// [2] desc
			s := renderObjOpt(data)
			ss := strings.SplitAfterN(string(s), " ", 3)

			objName := utils.LeftNullTrimer(ss[0])
			objName = strings.Trim(objName, " ")
			objOpt := utils.LeftNullTrimer(ss[1])
			objOpt = strings.Trim(objOpt, " ")
			objDesc := utils.LeftNullTrimer(ss[2])
			objDesc = strings.Trim(objDesc, " ")

			if _, ok := objs[objName]; ok {
				objs[objName] += fmt.Sprintf("| %v | `%v` |\n", objOpt, objDesc)
			} else {
				return "", fmt.Errorf("option for object %[1]v is defined but object %[1]v can not be found", objName)
			}

		case isObjName(data):
			// [0] objName
			// [1] desc
			s := renderObjName(data)
			ss := strings.SplitAfterN(string(s), " ", 2)

			objName := utils.LeftNullTrimer(ss[0])
			objName = strings.Trim(objName, " ")
			objDesc := utils.LeftNullTrimer(ss[1])
			objDesc = strings.Trim(objDesc, " ")

			if _, ok := objs[objName]; ok {
				return "", fmt.Errorf("object %v is defined multiple times in file %v", objName, path)
			} else {
				objs[objName] += fmt.Sprintf("\n*%v*: %v\n%v", objName, objDesc, tableHeader)
			}

		case isMandObjName(data):
			// [0] objName
			// [1] desc
			s := renderMandObjName(data)
			ss := strings.SplitAfterN(string(s), " ", 2)

			objName := utils.LeftNullTrimer(ss[0])
			objName = strings.Trim(objName, " ")
			objDesc := utils.LeftNullTrimer(ss[1])
			objDesc = strings.Trim(objDesc, " ")

			if _, ok := objs[objName]; ok {
				return "", fmt.Errorf("object %v is defined multiple times in file %v", objName, path)
			} else {
				objs[objName] += fmt.Sprintf("\n***%v***: %v\n%v", objName, objDesc, tableHeader)
			}

		case isMandOpt(data):
			if _, err := output.WriteString(utils.LeftNullTrimer(string(renderMandOpt(data)))); err != nil {
				return "", err
			}

		case isMandObjOpt(data):
			// [0] objName
			// [1] objOpt
			// [2] desc
			s := renderMandObjOpt(data)
			ss := strings.SplitAfterN(string(s), " ", 3)

			objName := utils.LeftNullTrimer(ss[0])
			objName = strings.Trim(objName, " ")
			objOpt := utils.LeftNullTrimer(ss[1])
			objOpt = strings.Trim(objOpt, " ")
			objDesc := utils.LeftNullTrimer(ss[2])
			objDesc = strings.Trim(objDesc, " ")

			if _, ok := objs[objName]; ok {
				objs[objName] += fmt.Sprintf("| **%v** | `%v` |\n", objOpt, objDesc)
			} else {
				return "", fmt.Errorf("option for object %[1]v is defined but object %[1]v can not be found", objName)
			}
		}
	}
	for _, v := range objs {
		if _, err := output.WriteString(v + "\n"); err != nil {
			return "", err
		}
	}
	return output.String(), nil
}
