package walkers

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	dt "github.com/lonestar137/gcpfm/src/datatypes"
	"google.golang.org/api/iterator"
)

/*
This function will be responsible for reading the dir based on platform.  The known dirs will be cached to a file/in-memory using encoding/gob lib or go-cache or golang-lru.

go-cache and golang-lru seem best option, they provide filebased or in-memory caching and handle expirations for you.
*/
func WalkDirectory(baseDir dt.Directory) {
	// List the files and directories in the basDir

	switch baseDir.Platform {
	case dt.WINDOWS, dt.UNIX:
		filepath.WalkDir(baseDir.BasePath, func(path string, info os.DirEntry, err error) error {
			fmt.Println("UNIX strated")
			// logic for what to do w/ each detected dir
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		})
	case dt.GCP:
		fmt.Println("GCP strated")
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
