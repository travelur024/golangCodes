package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/API", server.AddMidleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()

}
