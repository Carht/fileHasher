package main

import (
	"os"
	"fmt"
	"flag"
	"log"
	"path/filepath"
	"github.com/carht/fileHasher/singledir"
	"github.com/carht/fileHasher/walkerdir"
)

var (
	flag_path string
	flag_hash string
	flag_singledir string
)

func init() {
	flag.StringVar(&flag_path, "p", ".", "File path")
	flag.StringVar(&flag_hash, "h", "md5", "Hash algoritm")
}
   
func list_md5(paths []string) error {
	for _, file := range paths {
		fmt.Printf("%s, %x\n", file, singledir.Hmd5(file))
	}

	return nil	
}

func list_sha512(paths []string) error {
	for _, file := range paths {
		fmt.Printf("%s, %x\n", file, singledir.Hsha512(file))
	}

	return nil
}

func main() {
	flag_walkerp := flag.Bool("w", false, "Directory Walker")
	f := walkerdir.FileHasher{}
	flag.Parse()

	_, err := os.Stat(flag_path)
	if err != nil {
		log.Fatal(err)
	}

	path, err := filepath.Abs(flag_path)
	if err != nil {
		log.Fatal(err)
	}

	if *flag_walkerp {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if flag_hash == "md5" {
				if !info.IsDir() {
					f.Filename = path
					f.Hash, err = f.Md5(f.Filename)
					if err != nil {
						log.Fatal(err)
					} 
					
					fmt.Printf("%s, %x\n", f.Filename, f.Hash)
					
				} else {
					f.Filename = path
					f.Hash, err = f.Md5dir(f.Filename)
					if err != nil {
						log.Fatal(err)
					}
					
					fmt.Printf("%s, %x\n", f.Filename, f.Hash)
				}
			}
			
			if flag_hash == "sha512" {
				if !info.IsDir() {
					f.Filename = path
					f.Hash, err = f.Sha512(f.Filename)
					if err != nil {
						log.Fatal(err)
					}
					
					fmt.Printf("%s, %x\n", f.Filename, f.Hash)
				} else {
					f.Filename = path
					f.Hash, err = f.Sha512dir(f.Filename)
					if err != nil {
						log.Fatal(err)
					}
					
					fmt.Printf("%s, %x\n", f.Filename, f.Hash)
				}
			}
			
			return nil
		})
	} else {
		list_files, err := singledir.ReadDirNames(path)
		
		lf_fullpath, err := singledir.ToAbs(path, list_files)

		if flag_hash == "md5" {
			err = list_md5(lf_fullpath)
			if err != nil {
				log.Fatal(err)
			}
		}

		if flag_hash == "sha512" {
			err = list_sha512(lf_fullpath)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
