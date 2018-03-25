package main

import (
	"fmt"
	"net/http"

	"github.com/m-nakamura145/golang-docker-compose-sample/entity"
)

func hello_world(w http.ResponseWriter, r *http.Request) {
	sampleEntity := entity.Sample{"hoge", "fuga"}
	fmt.Fprintf(w, "Hello World. User: "+sampleEntity.User+" Password: "+sampleEntity.Password)
}

func main() {
	http.HandleFunc("/", hello_world)
	http.ListenAndServe(":5000", nil)
}
