package server

import (
	"bytes"
	"fmt"
	"net"

	"github.com/saurabhy27/redis-database/constants"
	"github.com/saurabhy27/redis-database/processor"
	request "github.com/saurabhy27/redis-database/request"
)

type ServerArgs struct {
	Port int
}

type Server struct {
	args             ServerArgs
	requestProcessor processor.RequestProcessorInterface
}

func NewServer(args ServerArgs, requestProcessor processor.RequestProcessorInterface) *Server {
	return &Server{args: args, requestProcessor: requestProcessor}
}

func (s *Server) Start() {
	fmt.Println("Starting connections....")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.args.Port))
	if err != nil {
		fmt.Println(fmt.Errorf("FAILED TO LISTEN TO ADDRESS %s: %w", fmt.Sprintf(":%d", s.args.Port), err))
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(fmt.Errorf("FAILED TO GET CONNECTION: %w", err))
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection Created")
	for {
		conn.Write([]byte("redis> "))
		buf := make([]byte, constants.ArgBufSize)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(fmt.Errorf("UNABLE TO READ DATA FROM TERMINAL: %w", err))
			return
		}
		data := buf[:n]
		// removing the \n from the end of the string
		if bytes.HasSuffix(data, []byte("\n")) {
			data = data[:len(data)-1]
		}
		fmt.Printf("Received %d bytes: %s\n", n, data)

		request, err := request.ParseProtocol(string(data))
		if err != nil {
			fmt.Println(fmt.Errorf("FAILED TO PARSE INPUT: %w", err))
			s.writeError(err, conn)
			continue
		}
		response, err := s.requestProcessor.Process(request)
		if err != nil {
			fmt.Println(fmt.Errorf("FAILED TO EXECUTE THE REQUEST: %w", err))
			s.writeError(err, conn)
			continue
		}
		s.writeSuccess(response.Value, conn)
	}
}

func (s *Server) writeError(err error, conn net.Conn) {
	conn.Write([]byte(fmt.Sprintf("ERR %s\n", err)))
}

func (s *Server) writeSuccess(value any, conn net.Conn) {
	switch v := value.(type) {
	case int:
		conn.Write([]byte(fmt.Sprintf("%d\n", v)))
	case string:
		conn.Write([]byte(fmt.Sprintf("%s\n", v)))
	case []string:
		for _, s := range v {
			conn.Write([]byte(fmt.Sprintf("%v\n", s)))
		}
	case map[float64]string:
		for k, v := range v {
			conn.Write([]byte(fmt.Sprintf("%v  %f\n", v, k)))
		}
	case nil:
		conn.Write([]byte("(nil)\n"))
	case []byte:
		if len(v) == 0 {
			conn.Write([]byte("(nil)\n"))
		} else {
			conn.Write([]byte(fmt.Sprintf("%s\n", v)))
		}
	default:
		fmt.Println("Type is unknown!")
	}
}
