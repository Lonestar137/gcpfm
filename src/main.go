package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/lonestar137/gcpfm/src/lib"
	"github.com/lonestar137/gcpfm/src/ui"
)

// Goal: make GCP ranger like filemanager, objective is to make GCP faster to navigate and use.
// You can press y on any dir/file and it will close the program and print the path to console.

func ShowAvailBuckets() string {
	//var cmd = "gsutil ls"
	//availBucketsOutput, err := exec.Command(cmd).Output()

	//availBucketsOutput, err := exec.Command("gsutil", "ls", "-R", "\"gs://cill-raccoon-gj/\"").Output()
	availBucketsOutput, err := exec.Command("gsutil", "ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	var parsedBucketsOutput = lib.ParseCmdOutput(availBucketsOutput)

	return ui.Home("Available buckets", parsedBucketsOutput)
}

func main() {

	// TODO Asap2, assign this to a value, and pass it to the exec command for the bucket.
	var bucketName string = ShowAvailBuckets()

	// TODO put this on a concurrent block, that writes to the log file.
	bucketFolderData, err := exec.Command("gsutil", "ls", "-R", "\""+bucketName+"\"").Output()
	if err != nil {
		fmt.Println("command failed to run.")
		log.Fatal(err)
	}
	var parsedBucketFolderData []string = lib.ParseCmdOutput(bucketFolderData)

	// TODO get the data from gsutil
	//var bucketpath string = ""
	//var cmd = "gsutil ls -R \"" + bucketpath + "\""
	//exec.Command(cmd)

	// TODO bubbletea loading icon while waiting for gsutil to return list.

	if bucketName != "" {
		if result := ui.Home(bucketName, parsedBucketFolderData); result != "" {
			fmt.Println(result)
		}
	}
}
