package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/archerwq/go-lab/application/webfront/rule"
)

type Server struct {
	mu    sync.RWMutex // guards the fields bellow
	mtime time.Time    // when the rule file was last modified
	rules []*rule.Rule
}

// NewServer construct a Server that reads rules from file with a peroid specified by poll.
func NewServer(file string, poll time.Duration) (*Server, error) {
	s := new(Server)
	if err := s.loadRules(file); err != nil {
		return nil, err
	}
	go s.refreshRules(file, poll)
	return s, nil
}

// Implements http.Handler interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h := s.handler(r); h != nil {
		h.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Not found.", http.StatusNotFound)
}

// refreshRules polls file periodically and refresh the Server's rules and last modified time.
func (s *Server) refreshRules(file string, poll time.Duration) {
	for {
		if err := s.loadRules(file); err != nil {
			log.Println(err)
		}
		time.Sleep(poll)
	}
}

// Load rules from file if it has been modified.
func (s *Server) loadRules(file string) error {
	fi, err := os.Stat(file)
	if err != nil {
		return err
	}

	modTime := fi.ModTime()
	if !modTime.After(s.mtime) && s.rules != nil {
		return nil // no change
	}

	rules, err := parseRules(file)
	if err != nil {
		return fmt.Errorf("parsing %s: %v", file, err)
	}

	s.mu.Lock()
	s.mtime = modTime
	s.rules = rules
	defer s.mu.Unlock()

	return nil
}

func parseRules(file string) ([]*rule.Rule, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var rules []*rule.Rule
	err = json.NewDecoder(f).Decode(&rules)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func (s *Server) handler(req *http.Request) http.Handler {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, r := range s.rules {
		if r.Match(req) {
			return r.Handler()
		}
	}
	return nil
}
