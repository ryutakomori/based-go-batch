package util

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "log"
)

func Writer(path string, body []byte) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		// Openエラー処理
		log.Println(err.Error())
	}
	defer file.Close()

	file.Write(([]byte)(body))
}

func Reader(path string, body []byte) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		// Openエラー処理
		log.Println(err.Error())
	}
	defer file.Close()
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func List(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {

		if file.IsDir() {
			paths = append(paths, List(filepath.Join(path, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(path, file.Name()))
	}

	return paths
}

func ListCSV(path string) []string {
	paths := List(path)
	var retpaths []string
	for _, path := range paths {
		if filepath.Ext(path) == ".csv" {
			retpaths = append(retpaths, path)
		}
	}

	return retpaths
}
