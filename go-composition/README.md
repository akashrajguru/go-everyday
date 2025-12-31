# Composition in Go
Composition in Go is a design principle where you build complex types by combining simpler types, rather than using inheritance (which Go doesn't have). It's about "has-a" relationships rather than "is-a" relationships.

## Key Concepts
### Embedding: 
Go allows you to embed types within structs, giving the outer struct access to the embedded type's methods and fields.
### Interfaces: 
Define behavior contracts that types can implement, enabling polymorphism and flexible composition.