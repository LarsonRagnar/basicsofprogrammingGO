package main

import (
	"sync"
)

type server struct {
	status string
	sync.Mutex
}

func main() {
	s := server{}
	for i := 0; i < 1001; i++ {
		go s.Running()
		go s.Down()
	}

}

func (s *server) Running() {
	s.Lock()
	s.status = "alive"
	s.Unlock()
}

func (s *server) Down() {
	s.Lock()
	s.status = "down"
	s.Unlock()
}
