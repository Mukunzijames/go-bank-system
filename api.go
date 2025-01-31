package main

import (
	"fmt"
	"net/http"
)

type APIServer struct {
	listenAddr string
}

func newapiserver(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) handleAccount (w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Starting API server on", s.listenAddr)
	return nil
}
func (s *APIServer) handleGetAccount (w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Starting API server on", s.listenAddr)
	return nil
}
func (s *APIServer) handleCreateAccount (w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Starting API server on", s.listenAddr)
	return nil
}
func (s *APIServer) handleDeleteAccount (w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Starting API server on", s.listenAddr)
	return nil
}
func (s *APIServer) handleUpdateAccount (w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Starting API server on", s.listenAddr)
	return nil
}
func (s *APIServer) handleTransfer (w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Starting API server on", s.listenAddr)
	return nil
}
