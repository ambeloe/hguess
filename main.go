package main

import (
	"fmt"
	"os"
)

type HashAlg interface {
	Name() string

	Hash(b []byte) string
}

func main() {
	os.Exit(rMain())
}

func rMain() int {
	var err error
	var f []byte

	var algs = []HashAlg{md2{}, md4{}, md5{}, sha1{}, sha256{}, sha512{}}

	for _, filename := range os.Args[1:] {
		f, err = os.ReadFile(filename)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error reading file: %s, %v\n", filename, err)
			return 1
		}
		for _, hash := range algs {
			fmt.Printf("%s %s %s\n", hash.Name(), hash.Hash(f), filename)
		}
	}

	return 0
}
