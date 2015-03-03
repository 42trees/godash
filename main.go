package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Woo struct {
	Id   int
	Name string
}

func main() {
	// Simple static webserver:
	//log.Fatal(http.ListenAndServe(":8082", http.FileServer(http.Dir("."))))
	mux := ApiMux()
	log.Println("Listening on port 8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}

func ApiMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./js"))))
	mux.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./img"))))

	mux.HandleFunc("/woo", woo)

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, nil)
	})
	return mux
}

func woo(w http.ResponseWriter, req *http.Request) {

	// Make sure that the writer supports flushing.
	//
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	var woo Woo
	woo.Id = 5
	woo.Name = "YEAH"

	response, _ := json.Marshal(woo)

	//
	for {

		// Write to the ResponseWriter, `w`.
		fmt.Fprintf(w, "data: %s\n\n", response)

		// Flush the response.  This is only possible if
		// the repsonse supports streaming.
		f.Flush()
		time.Sleep(30 * time.Second)
	}

	//	fmt.Fprintf(w, "%s", response)
}
