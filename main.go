package main

import (
	"encoding/json"
	"fmt"
	"github.com/viktorbenei/go-webtest/vendor/ansi"
	"log"
	"net/http"
)

var (
	ansiColorRed   = ansi.ColorFunc("red")
	ansiColorGreen = ansi.ColorFunc("green")
)

func init() {

}

type HelloResp struct {
	Msg string `json:"msg"`
}

// func writeJSONError(w http.ResponseWriter, err error) {
// 	w.Header().Set("Content Type", "application/json")
// 	w.WriteHeader(errorStatusCode(err))
// 	if jsonErr := json.NewEncoder(w).Encode(map[string]string{
// 		"description": err.Error(),
// 		"error":       errorName(err, "error"),
// 	}); jsonErr != nil {
// 		log.Println(ansiColorRed("Error: "), jsonErr)
// 	}
// }

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(ansiColorGreen("Request: "), r.URL)

	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)

	hellor := HelloResp{
		Msg: fmt.Sprintf("Hi there!, I love %s!", r.URL.Path[1:]),
	}

	if err := json.NewEncoder(w).Encode(&hellor); err != nil {
		log.Println(ansiColorRed("Error: "), err)
	}
	// fmt.Fprintf(w, "Hi there!, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Ready to server!")
	fmt.Println()
	http.ListenAndServe(":8080", nil)
}
