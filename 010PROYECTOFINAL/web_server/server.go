//el archivo main va a poder leer todo lo que este en el
//archivo server
package main

import "net/http"

//Struct server
type Server struct {
	//puerto del servidor para escuchar las conexiones
	port string
	//se agrega el atributo router que es un apuntador
	// al struct Router de router.go
	router *Router
}

//Funcion tipo global para ser leida en otros archivos
//Sirve para instanciar el servidor y que sea capaz de
// escuchar las conexiones recive el puerto que tiene que
// estar escuchando y devuelve el servidor como tal
func NewServer(port string) *Server {
	return &Server{
		port: port,
		//Router instanciado
		router: NewRouter(),
		//con esto el servidor ya es capaz de instanciar
		// el router y de tenerlo como propiedad
	}
}

//Neseria para que escuche peticiones...
//Funcion tipo receiver, del struct Server, devuelve un
//error en caso de que haya problemas al conectar
func (s *Server) Listen() error {
	//el router va a ser el encargado de tomar las urls y
	//procesarlas como se debe, crea el entry-point
	// los parametro son: el slash que es el punto de entrada,
	//y el handler es el router recien creado
	http.Handle("/", s.router)
	//ListenAndServe listens on the TPC network addres addre
	//and then calls Serve with handler to handle request on
	//incoming connections
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return nil
	}
	return nil
}
