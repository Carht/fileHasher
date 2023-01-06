package singledir

import (
	"io"
	"os"
	"log"
	"path/filepath"
	"crypto/sha512"
	"crypto/md5"
)


// If the input is a directory, return the list of files.
func ReadDirNames(dirpath string) ([]string, error) {
	f, err := os.Open(dirpath)
	if err != nil {
		return nil, err
	}

	files, err := f.Readdirnames(-1)
	f.Close()

	return files, nil
}


// Get the local paths and transform this to absolute path.
// If the path is "." calculate the absolute path from the file execution,
// else, use the original basepath and concatenate with the files inside this.
func ToAbs(basepath string, paths []string) ([]string, error) {
	full_path := []string{}
	
	for _, file := range paths {
		if basepath != "." {
			if basepath[len(basepath) - 1] == 47 {
				concat := basepath + file
				full_path = append(full_path, concat)
			} else {
				concat := basepath + "/" + file
				full_path = append(full_path, concat)
			}
		} else {
			path, err := filepath.Abs(file)

			if err != nil {
				return nil, err
			}

			full_path = append(full_path, path)
		}
	}

	return full_path, nil
}

// Hash md5sum
// Return the md5sum for a normal file, and for a directory, in this last case
// I use the full path as input string for the md5 algorithm, using the rationale
// of not posible repeated directory names.
func Hmd5(path string) []byte {
	fi, err := os.Lstat(path)

	if err != nil {
		log.Fatal(err)
	}

	if fi.IsDir() {
		md5 := md5.New()
		io.WriteString(md5, path)

		return md5.Sum(nil)
	} else {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		md5 := md5.New()
		if _, err := io.Copy(md5, file); err != nil {
			log.Fatal(err)
		}

		return md5.Sum(nil)
	}	
}

// Return the sha512 for a normal file, and for a directory, in this last case
// I use the full path as input string for sha512 algoritm, using the rationale
// of not posible repeated directory names.
func Hsha512(path string) []byte {
	fi, err := os.Lstat(path)

	if err != nil {
		log.Fatal(err)
	}

	if fi.IsDir() {
		sha512 := sha512.New()
		io.WriteString(sha512, path)

		return sha512.Sum(nil)
	} else {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		sha512 := sha512.New()
		if _, err := io.Copy(sha512, file); err != nil {
			log.Fatal(err)
		}

		return sha512.Sum(nil)
	}
}
