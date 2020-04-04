package scf_apigw_wrap

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Debug struct {
	mux  *http.ServeMux
	Host string
	Port int
}

func (d Debug) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		method := strings.ToUpper(r.Method)
		log.Printf("[%s] %s\n", method, r.RequestURI)
		body, err := ioutil.ReadAll(r.Body)
		defer func() { _ = r.Body.Close() }()
		if err != nil {
			fmt.Printf("\033[31m%s\033[0m\n", fmt.Sprint(err))
			w.WriteHeader(200)
			return
		}
		if len(body) > 0 {
			fmt.Println(string(body))
		}
		w.WriteHeader(200)
	})
	log.Println("listen", d.Port)
	return http.ListenAndServe(":"+strconv.Itoa(d.Port), mux)
}

func (d Debug) SendByte(data []byte) error {
	url := fmt.Sprintf("http://%s:%d", d.Host, d.Port)
	resp, err := http.Post(url, "application/octet-stream", bytes.NewReader(data))
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func (d Debug) SendString(data string) error {
	return d.SendByte([]byte(data))
}
