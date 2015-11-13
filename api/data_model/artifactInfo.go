package data_model

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type ArtifactInfo struct {
	content_len int64
	name        string
	mod_time    string
}

func NewArtifactInfo(header *http.Header) *ArtifactInfo {
	file_header := header.Get("Content-Disposition")
	fmt.Printf("Header[Content-Disposition]= %s\n", file_header)

	size_regexp := regexp.MustCompile("size=[0-9]*")
	name_regexp := regexp.MustCompile("filename=\".*\"")

	current_len := size_regexp.FindString(file_header)
	file_size, err := strconv.Atoi(strings.Split(current_len, "=")[1])
	if err != nil {
		//TODO you need to handle error
	}
	fmt.Printf("content-length string: %d\n", file_size)
	file_name := name_regexp.FindString(file_header)
	file_name = strings.Split(file_name, "\"")[1]
	fmt.Printf("filename string: %s\n", file_name)

	ml_str := header.Get("Last-Modified")

	return &ArtifactInfo{
		content_len: file_size,
		name:        file_name,
		mod_time:    ml_str,
	}
}
