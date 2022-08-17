package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/lonestar137/gcpfm/src/lib"
	"github.com/lonestar137/gcpfm/src/ui"
)

// Goal: make GCP ranger like filemanager, objective is to make GCP faster to navigate and use.
// You can press y on any dir/file and it will close the program and print the path to console.

func ShowAvailBuckets() string {
	//var cmd = "gsutil ls"
	//availBucketsOutput, err := exec.Command(cmd).Output()

	availBucketsOutput, err := exec.Command("ls", "-la").Output()
	if err != nil {
		log.Fatal(err)
	}
	var parsedBucketsOutput = lib.ParseCmdOutput(availBucketsOutput)
	ui.Home("Available buckets", parsedBucketsOutput)

	// TODO Asap, have Home return the selection on Enter as a string, kills two birds with one stone.

	return "ui.Home output here"

}

func main() {

	// TODO Asap2, assign this to a value, and pass it to the exec command for the bucket.
	ShowAvailBuckets()

	// TODO get the data from gsutil
	//var bucketpath string = ""
	//var cmd = "gsutil ls -R \"" + bucketpath + "\""
	//exec.Command(cmd)

	// TODO bubbletea loading icon while waiting for gsutil to return list.

	var sampleFolderData []string = []string{
		"gcs://sample-bucket/config/Code/logs",
		"gcs://sample-bucket/config/Code/logs/20220816T195842",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/exthost1",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/exthost1/output_logging_20220816T195848",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/exthost1/output_logging_20220816T200336",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/output_1_20220816T195844",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/output_1_20220816T200333",
		"gcs://sample-bucket/config/Code/logs/20220525T201131",
		"gcs://sample-bucket/config/Code/logs/20220525T201131/exthost1",
		"gcs://sample-bucket/config/Code/logs/20220525T201131/exthost1/output_logging_20220525T201139",
		"gcs://sample-bucket/config/Code/logs/20220525T201131/output_1_20220525T201135",
		"gcs://sample-bucket/config/Code/logs/20220514T230127",
		"gcs://sample-bucket/config/Code/logs/20220514T230127/exthost1",
		"gcs://sample-bucket/config/Code/Service Worker/CacheStorage/c8a89dc5a4dcee8d4646c4b8474ae324ddaac427/c66e042f-72ec-4292-b6a8-cccbea394816/index-dir",
		"gcs://sample-bucket/config/Code/Service Worker/CacheStorage/98e3acaec81ab4b1d4493535e1030a26de1af26a",
		"gcs://sample-bucket/config/Code/Service Worker/CacheStorage/98e3acaec81ab4b1d4493535e1030a26de1af26a/70e7ca36-8902-4d9f-b6cb-b380bf578efc",
		"gcs://sample-bucket/config/Code/Service Worker/CacheStorage/98e3acaec81ab4b1d4493535e1030a26de1af26a/70e7ca36-8902-4d9f-b6cb-b380bf578efc/index-dir",
		"gcs://sample-bucket/config/Code/Backups",
		"gcs://sample-bucket/config/Code/Backups/1660697528120",
		"gcs://sample-bucket/config/Code/Backups/1660697528120/userdata",
		"gcs://sample-bucket/config/Code/Backups/1660697528120/file",
		"gcs://sample-bucket/config/Code/CachedData/c3511e6c69bb39013c4a4b7b9566ec1ca73fc4d5/chrome/wasm/index-dir",
		"gcs://sample-bucket/config/Code/CachedData/c3511e6c69bb39013c4a4b7b9566ec1ca73fc4d5/chrome/js",
		"gcs://sample-bucket/config/Code/CachedData/c3511e6c69bb39013c4a4b7b9566ec1ca73fc4d5/chrome/js/index-dir",
	}

	var bucketName string = strings.Split(sampleFolderData[0], "/")[2]
	ui.Home(bucketName, sampleFolderData)
	// TODO asap3, fmt.Println(ui.Home) to print out the result of your selection.
}
