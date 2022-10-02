package createtcpwithgo

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	// 해당 Port listen
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		t.Fatal(err)
	}

	// graceful close
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())

	for {
		conn, err := listener.Accept()

		// Faile TCP HandShake OR Already Close Listener
		if err != nil {
			t.Fatal(err)
		}

		// goroutine
		go func(c net.Conn) {
			defer c.Close() // gracefull termination

			// TCP Logic
		}(conn)
	}
}
