package main

import (
	"fmt"
	"bytes"
	"encoding/hex"
	"encoding/binary"
)

func main() {
	//tr := bytes.NewBuffer(make([]byte, 10))
	//tr := new(bytes.Buffer)
	tr := bytes.NewBuffer([]byte{})
	//var a int = 10
	binary.Write(tr, binary.BigEndian, byte(17))
	//binary.Write(tr, binary.BigEndian, uint32(20480))
	fmt.Print(hex.Dump(tr.Bytes()))
	//b := tr.Bytes()
	
	//fmt.Print(string(b))
}