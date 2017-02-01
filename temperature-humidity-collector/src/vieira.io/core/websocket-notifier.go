package core

import (
	"io"
	"net"
	"log"
	"time"
	"net/http"
	"sync"
	"fmt"
	"github.com/googollee/go-socket.io"
	"encoding/json"
)

func listenAndServeWithClose(addr string, handler http.Handler) (sc io.Closer, err error) {
	var listener net.Listener
	srv := &http.Server{Addr: addr, Handler: handler}

	if addr == "" {
		addr = ":http"
	}

	listener, err = net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	go func() {
		err := srv.Serve(tcpKeepAliveListener{listener.(*net.TCPListener)})
		if err != nil {
			log.Println("HTTP Server Error - ", err)
		}
	}()

	return listener, nil
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

const measurementsChannel = "temperature-measurements"

func webSocketNotifier(listenPort int, measures <-chan Measure, trap <-chan bool, wg *sync.WaitGroup) func() {
	return func() {
		server, err := socketio.NewServer(nil)
		if err != nil { log.Fatalln("Socket.IO server start error", err) }
		server.On("connection", func(so socketio.Socket) {
			so.Join(measurementsChannel)
		})
		http.Handle("/socket.io/", server)

		listeningAddress := fmt.Sprintf(":%d", listenPort)
		srvCloser, err := listenAndServeWithClose(listeningAddress, nil)
		if err != nil { log.Fatalln("Socket.IO server start error", err) }
		log.Printf("Socket.IO server started on %s", listeningAddress)

		for {
			select {
			case <-trap:
				log.Println("Stopping Socket.IO server routine...")
				srvCloser.Close();
				if err != nil { log.Panicln("Socket.IO server stop error", err) }
				wg.Done()
				log.Println("Socket.IO server routine stopped.")
			case measure := <-measures:
				broadcastMeasure(measure, *server)
			}
		}

	}
}

func broadcastMeasure(measure Measure, server socketio.Server) {
	bytes, _ := json.Marshal(measure)
	log.Printf("Received payload : %s", string(bytes))
	server.BroadcastTo(measurementsChannel, "new-measure", string(bytes))
}