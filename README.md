#go-http-echo

A simple golang HTTP server to dump incoming requests
It will reate a file per request and respond with `200 OK`

## How to use

    go run main.go

This will start the HTTP server running on localhost, port 7893. 
To run on a different port, set the `SERVER_PORT` environment variable before running.

Once running, you can make requests from a browser (or any HTTP client) via `http://localhost:7893/`
