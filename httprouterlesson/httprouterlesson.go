package httprouterlesson

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func httpRouterLesson() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r * http.Request, params httprouter.Params){
		fmt.Fprintln(w, "halo dari golang")
	})
	server := http.Server{
		Handler: router,
		Addr: "localhost:8080",
	}
	server.ListenAndServe()
}