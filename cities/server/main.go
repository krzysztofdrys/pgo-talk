package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/krzysztofdrys/pgo-talk/cities/city"
	"gitlab.com/golang-commonmark/markdown"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func main() {
	dataset, err := city.Read()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/distance", func(w http.ResponseWriter, r *http.Request) {
		latStr := r.URL.Query().Get("lat")
		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "lat: %q", err)
			return
		}

		lngStr := r.URL.Query().Get("lng")
		lng, err := strconv.ParseFloat(lngStr, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "lng: %q", err)
			return
		}

		distStr := r.URL.Query().Get("dist")
		dist, err := strconv.ParseFloat(distStr, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "distStr: %q", err)
			return
		}

		result := city.Filter(dataset.Features, city.DistanceFilter{
			Lat:      lat,
			Lng:      lng,
			Distance: dist,
		})

		json.NewEncoder(w).Encode(result)
	})
	http.HandleFunc("/render", render)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func render(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	src, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	md := markdown.New(
		markdown.XHTMLOutput(true),
		markdown.Typographer(true),
		markdown.Linkify(true),
		markdown.Tables(true),
	)

	var buf bytes.Buffer
	if err := md.Render(&buf, src); err != nil {
		log.Printf("error converting markdown: %v", err)
		http.Error(w, "Malformed markdown", http.StatusBadRequest)
		return
	}

	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
