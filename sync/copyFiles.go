package sync

import (
	"io"
	"os"
)

func copyFiles(files []File) error {
	// 	// create a list of folder from the found files
	// 	folders := mccoy.Set{}
	for _, file := range files {
		// io.Copy()

		// input, err := ioutil.ReadFile(file.src)
		// if err != nil {
		// 	return err
		// }

		// err = ioutil.WriteFile(file.dest, input, 0644)
		// if err != nil {
		// 	return err
		// }

		// src := "words.txt"
		// dst := "words2.txt"

		fin, err := os.Open(file.src)
		if err != nil {
			return err
		}
		defer fin.Close()

		fout, err := os.Create(file.dest)
		if err != nil {
			return err
		}
		defer fout.Close()

		_, err = io.Copy(fout, fin)

		if err != nil {
			return err
		}
	}

	// 	// remove folders with children, os.MkdirAll() will create an entire path
	// 	for _, i := range folders.Items() {
	// 		needle := mccoy.ItemString(i)

	// 		for _, ii := range folders.Items() {
	// 			test := mccoy.ItemString(ii)

	// 			if strings.Contains(test, needle) && test != needle {
	// 				folders.Delete(i)
	// 				break
	// 			}
	// 		}
	// 	}

	// 	// make the folder structure
	// 	for _, i := range folders.Items() {
	// 		path := mccoy.ItemString(i)

	// 		err := os.MkdirAll(path, 0755)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}

	return nil
}
