package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
)

var EncryptionKey = []byte{97, 94, 49, 57, 117, 104, 37, 52, 55, 120, 55, 49, 101, 37, 115, 100}

func DecodeXOR(data []byte) []byte {
	for i, _ := range data {
		data[i] ^= EncryptionKey[i%16]
	}
	return data
}

func Decode(data []byte) (save Save, err error) {
	var result [][]byte
	for _, line := range bytes.Split(data, []byte{10}) {
		data = make([]byte, base64.StdEncoding.DecodedLen(len(line)))
		_, err := base64.StdEncoding.Decode(data, line)
		if err != nil {
			return Save{}, err
		}
		decoded := DecodeXOR(data)
		result = append(result, decoded)
	}

	err = json.Unmarshal(result[0], &save.MetaData)
	if err != nil {
		return Save{}, err
	}
	err = json.Unmarshal(result[1], &save.UserData)
	if err != nil {
		return Save{}, err
	}
	return
}

func (s *Save) Encode() (data []byte, err error) {
	metaJson, err := json.Marshal(s.MetaData)
	if err != nil {
		return
	}
	userJson, err := json.Marshal(s.UserData)
	if err != nil {
		return
	}
	var result = [][]byte{metaJson, userJson}
	for i, b := range result {
		result[i] = make([]byte, base64.StdEncoding.EncodedLen(len(b)))
		base64.StdEncoding.Encode(result[i], DecodeXOR(b))
	}
	return bytes.Join(result, []byte{10}), nil
}
