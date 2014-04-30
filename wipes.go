package main

import (
	"bufio"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var version = "0.0.1"

type conn struct {
	ws  *websocket.Conn
	out chan string
}

type server struct {
	broadcast   chan string
	register    chan *conn
	unregister  chan *conn
	connections map[*conn]bool
}

func (c *conn) readIgnoreLoop() {
	for {
		// read, but ignore.
		_, _, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func (c *conn) writeLoop() {
	for message := range c.out {
		err := c.ws.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func (s *server) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Whoops. Try using websockets!", 400)
		return
	} else if err != nil {
		return
	}

	c := &conn{out: make(chan string, 256), ws: ws}

	s.register <- c
	defer func() { s.unregister <- c }()
	go c.writeLoop()
	c.readIgnoreLoop()
}

func (s *server) pipeInput(r *bufio.Reader) {
	for {
		if line, err := r.ReadString('\n'); err == nil {
			s.broadcast <- line
		} else {
			// On EOF, we just exit.
			os.Exit(0)
		}
	}
}

func (s *server) run() {
	for {
		select {
		case c := <-s.register:
			s.connections[c] = true

		case c := <-s.unregister:
			delete(s.connections, c)
			close(c.out)

		case msg := <-s.broadcast:
			for c, _ := range s.connections {
				c.out <- msg
			}
		}
	}
}

var (
	addr       = flag.String("addr", ":8080", "http service address")
	staticPath = flag.String("static", ".", "static file path")
	versionPrint    = flag.Bool("v", false, "print the version and exit")
)

func main() {
	flag.Parse()

	if *versionPrint {
		println(version)
		os.Exit(0)
	}

	s := &server{broadcast: make(chan string), register: make(chan *conn), unregister: make(chan *conn), connections: make(map[*conn]bool)}

	go s.pipeInput(bufio.NewReader(os.Stdin))
	go s.run()

	http.Handle("/", http.FileServer(http.Dir(*staticPath)))
	http.HandleFunc("/_ws", s.wsHandler)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
