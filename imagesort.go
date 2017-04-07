package main

import "fmt"

const dirPath = "../../Testdata/"

func main() {

	//var files =
	var imagesFound, dirName = findFile(dirPath)

	fmt.Printf("This was the Size of the list found %d \n", len(imagesFound))

	var imageDates = getImageDates(imagesFound, dirName)

	//fmt.Println(imageDates)
	for i := 0; i < len(imageDates); i++ {
		fmt.Printf("This is the image found %s and the date of the image %s \n", imagesFound[i], imageDates[i])

	}

	fmt.Printf("End of main")
} // end main
