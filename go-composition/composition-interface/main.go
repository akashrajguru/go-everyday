package main

import "fmt"

// Define interfaces for behavior contracts
// Why it matters: By coding to interfaces rather than concrete types,
// we can swap implementations easily - real database in production, fake one in tests.
type Database interface {
	Save(data string) error
	Get(id string) (string, error)
}

type EmailSender interface {
	Send(to, subject, body string) error
}

// Real implementation of production code
// PostgresDB is a concrete type that implements the Database interface
// It has the Save and Get methods matching signature of the interface.
type PostgresDB struct {
	connectionString string
}

func (p PostgresDB) Save(data string) error {
	// Implement actual database save logic here
	fmt.Printf("Saving data to PostgreSQL: %s\n", data)
	return nil
}
func (p PostgresDB) Get(id string) (string, error) {
	return "data from PostgreSQL", nil
}
