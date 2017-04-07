package main

import (
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func getImageDates(imageList []string, dirLocation []string) []string {
	var dateList []string

	for i := 0; i < len(imageList); i++ {
		_, err := os.Stat(dirLocation[i] + imageList[i])
		if err != nil {
			if os.IsNotExist(err) {
				dateList = append(dateList, "File does not exist.")
			}
		} else {
			dateList = append(dateList, getImageDate(dirLocation[i]+imageList[i]))
		}
	} // end for loop

	return dateList
} // end getImageDates

func getImageDate(fname string) string {

	f, err := os.Open(fname)
	if err != nil {
		return err.Error()
	}
	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		return err.Error()
	}
	// Date and time of image was taken
	tm, _ := x.DateTime()

	return tm.String()
} // end getImageDate
