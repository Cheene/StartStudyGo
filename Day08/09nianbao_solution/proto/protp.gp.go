package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

//编码
func Encode(message string) ([]byte, error) {
	//1 获取长度
	var length = int32(len(message))
	//2 定义消息体
	var pkg = new(bytes.Buffer)
	//3 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//4 写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	//5 返回
	return pkg.Bytes(), nil
}

//解码
func Decode(reader *bufio.Reader) (string, error) {
	//1获取消息体的长度
	lengthByte, _ := reader.Peek(4)
	//2 获取存储长度位置的内容
	lengthBuff := bytes.NewBuffer(lengthByte)
	//3 读取值并且保存到 lengthBuff之中
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	//4 判断消息体的长度为0,返回
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	//5 接收信息
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil

}
