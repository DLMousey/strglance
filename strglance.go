package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("strglance can only be used as a pipe command")
		fmt.Println("Example: echo 'Hello, World!' | strglance")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	input = strings.TrimSpace(input)
	inputBytes := []byte(input)

	md5Hash := md5.Sum(inputBytes)
	sha1Sum := sha1.Sum(inputBytes)
	sha256Sum := sha256.Sum256(inputBytes)
	length := len(input)
	byteLength := len(inputBytes)

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	_, _ = fmt.Fprintf(w, "input\t%s\n", input)
	_, _ = fmt.Fprintf(w, "input (bytes)\t%s\n", hex.EncodeToString(inputBytes))
	_, _ = fmt.Fprintf(w, "length (string)\t%d\n", length)
	_, _ = fmt.Fprintf(w, "length (bytes)\t%d\n", byteLength)
	_, _ = fmt.Fprintf(w, "MD5 Hash\t%x\n", md5Hash)
	_, _ = fmt.Fprintf(w, "SHA1 Hash\t%x\n", sha1Sum)
	_, _ = fmt.Fprintf(w, "SHA256 Hash\t%x\n", sha256Sum[:])
	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
