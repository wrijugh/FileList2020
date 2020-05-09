package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	folder := "D:\\OneDrive - Microsoft\\Wriju_Personal\\E_books_OneDriveBus"

	getAllFiles(folder)
}

func getAllFiles(folder string) {
	err := filepath.Walk(folder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			var dir string = ""
			trueVal := info.IsDir()
			if trueVal == true {
				dir = "[Directory]"
			}

			println(dir, path)
			return nil
		})

	if err != nil {
		log.Println(err)
	}
}

/*
func getAllDirectories(directoryPath string) {

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
*/
