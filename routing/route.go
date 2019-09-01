package route

import (
	"github.com/fdirlikli/kube-ui/kubernetes"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"net/http"
)

var r *httprouter.Router
var renderer *render.Render

func GetRouter() *httprouter.Router {
	renderer = render.New()
	if r == nil {
		r = httprouter.New()
		initRoutes()
	}
	return r
}

func initRoutes() {
	r.GET("/", HomeHandler)
	r.GET("/createpod", createPod)
	r.GET("/deletepod", deletePod)
	r.GET("/getpods/:namespace", getPods)
	r.GET("/getpods", getPods)
	r.GET("/getnamespaces", getNamespaces)
}

func getPods(responseWriter http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ns := p.ByName("namespace")
	if ns == "all" {
		ns = ""
	}
	renderer.JSON(responseWriter, http.StatusOK, kubernetes.Service.GetAllPods(&ns))
}

func getNamespaces(responseWriter http.ResponseWriter, r *http.Request, p httprouter.Params) {
	renderer.JSON(responseWriter, http.StatusOK, kubernetes.Service.GetNamespaces())
}

func deletePod(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	keys, ok := r.URL.Query()["podname"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
	}
	defer handlePanic(rw)
	kubernetes.Service.DeletePod(&keys[0])
}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	renderer.HTML(rw, http.StatusOK, "example", "hede")
}

func createPod(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	podname, ok := r.URL.Query()["podname"]
	namespace, ok2 := r.URL.Query()["namespace"]
	if !ok || !ok2 {
		rw.WriteHeader(http.StatusBadRequest)
	}
	defer handlePanic(rw)
	pod := kubernetes.Service.CreatePod(&podname[0], &namespace[0])
	renderer.JSON(rw, http.StatusCreated, pod)

}

func handlePanic(rw http.ResponseWriter) {
	if r := recover(); r != nil {
		renderer.JSON(rw, http.StatusInternalServerError, r)
	}
}
