package main

import (
	"bytes"
	"encoding/binary"
	"runtime"
	"fmt"
	"os"
)

func IntToHex(arg int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, arg)
	CheckErr(err)
	return buffer.Bytes()
}

func CheckErr(err error) {
	var pos string
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		f := runtime.FuncForPC(pc)
		pos = fmt.Sprintf("%s\n\t%s : %d\n", file, f.Name(), line)
	}
	if err != nil {
		println("err occur: ", err, "\npos: ", pos)
		os.Exit(1)
	}
}
