package function

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func isValidReq(r *http.Request) bool {
	fmt.Fprintf(os.Stdout, "Method: %s; Path: %s; ContentLength: %d\n", r.Method, r.URL.Path, r.ContentLength)
	if r.Method == http.MethodPost && r.Body != nil {
		return true
	}

	return false
}

func MakeNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if !isValidReq(r) {
		http.Error(w, `{"result": 400}`, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	fmt.Fprintf(os.Stdout, "method: %s; url: %s; headers: ", r.Method, r.URL.Path)
	for k, v := range r.Header {
		fmt.Fprintf(os.Stdout, "%s => %s\n", k, v)
	}

	fmt.Fprintf(os.Stdout, "ContentLenght: %d (bytes); Body: ", r.ContentLength)
	io.Copy(os.Stdout, r.Body)
	fmt.Fprintln(os.Stdout, "")

	fmt.Fprintln(w, `{"result": 200}`)
}
