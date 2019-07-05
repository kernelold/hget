package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", headers)
	http.HandleFunc("/test/", test)
	if os.Getenv("HEADERGET_IP") == "" {
		os.Setenv("HEADERGET_IP", "0.0.0.0")
	}
	if os.Getenv("HEADERGET_PORT") == "" {
		os.Setenv("HEADERGET_PORT", "8080")
	}
	bind := fmt.Sprintf("%s:%s", os.Getenv("HEADERGET_IP"), os.Getenv("HEADERGET_PORT"))
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func test(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "OK \n%s \n", req.UserAgent())
}

func headers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	//Iterate over all header fields
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
	//Get value for a specified token
	fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])
}
