package socks5

import (
	"net"
	"testing"
)

func TestSocks5_Server(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:1085")
	if nil != err {
		return
	}
	for {
		_, err := ln.Accept()
		if nil != err {
			break
		}

	}
}
