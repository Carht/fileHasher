package main

import (
	"io"
	"os"
	"fmt"
	"flag"
	"log"
	"crypto/md5"
	"crypto/sha512"
	"path/filepath"
)

type FileHasher struct {
	filename string
	hash []byte
}

var (
	flag_path string
	flag_hash string
)

func init() {
	flag.StringVar(&flag_path, "p", ".", "File path")
	flag.StringVar(&flag_hash, "h", "md5", "Hash algoritm")
}

func (h *FileHasher) md5(path string) ([]byte, error) {
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

func (h *FileHasher) md5dir(path string) ([]byte, error) {
	md5 := md5.New()
	io.WriteString(md5, path)

	return md5.Sum(nil), nil
}

func (h *FileHasher) sha512(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sha512 := sha512.New()
	if _, err := io.Copy(sha512, file); err != nil {
		log.Fatal(err)
	}

	return sha512.Sum(nil), nil
}

func (h *FileHasher) sha512dir(path string) ([]byte, error) {
	sha512 := sha512.New()
	io.WriteString(sha512, path)

	return sha512.Sum(nil), nil
}
   
func main() {
	f := new(FileHasher)
	flag.Parse()

	_, err := os.Stat(flag_path)
	if err != nil {
		log.Fatal(err)
	}

	path, err := filepath.Abs(flag_path)
	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if flag_hash == "md5" {
			if !info.IsDir() {
				f.filename = path
				f.hash, err = f.md5(f.filename)
				if err != nil {
					log.Fatal(err)
				} 

				fmt.Printf("%s, %x\n", f.filename, f.hash)
				
			} else {
				f.filename = path
				f.hash, err = f.md5dir(f.filename)
				if err != nil {
					log.Fatal(err)
				}
				
				fmt.Printf("%s, %x\n", f.filename, f.hash)
			}
		}

		if flag_hash == "sha512" {
			if !info.IsDir() {
				f.filename = path
				f.hash, err = f.sha512(f.filename)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("%s, %x\n", f.filename, f.hash)
			} else {
				f.filename = path
				f.hash, err = f.sha512dir(f.filename)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("%s, %x\n", f.filename, f.hash)
			}
		}
		
		return nil
	})
}
