package main

import (
	"io"
	"os"
	"fmt"
	"flag"
	"log"
	"crypto/md5"
	"path/filepath"
)

type FileHasher struct {
	filename string
	hash []byte
}

var flag_path string

func init() {
	flag.StringVar(&flag_path, "p", ".", "File path")
}

func filemd5(path string) ([]byte , error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	md5 := md5.New()
	if _, err := io.Copy(md5, file); err != nil {
		log.Fatal(err)
	}

	return md5.Sum(nil), nil
}

func main() {
	fhs := new(FileHasher)

	_, err := os.Stat(flag_path)
	if err != nil {
		log.Fatal(err)
	}

	path, err := filepath.Abs(flag_path)
	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fhs.filename = path
			fhs.hash, err = filemd5(path)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s, %x\n", fhs.filename, fhs.hash)
		}
		
		return nil
	})
}
