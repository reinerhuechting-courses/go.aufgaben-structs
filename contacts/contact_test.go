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

func ExampleContact_HasTag() {
	c1 := Contact{"Max", "Mustermann", "+49 123 45678", []string{"Freunde", "Kollegen"}}

	fmt.Println(c1.HasTag("Freunde"))
	fmt.Println(c1.HasTag("Familie"))

	// Output:
	// true
	// false
}

func ExampleContact_AddTag() {
	c1 := Contact{"Max", "Mustermann", "+49 123 45678", []string{"Freunde", "Kollegen"}}

	c1.AddTag("Wichtig")
	c1.AddTag("Wichtig")

	fmt.Println(c1)

	// Output:
	// Max Mustermann
	//   Telefon: +49 123 45678
	//   Tags: [Freunde Kollegen Wichtig]
}
