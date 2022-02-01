package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func readFileStandardWay() {
	// get the file
	file, err := os.Open("README.md")
	if err != nil {
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

// since reading file is very common, there's a shorthand syntax for it
func readFileShorthandWay() {
	bs, err := ioutil.ReadFile("README.md")
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func createFile() {
	file, err := os.Create("file.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Hello, World!")
}

func readDirectoryContents() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func walkDirectories() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

func main() {
	readFileStandardWay()
	readFileShorthandWay()
	createFile()
	readDirectoryContents()
	walkDirectories()
}
