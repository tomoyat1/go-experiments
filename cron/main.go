package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() { fmt.Println("foo")})
	c.Start()

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

