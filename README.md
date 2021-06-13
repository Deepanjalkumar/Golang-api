# Golang-api

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	path, _ := os.Getwd()

	crc := filepath.Join(path + "/owasp.json")

	file_byte, _ := ioutil.ReadFile(crc)

	fmt.Println(string(file_byte))
}
