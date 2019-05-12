package main

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unsafe"
)

//Kode for funksjoner og for programmet.
// f.eks "go run compression.go hex"

func main() {
	args := os.Args
	file := args[1]

	d := returnHexASCII(file)
	a := returnBase64(d)
	compressBase64(a)

}

func readFile(file string) string {
	b, err := ioutil.ReadFile(file)
	fileValue := len(b)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	// Sjekker om fileValue er under 125, dette er for at det ikke skal bli for mange tegn, hvis den er over skriver den ut lengden på stringen i stede for stringen
	if fileValue < 125 {
		fmt.Println("Hex string:", str)
	} else {
		fmt.Println("Hex stringen er på", fileValue, "tegn")
	}
	return str
}

// Returnere en ascii/utf8 representasjon
func returnHexASCII(hex1 string) string {
	fileRead := readFile(hex1)

	ascii, err := hex.DecodeString(fileRead)
	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%s", ascii)
	fileStat, err := os.Stat(hex1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fra hex til ASCII:", str)
	fmt.Println("Størrelse i bytes", fileStat.Size())
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

	//Printer kun ut Base64 stringen hvis den er på under 100 tegn, hvis ikke informerer den kun om lengden på Base64 stringen
	if r < 100 {
		fmt.Println("Fra ASCII til base64:", e)
	} else {
		fmt.Println("Base64 er på", r, "tegn")
	}

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
