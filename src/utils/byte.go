package utils

import (
	"bytes"
	"net"
	"io"
)

func BytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

func ReadBytes(conn *net.TCPConn) (r []byte, e error) {

	buf := make([]byte, 1024)
	len := 0

	for {
		n, err := conn.Read(buf[len:])
		if n == 0 {
			// 取不到数据直接返回
			break
		}
		if n > 0 {
			// 取到数据
			len += n
		}
		if err != nil {
			if err != io.EOF {
				//Error Handler
			}

			break
		}
		if n < 1024 {
			break
		}
	}
	return buf[:len], nil
}
