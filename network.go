package goUtils

import (
	"net/http"
	"io/ioutil"
)

type HtmlContent struct {
	Header  map[string][]string
	Content []byte
}

type NetworkString string

func (inst *NetworkString)Get() (htmlContent HtmlContent, err error){
	response, err := http.Get(string(*inst))
	if err != nil{
		return
	}
	htmlContent.Header = response.Header
	htmlContent.Content, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil{
		return
	}
	return
}
