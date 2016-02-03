package protocol
import (
	"bytes"
	"encoding/binary"
)

const (
	//消息头内容
	TCPHeader = "atc.wiki"
	//消息头长度
	TCPHeaderLength = 8
	//发送内容长度
	TCPSaveDataLength = 4
)

//封包
func Packet(message []byte) []byte {
	return append(append([]byte(TCPHeader),IntToBytes(len(message))...), message...)
}

//解包
func Unpack(buf []byte, readerChan chan []byte) []byte{
	length := len(buf)

	var i int
	for i = 0; i < length; i ++ {
		//是否为一个完整消息头
		if length < i + TCPHeaderLength + TCPSaveDataLength {
			break
		}

		if string(buf[i:i + TCPHeaderLength]) == TCPHeader {
			//消息内容长度
			messageLength := BytesToInt(buf[i+TCPHeaderLength : i + TCPHeaderLength + TCPSaveDataLength])
			//是否为一个完整消息实体
			if length < i + TCPHeaderLength + TCPSaveDataLength + messageLength {
				break
			}
			//截取一个完整消息内容
			data := buf[i + TCPHeaderLength + TCPSaveDataLength : i + TCPHeaderLength + TCPSaveDataLength + messageLength]
			//发送给管道
			readerChan <- data

			//从下一个消息头读取
			i += TCPHeaderLength + TCPSaveDataLength + messageLength - 1
		}
	}

	//读完后初始化截断缓冲器
	if i == length {
		return make([]byte,0)
	}

	//返回截断缓冲内容,下次读取时继续
	return buf[i:]
}


//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}