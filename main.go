package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

var quotes []Quote

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "ERROR", http.StatusInternalServerError)
	}

	tmpl.Execute(w, nil)
}

func Getfile() {
	res, err := os.Open("db/quote.json")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	decoder := json.NewDecoder(res)
	if err := decoder.Decode(&quotes); err != nil {
		log.Fatal(err)
	}
}

func getquote() Quote {
	if len(quotes) == 0 {
		return Quote{
			ID:     0,
			Author: "Sistem",
			Text:   "Tidak ada kutipan tersedia.",
		}
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(quotes))
	return quotes[index]
}

func handleRandomQuote(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Quote-Key")
	if key != "iowehniuvht4iuhv8t3489tv82pqniwqvt8q4yvtiqyn8tn4yt83nt" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	quote := getquote()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("static/js"))))
	Getfile()
	http.HandleFunc("/quote", handleRandomQuote)
	http.ListenAndServe(":8080", nil)
}
