package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/xabi93/guards-calendar/cmd/guards-calendar/server"
)

func main() {
	if err := Run(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func Run(out io.Writer) error {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	port := env("PORT", "3000")

	fmt.Fprintf(out, "listening on %s:%s", host, port)

	return http.ListenAndServe(net.JoinHostPort(host, port), server.New())
}

func env(key, val string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return val
	}
	return v
}
