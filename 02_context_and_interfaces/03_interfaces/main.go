// Interfaces
// one of the most important aspects of OOP is the usage of interfaces. As go doesn't have classes, neither inheritance,
// interface implementation is done implicitly by the usage of structs and receiver functions.

package main

import "fmt"

// declaring a Employee interface
type Employee interface {
	getName() string
}

// creating 2 types, Programmer and Engineer that should implement our interface
type Programmer struct {
	Name string
}

type Engineer struct {
	Name string
}

// the actual interface implementation comes here, with these receiver functions
func (p *Programmer) getName() string {
	return p.Name
}

func (e *Engineer) getName() string {
	return e.Name
}

func employeeDetails(e Employee) {
	fmt.Printf("Employee name: %s\n", e.getName())
}

func main() {

	programmer := Programmer{
		Name: "Vitor José Sacramento",
	}

	engineer := Engineer{
		Name: "MacGyver",
	}

	fmt.Println("Our Employees details:")
	employeeDetails(&programmer)
	employeeDetails(&engineer)
}
