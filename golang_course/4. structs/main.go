package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo //embedded struct
	//contactInfp	// = contactInfo contactInfo
}

func main() {
	/*
		// type person struct {
		// 	firstName string
		// 	lastName  string
		// }
		var martin person
		// zero value assignment - nil isn't assigned
		// string 			 ""
		// int and float	 0
		// bool				 false
		// Example:
		// var martin person  -> martin{firstname: "", lastname: ""}
		fmt.Println(martin)         // o/p: {}
		fmt.Printf("%+v\n", martin) // o/p: {firstName: lastName:}
		// printf +v => prints keys and values

		// update struct's values
		martin.firstName = "Martin"
		martin.lastName = "Mason"
		fmt.Println(martin)
		fmt.Printf("%+v\n", martin)

		// other declaration methods
		alexa := person{"Alexa", "Anderson"} //order
		fmt.Println(alexa)
		alex := person{
			firstName: "Alex",
			lastName:  "Anderson",			// ',' is required at the end
		}

		fmt.Println(alex)
	*/

	jim := person{
		firstName: "Jim",
		lastName:  "Pa",
		contact: contactInfo{
			email:   "paji@hhh.com",
			zipcode: 123456,
		},
	}

	jimPointer := &jim

	// both jim and jimPointer can call and access all receiver functions and vars
	// fmt.Println((*jimPointer).contact) //valid
	// fmt.Println(jimPointer.contact)	// valid
	// jim.print()						// valid
	// jimPointer.print()				// valid

	jimPointer.updateFirstName("Jimmyyyyyy")
	fmt.Println(jim) //{Jim Pa {paji@hhh.com 123456}} // 'jim' value is copied & passed to updateFirstName (since the fn signature has (p person)). So all changes done in that function are local to it

	//(&jim).updateLastName("Paji") OR
	jimPointer.updateLastName("Paji")
	fmt.Println(jim) //{Jim Paji {paji@hhh.com 123456}} //  Reference to 'jim' is copied & passed to updateFirstName (since the fn signature has (p *person)). So all changes done in that function will reflect in the parent caller function

	// Pointer shortcut  - Go inference
	jim.updateLastName("Pajira") // Instead of using the pointer, just use the value and call the function. Go inferes and passes reference of 'jim/person' to the function
	fmt.Println(jim)

	a := []int{1, 2, 3}
	modSlice(a)
	fmt.Println(a) // 10 2 3 // slice contains pointer to array that stores data ( not the data itself ). So when slice is copied and passed over to the next function, pointer is passed over as well (which points to the same set of data i.e array in the caller and called function)

	b := best{"hi", "there"}
	b.test()
	fmt.Println(b)
}

type best []string

func (b best) test() {
	b[1] = "here"
}

func modSlice(a []int) {
	a[0] = 10
}

func (p person) print() {
	fmt.Println(p)
}

func (p person) printFull() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstName(newFirstName string) { // always pass 'person' as a value while calling this function (doesn't matter if the caller is a pointer or a value)
	p.firstName = newFirstName
	//p.print()
}

func (p *person) updateLastName(newLastName string) {
	p.lastName = newLastName
}
