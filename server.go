package main

import (
	"fmt"
	"net"
)

// Client represents a chat user
type Client struct {
	conn     net.Conn
	username string
}

// Server holds all active clients
type Server struct {
	clients map[string]*Client
}

// NewServer creates a new chat server
func NewServer() *Server {
	return &Server{
		clients: make(map[string]*Client),
	}
}

// Start runs the server
func (s *Server) Start(port string) error {
	// TODO: Listen for connections on this port
	// TODO: Accept new clients
	// TODO: Handle messages from clients
	return nil
}

// handleClient manages one client's connection
func (s *Server) handleClient(conn net.Conn) {
	// TODO: Get username
	// TODO: Handle messages
	// TODO: Clean up when client quits
}

func main() {
	server := NewServer()
	fmt.Println("Starting server on port 8080...")
	server.Start("8080")
}

