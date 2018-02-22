package main

import "strings"

// Join ...
func Join(path string, extend string) string {
	if strings.HasSuffix(path, "/") {
		return path + extend
	}
	return path + "/" + extend

}
