package goUtils

import (
	"io/ioutil"
)

type FileString string

func (inst *FileString)Open() ([]byte, error){
	return ioutil.ReadFile(*inst)
}
