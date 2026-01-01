package main

// Define interfaces for behavior contracts
type Database interface {
	Save(data string) error
	Get(id string) (string, error)
}
