package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/archerwq/go-lab/lib/syncutil"
	"github.com/gorilla/mux"
)

var dir = flag.String("dir", ".", "the directory to serve files from. Defaults to the current dir")
var wait = flag.Duration("graceful-timeout", time.Second*15,
	"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

func main() {
	flag.Parse()

	rtr := mux.NewRouter()
	rtr.HandleFunc("/user/{name:[a-z]+}/profile", profile).Methods("GET")
	rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(*dir))))

	rtr.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, _ := route.GetPathTemplate()
		qt, _ := route.GetQueriesTemplates()
		// p will contain regular expression is compatible with regular expression in Perl, Python, and other languages.
		// for instance the regular expression for path '/articles/{id}' will be '^/articles/(?P<v0>[^/]+)$'
		p, _ := route.GetPathRegexp()
		// qr will contain a list of regular expressions with the same semantics as GetPathRegexp,
		// just applied to the Queries pairs instead, e.g., 'Queries("surname", "{surname}") will return
		// {"^surname=(?P<v0>.*)$}. Where each combined query pair will have an entry in the list.
		qr, _ := route.GetQueriesRegexp()
		m, _ := route.GetMethods()
		log.Println(strings.Join(m, ","), strings.Join(qt, ","), strings.Join(qr, ","), t, p)
		return nil
	})

	http.Handle("/", rtr)

	srv := &http.Server{
		Handler: rtr,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening...")

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	syncutil.CtrC()

	// Create a deadline to wait for.
	ctx, _ := context.WithTimeout(context.Background(), *wait)
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func profile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}
