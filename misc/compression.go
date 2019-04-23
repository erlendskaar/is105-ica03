package main

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
)

//Kode for funksjoner og for programmet.

func main() {
	args := os.Args
	file := args[1]

	g := readFile(file)

	d := returnHexASCII(g)
	a := returnBase64(d)
	compressBase64(a)
}

func readFile(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	fmt.Println("Hex string:", str)

	return str
}

// Returnere en ascii/utf8 representasjon
func returnHexASCII(hex1 string) string {

	ascii, err := hex.DecodeString(hex1)
	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%s", ascii)

	fmt.Println("Fra hex til ASCII:", str)
	fmt.Printf("Størrelse i byte: %T, %d \n", str, unsafe.Sizeof(str))
	fmt.Println("Lengde: ", len(str))
	fmt.Println("")

	return str
}

// Returnere en base64 representasjon
func returnBase64(s string) string {

	// ASCII til base64
	e := base64.StdEncoding.EncodeToString([]byte(s))
	// Lengden av base64 strengen
	r := base64.StdEncoding.EncodedLen(len(e))

	fmt.Println("Fra ASCII til base64:", e)
	fmt.Printf("Størrelse i byte for base64: %T, %d \n", e, unsafe.Sizeof(e))
	fmt.Println("Lengden på stringen i base64:", r)

	return e
}

// Komprimerer til .gz
func compressBase64(b64String string) {

	newFile, err := os.Create("compression.gz")
	if err != nil {
		fmt.Print(err)
	}
	w := gzip.NewWriter(newFile)

	fmt.Println("Komprimerer nå til .gz")

	w.Write([]byte(b64String))

	w.Close()
}
