package main

import (
	"bytes"
	"encoding/binary"
	"os"
)

func IntToByte(num int64)[]byte{
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr(err)
	return buffer.Bytes()
}

func CheckErr(err error){
	if err != nil{
		println("err occur: ", err)
		os.Exit(1)
	}
}