package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	myfolder := `C:\file\intern_in_company\facade2\workflow_individual\file`

	files, _ := ioutil.ReadDir(myfolder)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fmt.Println(file.Name())
		}
	}
}
