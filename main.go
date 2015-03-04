package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Widget struct {
	Id       int
	Title    string
	Subtitle string
	Content  string
	Color    string
	Time     string
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

	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Now()
	now := t.Format(layout)

	var wd Widget
	wd.Id = 1
	wd.Title = "y7.mediasuite.proc"
	wd.Subtitle = "OK"
	wd.Content = "Everything is fine"
	wd.Color = "green"
	wd.Time = now
	var widgets []Widget

	widgets = append(widgets, wd)
	wd.Id = 2
	wd.Subtitle = "OH NO. STUFF IS BROKEN"
	wd.Content = "Everything is definitely not fine"
	wd.Color = "red"
	wd.Time = now
	widgets = append(widgets, wd)

	wd.Id = 3
	wd.Subtitle = "OH NO. STUFF IS BROKEN"
	wd.Content = "Everything is definitely not fine"
	wd.Color = "red"
	wd.Time = now
	widgets = append(widgets, wd)

	wd.Id = 4
	wd.Subtitle = "OH NO. STUFF IS AWESOME"
	wd.Content = "Everything is definitely not fine"
	wd.Color = "green"
	wd.Time = now
	widgets = append(widgets, wd)

	wd.Id = 5
	wd.Subtitle = "OH NO. STUFF IS AWESOME"
	wd.Content = "Everything is definitely not fine"
	wd.Color = "orange"
	wd.Time = now
	widgets = append(widgets, wd)

	wd.Id = 6
	wd.Subtitle = "PUPPIES"
	wd.Content = "Everything is AWESOME"
	wd.Color = "purple"
	wd.Time = now
	widgets = append(widgets, wd)

	response, _ := json.Marshal(widgets)

	//
	for {

		// Write to the ResponseWriter, `w`.
		fmt.Fprintf(w, "data: %s\n\n", response)

		// Flush the response.  This is only possible if
		// the repsonse supports streaming.
		f.Flush()
		time.Sleep(5 * time.Second)
	}

	//	fmt.Fprintf(w, "%s", response)
}
