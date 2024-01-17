package main

import (
	"fmt"
	"net/http"
	"openweb/views"
	"openweb/views/components"

	"github.com/a-h/templ"
)

func main() {
	component := views.Index()
	component2 := components.Foo()

	http.Handle("/", templ.Handler(component))
	http.Handle("/foo", templ.Handler(component2))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
