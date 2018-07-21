package main

import (
	"bytes"
	"encoding/binary"
	"os"
	"runtime"
	"fmt"
)

func IntToByte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr(err)
	return buffer.Bytes()
}

func CheckErr(err error) {
	var pos string
	//func Caller(skip int) (pc uintptr, file string, line int, ok bool) {
	pc, file, line, ok := runtime.Caller(1)
	if ok{
		f := runtime.FuncForPC(pc)
		//log.Println(f.Name())
		pos = fmt.Sprintf("%s\n\t%s : %d\n", file, f.Name(), line)
	}
	if err != nil {
		println("err occur: ", err, "\npos: ", pos)
		os.Exit(1)
	}
}
