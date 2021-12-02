package main

import "net"

//func connHandler(c net.Conn) {
//	for {
//		cnt, err := c.Read(buf)
//		c.Write(buf)
//	}
//}
//func main() {
//	server, err := net.Listen("tcp", ":1208")
//	for {
//		conn, err := server.Accept()
//		go connHandler(conn)
//	}
//}

func connHandler(c net.Conn) {
	for {
		c.Write(...)
		c.Read(...)
	}
}
func main() {
	conn, err := net.Dial("tcp", "localhost:1208")
	connHandler(conn)
}
