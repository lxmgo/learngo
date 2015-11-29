// Go 之 原子操作
//
// Copyright (c) 2015 - Batu <1235355@qq.com>
//
// 创建一个文件存放数据,在同一时刻,可能会有多个Goroutine分别进行对此文件的写操作和读操作.
// 每一次写操作都应该向这个文件写入若干个字节的数据,作为一个独立的数据块存在,这意味着写操作之间不能彼此干扰,写入的内容之间也不能出现穿插和混淆的情况
// 每一次读操作都应该从这个文件中读取一个独立完整的数据块.它们读取的数据块不能重复,且需要按顺序读取.
// 例如: 第一个读操作读取了数据块1,第二个操作就应该读取数据块2,第三个读操作则应该读取数据块3,以此类推
// 对于这些读操作是否可以被同时执行,不做要求. 即使同时进行,也应该保持先后顺序.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"os"
	"errors"
	"io"
)

//数据文件的接口类型
type DataFile interface {
	// 读取一个数据块
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号
	Rsn() int64
	// 获取最后写入的数据块的序列号
	Wsn() int64
	// 获取数据块的长度
	DataLen() uint32
}

//数据类型
type Data []byte

//数据文件的实现类型
type myDataFile struct {
	f *os.File	//文件
	fmutex sync.RWMutex //被用于文件的读写锁
	rcond   *sync.Cond   //读操作需要用到的条件变量
	woffset int64 // 写操作需要用到的偏移量
	roffset int64 // 读操作需要用到的偏移量
	dataLen uint32 //数据块长度
}

//初始化DataFile类型值的函数,返回一个DataFile类型的值
func NewDataFile(path string, dataLen uint32) (DataFile, error){
	//f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	f,err := os.Create(path)
	if err != nil {
		fmt.Println("Fail to find", f, "cServer start Failed")
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}

	df := &myDataFile{
		f : f,
		dataLen:dataLen,
	}
	//创建一个可用的条件变量(初始化),返回一个*sync.Cond类型的结果值,我们就可以调用该值拥有的三个方法Wait,Signal,Broadcast
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

//获取并更新读偏移量,根据读偏移量从文件中读取一块数据,把该数据块封装成一个Data类型值并将其作为结果值返回

func (df *myDataFile) Read() (rsn int64, d Data, err error){
	// 读取并更新读偏移量
	var offset int64
	for {
		offset = atomic.LoadInt64(&df.roffset)
		if atomic.CompareAndSwapInt64(&df.roffset, offset, (offset + int64(df.dataLen))){
			break
		}
	}

	//读取一个数据块,最后读取的数据块序列号
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	//读写锁:读锁定
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				//暂时放弃fmutex的 读锁,并等待通知的到来
				df.rcond.Wait()
				continue
			}
		}
		break
	}
	d = bytes
	return
}

func (df *myDataFile) Write(d Data) (wsn int64, err error){
	//读取并更新写的偏移量
	var offset int64
	for {
		offset = atomic.LoadInt64(&df.woffset)
		if atomic.CompareAndSwapInt64(&df.woffset, offset, (offset + int64(df.dataLen))){
			break
		}
	}


	//写入一个数据块,最后写入数据块的序号
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen){
		bytes = d[0:df.dataLen]
	}else{
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	//发送通知
	df.rcond.Signal()
	return
}

func (df *myDataFile) Rsn() int64{
	offset := atomic.LoadInt64(&df.roffset)
	return offset / int64(df.dataLen)
}

func (df *myDataFile) Wsn() int64{
	offset := atomic.LoadInt64(&df.woffset)
	return offset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func main(){
	//简单测试下结果
	var dataFile DataFile
	dataFile,_ = NewDataFile("./mutex_2015_1.dat", 10)

	var d=map[int]Data{
		1:[]byte("batu_test1"),
		2:[]byte("batu_tstt2"),
		3:[]byte("batu_test3"),
	}

	//写入数据
	for i:= 1; i < 4; i++ {
		go func(i int){
			wsn,_ := dataFile.Write(d[i])
			fmt.Println("write i=", i,",wsn=",wsn, ",success.")
		}(i)
	}

	//读取数据
	for i:= 1; i < 4; i++ {
		go func(i int){
			rsn,d,_ := dataFile.Read()
			fmt.Println("Read i=", i,",rsn=",rsn,",data=",d, ",success.")
		}(i)
	}

	time.Sleep(10 * time.Second)
}