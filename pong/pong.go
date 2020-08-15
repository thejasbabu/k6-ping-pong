package pong

import (
	"fmt"
	"log"
	"net/http"
)

type Ponger struct {
	listenerPort int
}

func NewPonger(port int) Ponger {
	return Ponger{listenerPort: port}
}

// Start the ponger and bind to listener port
func (p *Ponger) Start() {
	http.HandleFunc("/pong", p.handler)
	listen := fmt.Sprintf(":%d", p.listenerPort)
	fmt.Printf("Ponger listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}

func (p *Ponger) handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}
