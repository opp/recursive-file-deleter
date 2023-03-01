package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var fileToDelete *string = flag.String("file", ".DS_Store", "Choose file name to delete.")
	var searchPath *string = flag.String("path", "", "Choose file path to recursively search.")
	var runType *string = flag.String("run", "dry", "Choose between 'dry' run or a 'real' run, eg. -run=real")
	flag.Parse()

	fmt.Printf("Running '%s' run; searching for file '%s' in path '%s'\n", *runType, *fileToDelete, *searchPath)
	fileSearch(runType, searchPath, fileToDelete)
}

func fileSearch(run *string, searchPath *string, fileToDelete *string) {
	err := filepath.Walk(*searchPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatal(err)
			}
			if info.Name() == *fileToDelete {
				log.Printf("Found: %s\n", path)
				if *run == "real" {
					err = os.Remove(path)
					if err != nil {
						log.Fatal(err)
					}
					log.Println("Deleted: ", path)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	log.Println("Finished.")
}
