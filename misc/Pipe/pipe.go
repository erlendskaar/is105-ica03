// https://www.socketloop.com/references/golang-io-pipe-function-example
package main

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"unsafe"
)

const hexNummer = "546865736520776F7264"

func main() {

	k := returnHexASCII(hexNummer)

	l := returnBase64(k)

	compressBase64(l)

	returnHex()

	returnHexFromBase64()

}
func pipe() {
	// read and write with pipe
	pReader, pWriter := io.Pipe()

	// use base64 encoder
	b64Writer := base64.NewEncoder(base64.StdEncoding, pWriter)

	gWriter := gzip.NewWriter(b64Writer)

	// write text to be gzipped and push to pipe
	go func() {
		fmt.Println("Start writing")
		n, err := gWriter.Write([]byte("These words will be compressed and pushed into pipe"))

		fmt.Printf("len = %d\n", n)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		gWriter.Close()
		b64Writer.Close()
		pWriter.Close()
		fmt.Printf("Written %d bytes \n", n)
	}()

	// start reading from the pipe(reverse of the above process)
	//Husk at compression er bare en annen mainmetode og at stdout er print.

	// use base64 decoder to wrap pipe Reader
	b64Reader := base64.NewDecoder(base64.StdEncoding, pReader)

	// read gzipped text and decompressed the text
	gReader, err := gzip.NewReader(b64Reader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Look at the final output at the other side of the pipe

	// print out the text
	text, err := ioutil.ReadAll(gReader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", text)
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
	fmt.Println("")

	return e
}

// Implementert etter å ha sett oversettelse fra de heksadesimalerepresentasjonen
func returnHex() {
	src := []byte("These word")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)
	fmt.Println("")
}

//Koder fra Base64 til ASCII til heks.
func returnHexFromBase64() {
	str := "VGhlc2Ugd29yZA=="
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)

		return
	}
	fmt.Printf("%q\n", data)
}

func compressBase64(b64String string) {

	newFile, err := os.Create("compression.gz")
	if err != nil {
		fmt.Print(err)
	}
	w := gzip.NewWriter(newFile)

	fmt.Println("Komprimerer nå til .gz")
	fmt.Println("")

	w.Write([]byte(b64String))

	w.Close()
}
