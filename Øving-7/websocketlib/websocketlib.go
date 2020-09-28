package websocketlib

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"net"
	"net/http"
	"strings"
)

var clients []net.Conn

// MakeSocket : initiates a websocket
func MakeConn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		return nil, errors.New("http.ResponseWriter not of type http.Hijacker")
	}

	conn, rw, err := hj.Hijack()
	if err != nil {
		return nil, err
	}

	err = Handshake(rw.Writer, r)
	if err != nil {
		return nil, err
	}

	clients = append(clients, conn)

	return conn, nil
}

// Handshake : executes an handshake with client
func Handshake(bw *bufio.Writer, r *http.Request) error {
	hash := func(key string) string {
		h := sha1.New()
		h.Write([]byte(key))
		h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))

		return base64.StdEncoding.EncodeToString(h.Sum(nil))
	}(r.Header.Get("Sec-WebSocket-Key"))

	lines := []string{
		"HTTP/1.1 101 Web Socket Protocol Handshake",
		"Server: go/echoserver",
		"Upgrade: WebSocket",
		"Connection: Upgrade",
		"Sec-WebSocket-Accept: " + hash,
		"",
		"", // required for extra CRLF
	}

	_, err := bw.Write([]byte(strings.Join(lines, "\r\n")))
	return err
}

// ReadFromClient : reads message from client
func ReadFromClient(conn net.Conn) ([]byte, error) {
	msg := make([]byte, 1024)
	_, err := conn.Read(msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// MessageAllClients : sends message to all connected clients
func MessageAllClients(conn net.Conn, msg []byte) error {
	for _, client := range clients {
		if client.RemoteAddr() == conn.RemoteAddr() {
			_, err := client.Write([]byte("Message sent to all clients"))
			if err != nil {
				return err
			}
		} else {
			_, err := client.Write(msg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
