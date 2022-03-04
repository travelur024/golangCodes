package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}
func (r *Router) handler(w http.ResponseWriter, request *http.Request) {
	handler, exist := r.FindHandler(request.URL.Path)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error brooo")
		return
	}
	handler(w, request)
}
func handler02(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saludos, bienvenido 002 a %s!", r.URL.Path[1:])
}
func handler03(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Saludos, bienvenido 003 a %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "Saludos, Naty! Chiquita hermosa bonita preciosa bb, Me encantas!!!")

}

func middleware(originalHandler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flag := false
		fmt.Println("Checking authentication")
		if flag {
			originalHandler(w, r)
		} else {
			fmt.Fprintf(w, "Authentication false")
			fmt.Println("Authentication false")
			return
		}
	})
}

func allMiddlewares(f http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	cont := 0
	for _, mid := range middlewares {
		f = mid(f)
		cont++
	}
	fmt.Println("Paso por: ", cont, "Middleware:", f)
	cont = 0
	return f
}

func main() {
	rute := Router{
		rules: make(map[string]http.HandlerFunc),
	}
	rute.rules["/end"] = handler03
	autenticado := allMiddlewares(handler02, middleware)
	rute.rules["/home"] = autenticado

	http.HandleFunc("/", rute.handler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("No se inicializo el servidor")
	}

}
