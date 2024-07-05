package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)

var DataMap map[string]string

func init() {
    file, err := ioutil.ReadFile("data.json")
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(file, &DataMap)
    if err != nil {
        log.Fatal(err)
    }
}

func handleData(w http.ResponseWriter, r *http.Request) {
    jsonData, err := json.Marshal(DataMap)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}

func handleDataCode(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    data, exists := DataMap[code]
    if !exists {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Data for code %s: %s", code, data)
}

func main() {
    http.HandleFunc("/datas", handleData)
    http.HandleFunc("/data", handleDataCode)

    fmt.Println("Server listening on port 80")
    log.Fatal(http.ListenAndServe(":80", nil))
}
