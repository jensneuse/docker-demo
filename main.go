package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
)

var (
	PORT          = os.Getenv("PORT")
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""\___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
         \    \        __/
          \____\______/

          |          |
       __ |  __   __ | _  __   _
      /  \| /  \ /   |/  / _\ |
      \__/| \__/ \__ |\_ \__  |`))
	})

	fmt.Printf("Listening on: %s\n",PORT)
	err := http.ListenAndServe("0.0.0.0:" + PORT,http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}
