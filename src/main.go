package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	storage "cloud.google.com/go/storage"
	"github.com/lonestar137/gcpfm/src/lib"
	"github.com/lonestar137/gcpfm/src/ui"
	"google.golang.org/api/iterator"
)

// Goal: make GCP ranger like filemanager, objective is to make GCP faster to navigate and use.
// You can press y on any dir/file and it will close the program and print the path to console.

func ShowAvailBuckets() string {
	//var cmd = "gsutil ls"
	//availBucketsOutput, err := exec.Command(cmd).Output()

	//availBucketsOutput, err := exec.Command("gsutil", "ls", "-R", "\"gs://cill-raccoon-gj/\"").Output()
	// availBucketsOutput, err := exec.Command("gsutil", "ls").Output()
	// availBucketsOutput, err := exec.Command("cmd.exe", "/c", "ls").Output()
	availBucketsOutput, err := exec.Command("ls", "-R", "-d", "/usr").Output()
	if err != nil {
		log.Fatal(err)
	}
	var parsedBucketsOutput = lib.ParseCmdOutput(availBucketsOutput)

	return ui.Home("Available buckets", parsedBucketsOutput)
}

type Directory struct {
	title, desc, basePath string
	platform              Platform
}

// enum type creation w/ companion pieces
type Platform int

const (
	GCP Platform = iota // iota keyword instantiates all the contants w/ 0 to N values automatically, i.e. 0,1,2.  This is a unique integer, it prevents conflicts with other constants from other types .
	WINDOWS
	UNIX
)

// companion function to Platform type, now if we try to print our platform, it will use this function to print out the result instead of printing the integer.
func (s Platform) String() string {
	switch s {
	case GCP:
		return "GCP"
	case WINDOWS:
		return "WINDOWS"
	case UNIX:
		return "UNIX"
	default:
		return "Unknown"
	}
}

func walkDirectory(baseDir Directory) {
	// List the files and directories in the basDir

	switch baseDir.platform {
	case WINDOWS, UNIX:
		filepath.WalkDir(baseDir.basePath, func(path string, info os.DirEntry, err error) error {
			// logic for what to do w/ each detected dir
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		})
	case GCP:
		// Create a client for interacting with the Cloud Storage API
		client, err := storage.NewClient(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// List the objects in the bucket
		it := client.Bucket("my-bucket").Objects(context.Background(), nil)
		for {
			obj, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			// Print the object's path
			fmt.Println(filepath.Join("gs://my-bucket", obj.Name))
		}

	default:
		fmt.Println("Invalid platform.")
	}

}

func main() {

	// List the files and directories in the current directory

	var root Directory = Directory{title: "Test", basePath: ".", desc: "just a test dir", platform: UNIX}
	walkDirectory(root)

}
