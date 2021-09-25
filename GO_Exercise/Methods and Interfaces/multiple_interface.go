// Go program to illustrate the
// concept of multiple interfaces
package main

import "fmt"

// Interface 1
type AD interface {
	details()
}

// Interface 2
type AA interface {
	articles()
}

// Structure
type Author_details struct {
	a_name  string
	branch  string
	college string
	year    int
	salary  int
	revenue int16
}

// Implementing method
// of the interface 1
func (a Author_details) details() {

	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nRevenue: %d", a.revenue)

}

// Implementing method
// of the interface 2
func (a Author_details) articles() {

	pendingarticles := a.revenue - int16(a.salary)
	fmt.Printf("\nTotal profit: %d", pendingarticles)
}

// Main value
func main() {

	// Assigning values
	// to the structure
	values := Author_details{
		a_name:  "kloud1",
		branch:  "CS",
		college: "MIT",
		year:    2022,
		salary:  50000,
		revenue: 10000,
	}

	// Accessing the method
	// of the interface 1
	var i1 AD = values
	i1.details()

	// Accessing the method
	// of the interface 2
	var i2 AA = values
	i2.articles()

}
