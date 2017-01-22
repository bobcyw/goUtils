package goUtils

import "testing"

func TestNetworkString_Get(t *testing.T) {
	baidu := NetworkString("http://www.baidu.com")
	htmlContent, err := baidu.Get()
	if err != nil{
		t.Fatal(err.Error())
	}
	store := FileString("temp/baidu.html")
	err = store.Write(htmlContent.Content)
	if err != nil{
		t.Fatal(err.Error())
	}
}
