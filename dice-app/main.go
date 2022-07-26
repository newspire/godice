package main

import (
	"fmt"
	"log"
	"net/http"

	"newspire.org/die"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<A HREF=\"./roll\">Roll</A>")
	})

	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		Die := die.NewDie(6)
		Die.Roll()
		fmt.Fprintf(w, "%d", Die.Value())
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
