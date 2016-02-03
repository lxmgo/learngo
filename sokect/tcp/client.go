package tcp
import (
	"net"
	"fmt"
	"time"
	"../protocol"
)


func write(conn net.Conn){
	for i := 0; i < 100; i ++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"messagessss\"}"
		_,err := conn.Write(protocol.Packet([]byte(words)))
		CheckErr(err)
	}
	fmt.Println("write end.")
}

func RunClient(){
	addr := "127.0.0.1:9999"
	tcpAdd, err := net.ResolveTCPAddr("tcp4", addr)
	CheckErr(err)

	conn, err := net.DialTCP("tcp", nil, tcpAdd)
	CheckErr(err)

	defer conn.Close()
	fmt.Println("connect success.")
	go write(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}


