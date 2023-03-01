package main

func main() {
    server := New("localhost:3000", NewRepositories())
    server.Listen()
}
