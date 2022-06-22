package rip

import (
	"net"

	"github.com/sharpvik/rip/proto"
	riptcp "github.com/sharpvik/rip/tcp"
)

type Server struct {
	Handler
	listener net.Listener
}

func NewServer(h Handler) *Server {
	return &Server{
		Handler: h,
	}
}

// ListenAndServeTCP opens a TCP connection and handles traffic according to
// the RIP/TCP conventions.
func (s *Server) ListenAndServeTCP(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	s.listener = listener
	return s.acceptConnectionsTCP()
}

func (s *Server) ServeTCP(conn net.Conn) (err error) {
	req, e := riptcp.ReadRequest(conn)
	if e != nil {
		return riptcp.SendResponse(conn, proto.ResponseError(e))
	}
	return riptcp.SendResponse(conn, s.Handle(req))
}

func (s *Server) acceptConnectionsTCP() (err error) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		go s.handleConnectionTCP(conn)
	}
}

func (s *Server) handleConnectionTCP(conn net.Conn) (err error) {
	for {
		if err = s.ServeTCP(conn); err != nil {
			return
		}
	}
}
