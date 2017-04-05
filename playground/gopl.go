package main

import (
	"fmt"
	"github.com/master-g/golandlord/playground/bp"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func bluePay(w http.ResponseWriter, r *http.Request) {
	keys := []string{
		"cmd",
		"bt_id",
		"msisdn",
		"operator",
		"paytype",
		"price",
		"productid",
		"status",
		"t_id",
		"currency",
		"interfacetype",
		"encrypt",
	}

	queries := r.URL.Query()

	for _, v := range keys {
		fmt.Println(v, queries.Get(v))
	}
	io.WriteString(w, "proceed")
}

func main() {
	handler := new(bp.Handler)
	handler.RegisterRouter("/", hello)
	handler.RegisterRouter("/BlueNotification", bluePay)
	server := http.Server{
		Addr:    ":8000",
		Handler: handler,
	}
	server.ListenAndServe()
}
