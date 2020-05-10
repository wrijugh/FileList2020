package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	folder := "D:\\OneDrive - Microsoft\\Wriju_Personal\\E_books_OneDriveBus"
	dirs := getOnlySubDirectories(folder)

	dataSlice := []string{""}
	for _, dir := range dirs {
		files := getAllFilesFromOneDirectory(dir)
		dataSlice = append(dataSlice, files...) //Learned when appending a slice to another we need to use ...
	}

	//To Print line by line
	/*textToWrite := "List Generated in datehere " //TODO: FormatDate and also mention the root folder
	for _, f := range dataSlice {
		textToWrite = textToWrite + f + "\\r\\n"
	}*/

	//Write to file
	writeToFile("list.txt", folder, dataSlice)
}

//Get files inside one folder
func getAllFilesFromOneDirectory(directoryPath string) []string {
	fileSlice := []string{""}
	if directoryPath == "" {
		return nil
	}

	f, _ := os.Open(directoryPath)

	files, _ := f.Readdir(-1)
	f.Close()

	parentDir := filepath.Dir(directoryPath)
	onlyDir := strings.ReplaceAll(directoryPath, parentDir+"\\", "")
	fileSlice = append(fileSlice, onlyDir)

	dashStr := ""
	for i := 0; i < len(onlyDir); i++ {
		dashStr = dashStr + "="
	}
	fileSlice = append(fileSlice, dashStr)

	i := 1
	for _, file := range files {
		if !file.IsDir() { //not to add only the folder name. This is not needed in file list
			//fileSlice,
			fileSlice = append(fileSlice, strconv.Itoa(i)+". "+file.Name())
			fmt.Println(i, file.Name())
			i++
		}
	}

	return fileSlice
}

//Only get all the directories then for each directory find the files
func getOnlySubDirectories(folder string) []string {
	dirSlice := []string{""}
	filepath.Walk(folder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				dirSlice = append(dirSlice, path)
			}
			return nil
		})

	return dirSlice
}

func writeToFile(fileName, rootFolder string, sliceToWrite []string) {
	/*writeByte := []byte(textToWrite)
	err := ioutil.WriteFile(fileName, writeByte, 0644)
	if err != nil {
		panic(err)
	}*/
	//Create a header like
	//Folder name and created on
	//Created using Go

	// hline2 :=
	file, err := os.OpenFile(fileName, os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	hline1 := "List for the folder " + rootFolder
	//hline2 := "On " +
	_, _ = datawriter.WriteString(hline1 + "\n")
	//_, _ = datawriter.WriteString(hline2 + "\n")
	//This is not writing a multiline -----!!!!!!!!!!!
	for _, data := range sliceToWrite {
		_, _ = datawriter.WriteString(data + "\n")

	}

	datawriter.Flush()
	file.Close()
}
