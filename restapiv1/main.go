package main

func main() {
	server := NewAPIServer("0.0.0.0:6000", nil)
	server.Serve()
}
