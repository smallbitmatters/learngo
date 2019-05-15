// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
	// Most other languages do this:
	try {
		// open a file
		// throws an exception
	} catch (ExceptionType name) {
		// handle the error
	} finally {
		// close the file
	}

	// Go way:
	file, err := // open the file
	if err != nil {
		// handle the error
	}
	// close the file
*/

func main() {
	files := []string{
		"pngs/cups-jpg.png",
		"pngs/forest-jpg.png",
		"pngs/golden.png",
		"pngs/work.png",
		"pngs/shakespeare-text.png",
		"pngs/empty.png",
	}

	pngs := detect(files)

	fmt.Printf("Correct Files:\n")
	for _, png := range pngs {
		fmt.Printf(" + %s\n", png)
	}
}

func detect(filenames []string) (pngs []string) {
	const pngHeader = "\x89PNG\r\n\x1a\n"

	buf := make([]byte, len(pngHeader))

	for _, filename := range filenames {
		if read(filename, buf) != nil {
			continue
		}

		if bytes.Equal([]byte(pngHeader), buf) {
			pngs = append(pngs, filename)
		}
	}
	return
}

// read reads len(buf) bytes to buf from a file
func read(filename string, buf []byte) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	if fi.Size() <= int64(len(buf)) {
		return fmt.Errorf("file size < len(buf)")
	}

	_, err = io.ReadFull(file, buf)
	return err
}

func detectPNGUnsafeAndVerbose(filenames []string) (pngs []string) {
	const pngHeader = "\x89PNG\r\n\x1a\n"

	buf := make([]byte, len(pngHeader))

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			continue
		}

		fi, err := file.Stat()
		if err != nil {
			file.Close()
			continue
		}

		if fi.Size() <= int64(len(pngHeader)) {
			file.Close()
			continue
		}

		_, err = io.ReadFull(file, buf)
		file.Close()
		if err != nil {
			continue
		}

		if bytes.Equal([]byte(pngHeader), buf) {
			pngs = append(pngs, filename)
		}
	}
	return
}

// defers don't run when the loop block ends
func detectPNGWrong(filenames []string) (pngs []string) {
	const pngHeader = "\x89PNG\r\n\x1a\n"

	buf := make([]byte, len(pngHeader))

	for _, filename := range filenames {
		fmt.Printf("processing: %s\n", filename)

		file, err := os.Open(filename)
		if err != nil {
			continue
		}

		defer func() {
			fmt.Printf("closes    : %s\n", filename)
			file.Close()
		}()

		fi, err := file.Stat()
		if err != nil {
			continue
		}

		if fi.Size() <= int64(len(pngHeader)) {
			continue
		}

		_, err = io.ReadFull(file, buf)
		if err != nil {
			continue
		}

		if bytes.Equal([]byte(pngHeader), buf) {
			pngs = append(pngs, filename)
		}
	}
	return
}