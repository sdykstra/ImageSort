package main

import (
	"io/ioutil"
	"log"
)

func findFile(dirName string) ([]string, []string) {

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	var fileListNames []string
	var dirLists []string
	for _, file := range files {
		if file.IsDir() != true {
			fileListNames = append(fileListNames, file.Name())
			dirLists = append(dirLists, dirName)
		} else {
			var temp []string
			temp, dirLists = findFile(dirName + file.Name() + "/")
			for j := 0; j < len(temp); j++ {
				if file.IsDir() != true {
					fileListNames = append(fileListNames, file.Name()+"/"+temp[j])
					dirLists = append(dirLists, dirName+file.Name()+"/")
				}
			}
		}
	}
	return fileListNames, dirLists

}
