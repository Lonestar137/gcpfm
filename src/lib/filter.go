package lib

import "regexp"

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

// type model FolderStack(folderName: list)
// func StackFolderPop(folderName: string) list {}
// func StackFolderPush(folderName: string) list {}

// quick searches from stack root
//func SearchFolder(lsCommandOutput: string) string {}
