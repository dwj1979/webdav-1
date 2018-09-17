package main

import (
	"flag"
	"golang.org/x/net/webdav"
	"log"
	"net/http"
)

func main() {
	var (
		listen, dir, username, password string
		srv                             http.Handler
	)

	flag.StringVar(&listen, "listen", ":8080", "bind address to listen on")
	flag.StringVar(&dir, "dir", "./", "directory to serve")
	flag.StringVar(&username, "user", "", "basic authentication user")
	flag.StringVar(&password, "password", "", "basic authentication password")
	flag.Parse()

	srv = &webdav.Handler{
		FileSystem: webdav.Dir(dir),
		LockSystem: webdav.NewMemLS(),
	}

	if len(username) > 0 && len(password) > 0 {
		handler := srv
		srv = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, _ := r.BasicAuth()
			if user != username || pass != password {
				w.Header().Set("WWW-Authenticate", "Basic realm=\"WebDAV\"")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			handler.ServeHTTP(w, r)
		})
	}

	log.Fatalln(http.ListenAndServe(listen, srv))
}
