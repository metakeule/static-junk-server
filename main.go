package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
)

var dir *string = flag.String("dir", ".", "directory for files.")
var port *int = flag.Int("port", 8080, "listening port")
var host *string = flag.String("host", "localhost", "listening host")

func tryToLaunch(host string, port int) {
	if port > 8090 {
		fmt.Println("tried more than 10 ports, giving up!")
		return
	}
	h := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("listening on http://%s\n", h)
	err := http.ListenAndServe(h, nil)
	if err != nil {
		fmt.Println(err)
		tryToLaunch(host, port+1)
	}
}

func main() {
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	p := path.Join(wd, *dir)

	http.Handle("/", http.FileServer(http.Dir(p)))
	tryToLaunch(*host, *port)
}
