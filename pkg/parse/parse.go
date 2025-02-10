package parse

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/gfkoeb/snake_case_ify/pkg/convert"
)

func ProcessFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	outputString := strings.Builder{}

	re := regexp.MustCompile(`(?i)((?:function|fn|func|def)\s+)(\w+)`)

	for scanner.Scan() {
		line := scanner.Text()
		newLine := re.ReplaceAllStringFunc(line, func(match string) string {
			groups := re.FindStringSubmatch(match)

			if len(groups) < 3 {
				return match
			}

			funcPrefix := groups[1]
			funcName := groups[2]

			newName := convert.Convert(funcName)
			return funcPrefix + newName
		})
		outputString.WriteString(newLine + "\n")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return outputString.String(), nil
}
