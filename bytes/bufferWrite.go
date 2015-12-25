// Go 之 bytes.buffer Write
//
// Copyright (c) 2015 - Batu <1235355@qq.com>
//
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	//newBuffer 整形转换成字节
	var n int = 10000
	intToBytes := IntToBytes(n)
	fmt.Println("==========int to bytes========")
	fmt.Println(intToBytes)
	//NewBufferString
	TestBufferString()
	//write
	BufferWrite()
	//WriteString
	BufferWriteString()
	//WriteByte
	BufferWriteByte()
	//WriteRune
	BufferWriteRune()

}


func IntToBytes(n int) []byte {
	x := int32(n)
	//创建一个内容是[]byte的slice的缓冲器
	//与bytes.NewBufferString("")等效
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func TestBufferString(){
	buf1:=bytes.NewBufferString("swift")
	buf2:=bytes.NewBuffer([]byte("swift"))
	buf3:=bytes.NewBuffer([]byte{'s','w','i','f','t'})
	fmt.Println("===========以下buf1,buf2,buf3等效=========")
	fmt.Println("buf1:", buf1)
	fmt.Println("buf2:", buf2)
	fmt.Println("buf3:", buf3)
	fmt.Println("===========以下创建空的缓冲器等效=========")
	buf4:=bytes.NewBufferString("")
	buf5:=bytes.NewBuffer([]byte{})
	fmt.Println("buf4:", buf4)
	fmt.Println("buf5:", buf5)
}

func BufferWrite(){
	fmt.Println("===========以下通过Write把swift写入Learning缓冲器尾部=========")
	newBytes := []byte("swift")
	//创建一个内容Learning的缓冲器
	buf := bytes.NewBuffer([]byte("Learning"))
	//打印为Learning
	fmt.Println(buf.String())
	//将newBytes这个slice写到buf的尾部
	buf.Write(newBytes)
	fmt.Println(buf.String())
}

func BufferWriteString(){
	fmt.Println("===========以下通过Write把swift写入Learning缓冲器尾部=========")
	newString := "swift"
	//创建一个string内容Learning的缓冲器
	buf := bytes.NewBufferString("Learning")
	//打印为Learning
	fmt.Println(buf.String())
	//将newString这个string写到buf的尾部
	buf.WriteString(newString)
	fmt.Println(buf.String())
}

func BufferWriteByte(){
	fmt.Println("===========以下通过WriteByte把swift写入Learning缓冲器尾部=========")
	var newByte byte = '!'
	//创建一个string内容Learning的缓冲器
	buf := bytes.NewBufferString("Learning")
	//打印为Learning
	fmt.Println(buf.String())
	//将newString这个string写到buf的尾部
	buf.WriteByte(newByte)
	fmt.Println(buf.String())
}

func BufferWriteRune(){
	fmt.Println("===========以下通过WriteRune把\"好\"写入Learning缓冲器尾部=========")
	var newRune = '好'
	//创建一个string内容Learning的缓冲器
	buf := bytes.NewBufferString("Learning")
	//打印为Learning
	fmt.Println(buf.String())
	//将newString这个string写到buf的尾部
	buf.WriteRune(newRune)
	fmt.Println(buf.String())
}