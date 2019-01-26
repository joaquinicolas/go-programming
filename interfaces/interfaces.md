# Interfaces

Interfaces express generalization or abstraction of other's types behaviors.
In Go, interfaces are satisfied implicitly. This is useful for types that are defined in packages
that you don't control.

### As contracts
Interfaces might define a contract between a function and its callers. On the one hand, 
the contract requires that the caller provide a value of a concrete type with the right signature
and behavior. 
On the other hand, the interface guarentee that the function will do its job with a value 
that satisfies the interface. 

By convention, one method interfaces are named with the suffix -er. 


### Interface satisfaction

A type satisfies an interface when it possesses all the methods that
the interface requires. 
*os.File satisfies io.Reader, io.Closer, io.Writer, and ReadWriter.

"A type is a particular interface type"

Like an envelope that wraps and conceals a letter, only methods revealed
by the interface may be called, even if the concrete type has others.

```go
os.Stdout.Write([]byte("hello"))
var w io.Writer
w = os.Stdout
w.Write([]byte("hello"))
w.Close() // compile error: io.Writer lacks Close method
```