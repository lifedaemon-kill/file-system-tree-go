package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	err := getTree(&out, path, printFiles)
	if err != nil {
		return err
	}
	return nil
}
func getTree(out *io.Writer, path string, printFiles bool) error {
	arr, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, obj := range arr {
		if obj.IsDir() {
			err := getTree(out, path+"/"+obj.Name(), printFiles)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(obj.Name())
		}
	}
	return nil
}
