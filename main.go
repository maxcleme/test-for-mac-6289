package main

import (
	"net/http"

	"github.com/maxcleme/go-ieproxy"
)

func main() {
	r, err := http.NewRequest("POST", "https://sessions.bugsnag.com", nil)
	if err != nil {
		panic(err)
	}

	if _, err := ieproxy.GetProxyFunc()(r); err != nil {
		panic(err)
	}
}
