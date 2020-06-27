package socks5

import (
	"fmt"
	"log"
	"net"
)

type ServerBindConfig struct {
	network string
	address string
}

type Server struct {
	BindConfigs []ServerBindConfig
}

func (s *Server) Listen() {
	go func() {
		for _, bindConfig := range s.BindConfigs {
			ln, err := net.Listen(bindConfig.network, bindConfig.address)
			if nil != err {
				errMsg := fmt.Sprintf("listen %v fail.", bindConfig.address)
				log.Fatalln(errMsg)
			}
			conn, err := ln.Accept()
			if nil != err {
				errMsg := fmt.Sprintf("accept %v fail.", bindConfig.address)
				log.Fatalln(errMsg)
			}
			go s.HandleConn(&conn)
		}
	}()
}

func (s *Server) HandleConn(conn *net.Conn) {

}
