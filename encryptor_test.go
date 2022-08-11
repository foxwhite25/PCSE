package main

import (
	"io/ioutil"
	"testing"
)

func TestDecodeEncode(t *testing.T) {
	bs, err := ioutil.ReadFile("./save.pcsave")
	if err != nil {
		t.Error(err)
		return
	}
	result, err := Decode(bs)
	if err != nil {
		t.Error(err)
		return
	}
	println(result.MetaData.CustomName.CustomName)
	bs, err = result.Encode()
	if err != nil {
		t.Error(err)
		return
	}
	err = ioutil.WriteFile("./out.pcsave", bs, 777)
	if err != nil {
		t.Error(err)
		return
	}
}
