package route

import "github.com/gorilla/mux"

func NewRoute()*mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	return myRouter
}