package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServerHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Bye"))
}
