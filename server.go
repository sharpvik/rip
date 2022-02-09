package rfip

import "net"

type Server struct {
	Resolver
	listener net.Listener
}

func NewServer(rsvr Resolver) *Server {
	return &Server{
		Resolver: rsvr,
	}
}

func NewServerWithResolver(master interface{}) *Server {
	return NewServer(NewResolver(master))
}

func (s *Server) ListenAndServe(addr string) (err error) {
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
	req, err := ReadRequest(conn)
	if err != nil {
		return
	}
	resp, err := s.Resolve(req)
	if err != nil {
		return
	}
	_, err = resp.Send(conn)
	return
}
