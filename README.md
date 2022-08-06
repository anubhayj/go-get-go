# go-get-go

# What is go ?
Go is a statically typed, compiled programming language.

   How to use Go ? 
   Don't bring any baggage , this language misses out a lot of stuffs, but compensates on it in its own way.
   
   Criticisms on Go
   
   
     --> try,catch is missing.
     --> No semicolons
     --> Lack of Function Overloading and Default Values for Arguments.
     --> Generics

# Installations.
Homepage

golang.org
https://go.dev/dl/


# Go Code organization
Go programs are organized into packages.
A repository contains one or more modules. 

# Some common components.
goPath:

# variables:
  Variables are statically-typed, ,there are two types in which we can define a variable.
  ```go
   var num int // Explicitly typed
   var num := 10 // Implicitly typed,the compiler will assign the variable type to match the type of the initializer.
   const num = 20  // Defining constants
  ``` 
## packages : 
  Go applications are organized in packages. A package is a collection of source files located in the same directory. All source files in a directory must share the same package name. When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter can be used / accessed. The recommended style of naming in Go is that identifiers will be named using **camelCase**, except for those meant to be accessible across packages which should be **PascalCase**.
  
## functions :
```go
package helloGo

//public functions
func Hello (name string) string {
    return hi(name)
}

//private functions
func hi (name string) string {
    return "hi " + name
}
```
**Notice the function names.**

## interfaces :

Marshalling and UnMarshalling :
structs :
