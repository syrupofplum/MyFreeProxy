package socks5

type ClientMethods struct {
	VER      byte
	NMETHODS byte
	METHODS  []byte
}

type ServerMethods struct {
	VER    byte
	METHOD byte
}
