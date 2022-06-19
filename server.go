package rip

import (
	"net"
)

type Server struct {
	*resolver
	listener net.Listener
}

// ListenAndServeTCP opens a TCP connection and handles traffic according to
// the RIP/TCP conventions.
func (s *Server) ListenAndServeTCP(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	s.listener = listener
	return s.acceptConnections()
}

func (s *Server) acceptConnections() (err error) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) (err error) {
	for {
		if err = s.handleRequest(conn); err != nil {
			return
		}
	}
}

func (s *Server) handleRequest(conn net.Conn) (err error) {
	req, e := ReadRequest(conn)
	if e != nil {
		return ResponseError(e).Send(conn)
	}
	return s.Handle(req).Send(conn)
}
