package contacts

import "fmt"

func ExampleContact_String() {
	c1 := Contact{"Max", "Mustermann", "+49 123 45678", []string{}}

	fmt.Println(c1)

	// Output:
	// Max Mustermann
	//   Telefon: +49 123 45678
	//   Tags: []
}
