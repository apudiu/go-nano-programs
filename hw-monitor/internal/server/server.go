package server

import (
	"context"
	"fmt"
	"github.com/apudiu/go-nano-programs/hwmonitor/config"
	"github.com/apudiu/go-nano-programs/hwmonitor/templates"
	"github.com/coder/websocket"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	subscriberMessageBuffer int
	subscribers             map[*Subscriber]struct{}
	subscriberMu            sync.Mutex
	Mux                     http.ServeMux
}

func (s *Server) subscriberHandler(w http.ResponseWriter, r *http.Request) {
	err := s.subscribe(r.Context(), w, r)
	if err != nil {
		fmt.Println("Error subscribing:", err)
	}
}

func (s *Server) subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var ws *websocket.Conn

	newSubscriber := &Subscriber{
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

func (s *Server) addSubscriber(sc *Subscriber) {
	// as this is a concurrent situation, we need to safely write to the map
	s.subscriberMu.Lock()
	s.subscribers[sc] = struct{}{}
	s.subscriberMu.Unlock()

	fmt.Println("Added Subscriber", sc)
}

func (s *Server) Broadcast(msg []byte) {
	// again need to be concurrency safe
	s.subscriberMu.Lock()
	defer s.subscriberMu.Unlock()

	// broadcast the message to all subscribers
	for sc := range s.subscribers {
		sc.messages <- msg
	}
}

func New() *Server {
	s := &Server{
		subscriberMessageBuffer: 10,
		subscribers:             make(map[*Subscriber]struct{}),
	}

	s.Mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
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

	s.Mux.HandleFunc("GET /ws", s.subscriberHandler)

	return s
}

type Subscriber struct {
	messages chan []byte
}
