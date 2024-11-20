package main

import (
	"context"
	"fmt"
	"github.com/apudiu/go-nano-programs/hwmonitor/config"
	"github.com/apudiu/go-nano-programs/hwmonitor/internal/hardware"
	"github.com/apudiu/go-nano-programs/hwmonitor/templates"
	"github.com/coder/websocket"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

type server struct {
	subscriberMessageBuffer int
	subscribers             map[*subscriber]struct{}
	subscriberMu            sync.Mutex
	mux                     http.ServeMux
}

func (s *server) subscriberHandler(w http.ResponseWriter, r *http.Request) {
	err := s.subscribe(r.Context(), w, r)
	if err != nil {
		fmt.Println("Error subscribing:", err)
	}
}

func (s *server) subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var ws *websocket.Conn

	newSubscriber := &subscriber{
		messages: make(chan []byte, s.subscriberMessageBuffer),
	}
	s.addSubscriber(newSubscriber)

	// create ws
	ws, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}
	defer ws.CloseNow()

	ctx = ws.CloseRead(ctx)
	for {
		select {
		case msg := <-newSubscriber.messages:
			ctx2, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()
			if err2 := ws.Write(ctx2, websocket.MessageText, msg); err2 != nil {
				return err2
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *server) addSubscriber(sc *subscriber) {
	// as this is a concurrent situation, we need to safely write to the map
	s.subscriberMu.Lock()
	s.subscribers[sc] = struct{}{}
	s.subscriberMu.Unlock()

	fmt.Println("Added subscriber", sc)
}

func (s *server) broadcast(msg []byte) {
	// again need to be concurrency safe
	s.subscriberMu.Lock()
	defer s.subscriberMu.Unlock()

	// broadcast the message to all subscribers
	for sc := range s.subscribers {
		sc.messages <- msg
	}
}

func newServer() *server {
	s := &server{
		subscriberMessageBuffer: 10,
		subscribers:             make(map[*subscriber]struct{}),
	}

	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.New("index")
		tmpl, err := tmpl.Parse(templates.MainFile)
		if err != nil {
			log.Fatalln("Failed to parse template:", err)
		}

		err2 := tmpl.Execute(w, map[string]any{
			"url": config.Conf.GetWsUrl(),
		})
		if err2 != nil {
			log.Fatalln("Failed to execute template:", err2)
		}
	})

	s.mux.HandleFunc("/ws", s.subscriberHandler)

	return s
}

type subscriber struct {
	messages chan []byte
}

func main() {
	srv := newServer()

	go func(s *server) {
		for {
			sysInfo, err := hardware.GetSystemInfo()
			if err != nil {
				fmt.Println(err)
			}

			cpuInfo, err := hardware.GetCpuInfo()
			if err != nil {
				fmt.Println(err)
			}

			diskInfo, err := hardware.GetDiskInfo()
			if err != nil {
				fmt.Println(err)
			}

			now := time.Now().Format(time.DateTime)
			msg := []byte(`
      			<div hx-swap-oob="innerHTML:#update-timestamp">
        			<p><i style="color: green" class="fa fa-circle"></i> ` + now + `</p>
      			</div>
      			<div hx-swap-oob="innerHTML:#system-data">` + sysInfo + `</div>
      			<div hx-swap-oob="innerHTML:#cpu-data">` + cpuInfo + `</div>
      			<div hx-swap-oob="innerHTML:#disk-data">` + diskInfo + `</div>`,
			)

			s.broadcast(msg)
			time.Sleep(3 * time.Second)
		}
	}(srv)

	log.Fatalln(
		http.ListenAndServe(":"+config.Conf.Port, &srv.mux),
	)
}
