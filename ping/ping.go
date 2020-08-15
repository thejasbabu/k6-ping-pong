package ping

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Pinger struct {
	pongEndpoint string
	listenerPort int
	httpClient   *http.Client
}

func NewPinger(endpoint string, port int) Pinger {
	defaultRoundTripper := http.DefaultTransport
	defaultTransport := defaultRoundTripper.(*http.Transport)
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	client := &http.Client{Transport: defaultTransport}

	return Pinger{pongEndpoint: endpoint, listenerPort: port, httpClient: client}
}

// Start the pinger and bind to listener port
func (p *Pinger) Start() {
	http.HandleFunc("/ping", p.handler)
	listen := fmt.Sprintf(":%d", p.listenerPort)
	fmt.Printf("Pinger listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}

func (p *Pinger) handler(w http.ResponseWriter, r *http.Request) {
	apiEndpoint := fmt.Sprintf("%s/pong", p.pongEndpoint)
	resp, err := p.httpClient.Get(apiEndpoint)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error connecting pong service"))
		return
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte("Non-successful status recieved"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
