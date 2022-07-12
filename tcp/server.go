package riptcp

import (
	"net"

	"github.com/sharpvik/rip"
)

type server struct {
	rip.Handler
	listener net.Listener
}

func NewServer(h rip.Handler) *server {
	return &server{
		Handler: h,
	}
}

// ListenAndServe opens a TCP connection and handles traffic according to the
// RIP/TCP conventions.
func (s *server) ListenAndServe(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	s.listener = listener
	return s.acceptConnectionsTCP()
}

func (s *server) ServeTCP(conn net.Conn) (err error) {
	req, e := ReadRequest(conn)
	if e != nil {
		return SendResponse(conn, rip.ResponseError(e))
	}
	return SendResponse(conn, s.Handle(req))
}

func (s *server) acceptConnectionsTCP() (err error) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		go s.handleConnectionTCP(conn)
	}
}

func (s *server) handleConnectionTCP(conn net.Conn) (err error) {
	for {
		if err = s.ServeTCP(conn); err != nil {
			return
		}
	}
}
