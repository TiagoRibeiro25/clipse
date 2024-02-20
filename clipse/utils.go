package main

import (
	"bytes"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

/* General purpose functions to be used by other modules
 */

// Avoids repeat code by handling errors in a uniform way
func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// Contains checks if a string exists in the most recent 3 items
func contains(slice []ClipboardItem, str string) bool {
	if len(slice) > 3 {
		slice = slice[:3]
	}
	for _, item := range slice {
		if item.Value == str {
			return true
		}
	}
	return false
}

// Shortens string val to show in list view
func shorten(s string) string {
	if len(s) <= maxChar { // maxChar defined in constants.go
		return strings.ReplaceAll(s, "\n", "\\n")
	}
	return strings.ReplaceAll(s[:maxChar-3], "\n", "\\n") + "..."
}

func checkDataType(data string) string {
	/*
	   Confirms if clipboard data is currently folding a file vs a string
	*/
	dataBytes := []byte(data)
	reader := bytes.NewReader(dataBytes)

	_, err := png.Decode(reader)
	if err == nil {
		return "png"
	}
	_, err = jpeg.Decode(reader)
	if err == nil {
		return "jpg"
	}
	_, err = gif.Decode(reader)
	if err == nil {
		return "gif"
	}

	return "text"

}
