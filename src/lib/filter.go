package lib

import (
	"regexp"
	"strings"
)

func Filter(regexToMatchAgainst string, lsCommandOutput []string) []string {
	// NOTES: isn't it sorted by default? Maybe use a better search alg

	var folderRegex, _ = regexp.Compile(regexToMatchAgainst)

	var filteredList []string

	for i := 1; i < len(lsCommandOutput); i++ {
		var folderPath string = lsCommandOutput[i]
		var match bool = folderRegex.MatchString(folderPath)

		if match {
			filteredList = append(filteredList, folderPath)

		}
	}

	return filteredList
}

func ParseCmdOutput(cmdOutput []byte) []string {
	var cmdOutputAsString string = string(cmdOutput)
	var parsedCmdOutput []string

	parsedCmdOutput = strings.Split(cmdOutputAsString, "\n")

	return parsedCmdOutput
}

// quick searches from stack root
//func SearchFolder(lsCommandOutput: string) string {}
