package main

import (
	"fmt"
	"time"

	dt "github.com/lonestar137/gcpfm/src/datatypes"
	walk "github.com/lonestar137/gcpfm/src/lib/walkers"
	ds "github.com/lonestar137/gcpfm/src/ui/input/directory"
	start "github.com/lonestar137/gcpfm/src/ui/input/startlist"
)

// func ShowAvailBuckets() string {
// 	//var cmd = "gsutil ls"
// 	//availBucketsOutput, err := exec.Command(cmd).Output()

// 	//availBucketsOutput, err := exec.Command("gsutil", "ls", "-R", "\"gs://cill-raccoon-gj/\"").Output()
// 	// availBucketsOutput, err := exec.Command("gsutil", "ls").Output()
// 	// availBucketsOutput, err := exec.Command("cmd.exe", "/c", "ls").Output()
// 	availBucketsOutput, err := exec.Command("ls", "-R", "-d", "/usr").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var parsedBucketsOutput = lib.ParseCmdOutput(availBucketsOutput)

// 	return ui.Home("Available buckets", parsedBucketsOutput)
// }

/*
Goal: make GCP ranger like filemanager, objective is to make GCP faster to navigate and use.
You can press y on any dir/file and it will close the program and print the path to console.
*/
func main() {
	// TODO: SEPARATION OF CONCERNS, separate out each piece:
	// 1. user interface
	// 2. business logic
	// 3. data access

	var choice string = start.StartMenu()
	fmt.Println("choice:" + choice)
	var selectedPlatform dt.Platform = dt.GetPlatformType(choice)

	modelSettings := ds.SpawnInputMenu()

	for i := range modelSettings.Inputs {
		fmt.Println(modelSettings.Inputs[i].Value())
	}

	eventLoop := time.Tick(5 * time.Second)

	var root dt.Directory = dt.Directory{Title: "Test", BasePath: ".", Desc: "just a test dir", Platform: selectedPlatform}
	go walk.WalkDirectory(root)

	for range eventLoop {
		// List the files and directories in the current directory -- business logic
		go walk.WalkDirectory(root) // -- I want to put this in a go routine.

		// data cache access

		// user interface
		// - option for root.platform = "your choice"

	}

}
