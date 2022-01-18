package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func main() {
	// non-cryptographic
	h1 := crc32.NewIEEE()
	h1.Write([]byte("test"))
	v := h1.Sum32()
	fmt.Println(v)

	h2, err := getHash("README.md")
	if err != nil {
		return
	}
	fmt.Println(h2)

	// cryptographic
	h3 := sha1.New()
	h3.Write([]byte("test"))
	bs := h3.Sum([]byte{})
	fmt.Println(bs)
}
