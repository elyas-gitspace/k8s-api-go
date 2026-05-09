package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"
)

// Item représente un objet dans notre liste
type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// Notre "base de données" en mémoire
var (
    items  = []Item{}
    nextID = 1
    mu     sync.Mutex
)

// GET /health
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// GET /items
func getItemsHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// POST /items
func createItemHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    var item Item
    json.NewDecoder(r.Body).Decode(&item)
    item.ID = nextID
    nextID++
    items = append(items, item)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            getItemsHandler(w, r)
        } else if r.Method == http.MethodPost {
            createItemHandler(w, r)
        }
    })

    log.Println("API démarrée sur le port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}