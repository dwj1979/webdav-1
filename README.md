# WebDAV #

This is a tiny WebDAV server written in Go, suited for exposing a file system, optionally protected by basic authentication.

## Usage ##

    -dir string
          directory to serve (default "./")
    -listen string
          bind address to listen on (default ":8080")
    -password string
          basic authentication password
    -user string
          basic authentication user name
