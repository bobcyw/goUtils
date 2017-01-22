package goUtils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"encoding/json"
)

type FileString string

func (inst *FileString)ReadObj(obj interface{}) error{
	content, err := inst.Read()
	if err != nil{
		return err
	}
	if err = json.Unmarshal(content, obj); err != nil{
		return err
	}
	return nil
}

func (inst *FileString)Read() ([]byte, error){
	return ioutil.ReadFile(string(*inst))
}

func (inst *FileString)Write(data []byte) error{
	//check path
	path, err := filepath.Abs(string(*inst))
	if err != nil{
		return err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
    os.Mkdir(path, 0644)
	}
	//write file
	return ioutil.WriteFile(string(*inst), data, 0644)
}

func (inst *FileString)Delete()error {
	return os.Remove(string(*inst))
}

