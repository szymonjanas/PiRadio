package main

import (
    "html/template"
    "net/http"
    "log"
    "fmt"
)

type StationsName struct {
    Name string
}

type StationsPageData struct {
    Stations []StationsName
}

type PlayingDetails struct {
    Station string
    Title string
    Url string
}

type ViewPageData struct {
    StationsData StationsPageData
    PlayingData PlayingDetails
}

func getStations() []StationsName {
    return [] StationsName {
        StationsName {"Heart"}, 
        StationsName {"RMF-FM"},
        StationsName {"RADIO-ZET"},
        StationsName {"Anty-Radio"},
    } 
}

func viewHandler (w http.ResponseWriter, r *http.Request){
    tmpl, err := template.ParseFiles("server.html")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    details := ViewPageData {
        StationsData: StationsPageData{
            Stations: getStations(),
        },
        PlayingData: PlayingDetails {
            Station: "station name",
            Title: "title name",
            Url: "url link",
        },
    }

    tmpl.Execute(w, details)
}

func playHandler(w http.ResponseWriter, r *http.Request){
    http.Redirect(w, r, "/radio/", http.StatusFound)
    fmt.Println("PLAY")
} 

func stopHandler(w http.ResponseWriter, r *http.Request){
    http.Redirect(w, r, "/radio/", http.StatusFound)
    fmt.Println("STOP")
} 


func setHandler(w http.ResponseWriter, r *http.Request){
    body := r.FormValue("Stations")
    http.Redirect(w, r, "/radio/", http.StatusFound)
    fmt.Println(body)
} 

func main() {
    http.HandleFunc("/radio/play", playHandler)
    http.HandleFunc("/radio/stop", stopHandler)
    http.HandleFunc("/radio/set", setHandler)
    http.HandleFunc("/radio/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
