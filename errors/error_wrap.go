package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ioutil.ReadFile(filepath.Join(home, ".settings.xml"))
	// use go 1.13
	return config, errors.Wrap(err, "could not read config")


}

func main() {
	_, err := ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
