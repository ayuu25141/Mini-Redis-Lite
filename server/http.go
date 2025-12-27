package server

import (
	"encoding/json"
	"net/http"

	"Minireddis/storage"
)

type Server struct {
	Store *storage.Store
}

func StartHTTP(addr string, store *storage.Store) {
	s := &Server{Store: store}

	http.HandleFunc("/set", s.setHandler)
	http.HandleFunc("/get", s.getHandler)

	println("🚀 HTTP API running on", addr)
	http.ListenAndServe(addr, nil)
}
func (s *Server) setHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	var body struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	s.Store.Set(body.Key, body.Value)

	w.Write([]byte("OK"))
}
func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	val, ok := s.Store.Get(key)
	if !ok {
		w.Write([]byte("nil"))
		return
	}

	w.Write([]byte(val))
}
