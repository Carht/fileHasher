package walkerdir

import (
	"io"
	"os"
	"log"
	"crypto/md5"
	"crypto/sha512"
)

type FileHasher struct {
	Filename string
	Hash []byte
}

// Return the md5sum for a normal file.
func (h *FileHasher) Md5(path string) ([]byte, error) {
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

// Return the md5sum for a directory.
func (h *FileHasher) Md5dir(path string) ([]byte, error) {
	md5 := md5.New()
	io.WriteString(md5, path)

	return md5.Sum(nil), nil
}

// Return the sha512 for a normal file.
func (h *FileHasher) Sha512(path string) ([]byte, error) {
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

// Return the sha512 for a directory.
func (h *FileHasher) Sha512dir(path string) ([]byte, error) {
	sha512 := sha512.New()
	io.WriteString(sha512, path)

	return sha512.Sum(nil), nil
}
