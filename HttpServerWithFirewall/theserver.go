package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	info *log.Logger
)

func initial(traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func validSource(c net.Conn) bool {
	var sourceAddress = c.RemoteAddr()
	info.Println(sourceAddress.String())
	if strings.HasPrefix(sourceAddress.String(), "192.168.1.12") {
		return true
	}
	return true
}

func connStateListener(c net.Conn, cs http.ConnState) {
	if !validSource(c) {
		c.Close()
	}
	//fmt.Printf("CONN STATE: %v, %v\n", cs, c)
}

func main() {

	initial(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.RemoteAddr))
	})

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ConnState:      connStateListener,
	}
	panic(s.ListenAndServe())
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
