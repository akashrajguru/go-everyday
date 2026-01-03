package main

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
