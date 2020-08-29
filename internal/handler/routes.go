package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitPagesRouter(router *httprouter.Router) {
	// router.GET()
}

func InitAdminRouter(router *httprouter.Router) {
	// router.GET()
}

func InitPostRouter(router *httprouter.Router) {
	// router.GET()
}

func InitStaticFiles(router *http.Router) {
	// Serve static files for Assets
	router.ServeFiles("/styles/*filepath", http.Dir(fmt.Sprintf("%v/styles/", "nanti")))
	router.ServeFiles("/scripts/*filepath", http.Dir(fmt.Sprintf("%v/scripts/", "nanti")))
	router.ServeFiles("/img/*filepath", http.Dir(fmt.Sprintf("%v/img/", "nanti")))
	router.ServeFiles("/fonts/*filepath", http.Dir(fmt.Sprintf("%v/fonts/", "nanti")))
}
