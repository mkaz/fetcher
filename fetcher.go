/*
 * fetcher.go - http request helper
 * https://github.com/mkaz/fetcher
 *
 * A library to make it a bit easier to do HTTP fetches
 * Includes adding headers, post forms, params
 *
 */

package fetcher

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Fetcher struct {
	Url, Method           string
	Params, Header, Files map[string]string
}

func (f Fetcher) Execute() (result string, err error) {

	var reqBody io.Reader
	var contentType string

	// check if post and add post params
	if f.Method == "POST" {
		reqBody, contentType, err = f.createPostBody()
		if err != nil {
			return "", err
		}
	}

	// build request object
	client := &http.Client{}
	req, err := http.NewRequest(f.Method, f.Url, reqBody)
	if err != nil {
		return "", err
	}

	if f.Method == "POST" {
		req.Header.Add("Content-Type", contentType)
	}

	// add header objects
	for k, v := range f.Header {
		req.Header.Add(k, v)
	}

	// execute request object
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// process response
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	result = string(body)
	return
}

func NewFetcher() (f Fetcher) {
	f.Params = map[string]string{}
	f.Header = map[string]string{}
	f.Files = map[string]string{}
	return f
}

func (f Fetcher) createPostBody() (body io.Reader, contentType string, err error) {

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// add files if we are uploading a file
	for k, v := range f.Files {
		file, err := os.Open(v)
		if err != nil {
			return nil, "", err
		}

		part, err := writer.CreateFormFile(k, filepath.Base(v))
		if err != nil {
			return nil, "", err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return nil, "", err
		}
	}
	
	// add parameters if there are parameters
	for k, v := range f.Params {
		_ = writer.WriteField(k, v)
	}
	
	err = writer.Close()
	if err != nil {
		return
	}

	contentType = writer.FormDataContentType()
	body = &b
	return body, contentType, nil

}
