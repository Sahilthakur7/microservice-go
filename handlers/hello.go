package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {

}

func (h* Hello) ServeHTTP(rw http.ResponseWriter , r *http.Request){
	log.Println("hello world")
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Error", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", data)
}