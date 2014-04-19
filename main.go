package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
)

var dir *string = flag.String("dir", ".", "directory for files.")
var port *int = flag.Int("port", 8080, "Port to listen for requests on.")

func tryToLaunch(port int) {
	if port > 8090 {
		fmt.Println("tried more than 10 ports, giving up!")
		return
	}
	fmt.Printf("listening on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
		tryToLaunch(port + 1)
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
	tryToLaunch(*port)
}
