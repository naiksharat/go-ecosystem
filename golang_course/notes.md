Package
-------

* Multiple go files for a similar purpose
* 2 types
    * executable package 
        * starts with - package main 
        * on build spits out - main
        * Requirement: must have a main function
    * reusable / helper / dependency package 
        * starts with package name
        * doesn't spit out anything
        * Example: Standard libs - fmt, math, debug etc


Variables and data types:
-------------------------
* var card string = "Ace of Spades"
* types:
    * bool, int, float64, string


```
package main

import "fmt"

// var deckSize int = 52 	// valid		
// var deckSize = 52		// valid	
// deckSize := 52  			// invalid in global scope
// const decksize = 52 		// valid



func main() {
	//var card string = "Ace of Spades" // full declaration and value assignment
	card := "Ace of Spades"             // infer type and enforce.
	//card = "Ace of Diamonds"			// change value after declaration
	// card = 53						// error - type mismatch
	fmt.Println(card)

	//var deckSize int					// declare
	//deckSize = 53						// assign
	//fmt.Println(deckSize) 			//local > global
    //i, careds := 1, 2                 // multiple

	// var deckSize = 52		// valid	


}

```


array and slices:
-----------------

* array - fixed length array (list of records of the same data type)
* slice - variable length array (list of records of the same data type)

```
package main

import "fmt"

func main() {
    //var cards []string = []string{"",""}
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Clubs")       // not an in place append. a new string is returned

	for i, card := range cards {                // slice iteration
		fmt.Println(i, card)
	}

    for i := range cards {
        fmt.Println(i)                          // Valid but prints only indices, not card values
    }

    for _, card := range cards {
        fmt.Println(card)                       // ignore indices and print cards one by one
    }

	_ = NoCliArgError{}.Error()					// if one arg that's ignored

}

func newCard() string {
	return "Five of Diamonds"
}

```

access and length: 
a[0]
a[len(d) - 1]

slicing:
a[:5] 
a[5:]

packages/libs:
----------------
```
state.go
--------
package main

import "fmt"

func printState() {
	fmt.Println("berlin")
}

/*
// go run main.go state.go
*/

package main

func main() {
	printState()
}

```

strings:
strings.Split()
strings.Join()

os:
os.Exit()
os.Remove()
os.Create()

io/util:
ioutil.ReadFile()
ioutil.WriteFile()

math/rand:
rand.Intn()

time:
time.UnixNano()

Types and receiver functions: OO in Go
----------------------------------------

```
type deck []string

func (d deck) print() {
	for card := range d {
		fmt.Println(card)
	}
}

note:
d = this/self

usage: 
d = deck{}
d.print()

typecasting:
deck to string : string(d)
string to deck : deck([]string{})
```

Return values:
--------------

* return values 
```
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
```

* read returned values

```
a, b = deal()
```

* swap

a, b = b, a


if else, nil and error handling:
----------------------
```
nil ==> None

if err != nil {
    handle error/ exit
}


```

switch case - select case
----------------------------
	* https://yourbasic.org/golang/switch-statement/ -- fallthrough - continue to next case
	* Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.
```
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```


structs & receiver functions with structs
------------------------------------------
```
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

	jim.updateFirstName("Jimmyyyyyyy") // pass by value
	jim.print()
	// {Jim Pa {paji@hhh.com 123456}}
	// No change
	// main(has Jim Pa) --(passes Jim Pa and calls)--> updateFirstName(has a copy of Jim Pa) --(updates Jim Pa to Jimmy Pa)--> returns to main

	//jimPointer := &jim
	fmt.Println(&jim)
}

func (p person) print() {
	fmt.Println(p)
}

func (p person) printFull() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
	p.print()
}

```

pointers and gotchas:
----------------------

```
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
    // fmt.Println(jim.contact)	          // valid
	// fmt.Println((*jimPointer).contact) //valid
	// fmt.Println(jimPointer.contact) 	// valid
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
}

func (p person) updateFirstName(newFirstName string) { // always pass 'person' as a value while calling this function (doesn't matter if the caller is a pointer or a value)
	p.firstName = newFirstName
	//p.print()
}

func (p *person) updateLastName(newLastName string) { // always pass 'person' as a reference (reference/pointer to person) while calling this function (doesn't matter if the caller is a pointer or a value)
	p.lastName = newLastName
}

```

Notes:

* convert value to ref :  ref := &value
* convert ref to value :  value = *ref
* (* in function signature) i.e receiver type/arguments : Indicates the type of data 
    * If the function is expecting a pointer - 
        `func (p *person) upLastName(lastNameHeapPointer *lastName) { val = *lastName}`
    * function call : 
        `(&per).upLastName(&lastName)`


Pointer gotchas: value types and reference types
-------------------------------------------------

* Go is a pass by value language
* Slice = [capacity, length, ptr to array head that contains data (-> [array 1, 2, 3, 4])]
* slice contains pointer to an array that stores data ( slice does not the values i.e 1,2,3,4 in the above example ). So when slice is copied and passed over to the next function (pass by value i.e capacity, length and pointer to array),pointer is passed over as well (which points to the same set of data i.e array in the caller and called function)
* No need of pointers to pass the reference here, because pointer is already passed by nature

Example:
```
	a := []int{1, 2, 3}
	modSlice(a)
	fmt.Println(a) // 10 2 3 not 1,2,3

}

func modSlice(a []int) {
	a[0] = 10
}

----
func main() {
	b := best{"hi", "there"}
	b.test()
	fmt.Println(b) // Hi here
}

type best []string

func (b best) test() {
	b[1] = "here"
}



````

Value types :  
* pointers are needed to pass the reference)
* int, float, string, bool, struct"

Reference types : 
* a special structure which holds some metadata and a pointer to data. So when these values are copied over the comprising pointer is copied over as well)
* pointer, slice, map, channels



Maps:
-----
```
func main() {
	// declaration
	// 1
	m := map[string]int{
		"blue":  1,
		"green": 2,
	}

	// 2
	n := make(map[int]int)
	n[10] = 1
	fmt.Println(n)

	//update
	m["red"] = 1
	fmt.Println(m)

	// iteration
	for color, number := range m {
		fmt.Println(color, number)
	}
    
    // delete
	delete(m, "green")
	fmt.Println(m)

}
```

| maps | struct |
| ---- | ------ |
| all keys of same type, all values of same type | no restriction |
| ref type | value type |
| access = v[key], indexed | access = v.key, not indexed |
| elems can be added after creation | fixed structure |
| collection of a thing | a thing's properties |


tests:
------

``` deck_test.go
go test 
package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
```

interfaces:
------------
Syntax:
```
type nameInterface interface {
	interfaceFunctionThatWillBeImoplementedByOthers(int) string
	anotherFn() int
}

type implementer struct{}

func (i implementer) interfaceFunctionThatWillBeImoplementedByOthers(a int) string {
	//def
}

func (i implementer) anotherFn() int {
	// def
}


// interface consisting of multiple interfaces
type compoundInterface interface {
	nameInterface
	anotherSuchInterface
}

```

Notes:
* Interfaces are implicit: 
	* Members/implementers are not specifically assigned to an interface. 
	* As soon as an implementer implements the interface method, the implementer becomes an honorary member of the interface
* Helps in dependency inversion ( in a statically typed environment this is important to pass different types of args to a function. Dynamically typed languages don't have this probelm)
	```
	func function(intf interface_type) {
		intf.interface_method()
	}

	// look at the arg below. Instead of interface_type (as per function implementation I can pass in implementer1/2/3
	function(implementer_type_1) // implementer type 1 -> implements intf.interface_method()
	function(implementer_type_2) // implementer type 2 -> implements intf.interface_method()
	function(implementer_type_3) // implementer type 3 -> implements intf.interface_method()
	```
* Dependency inversion aids - Open Closed principle (open for extension, closed for modification) and aids - code reuse
* Match all interface:
	* declaration: type matchAll interface{}
	* Usage: func acceptAny(any matchAll) {}
	* Usage: func acceptAny(any interface{}) {}
	* any could be - int, string, strict chan etc

``` Example 1:
package main

import "fmt"

type triangle struct {
	height float64
	width  float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.width * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {

	t := triangle{1, 2}
	s := square{1}

	printArea(t)
	printArea(s)
}
```


Example 2:
Custom writers, custom exceptions
os.exit
os.open
os.args


io.copy(dest, src)
``` 
package main

import (
	"fmt"
	"io"
	"os"
)

type NoCliArgError struct{}

func (ce NoCliArgError) Error() string {
	fmt.Println("Len < 2. No CLI Arg")
	return ""
}

type FileDoesntExist struct{}

func (ce FileDoesntExist) Error() string {
	fmt.Println("File doesn't exist")
	return ""
}

type customWriter struct{}

func (cw customWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs), "from this")
	return len(bs), nil
}

func main() {
	if len(os.Args) < 2 {
		_ = NoCliArgError{}.Error()
		os.Exit(1)
	}

	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		_ = FileDoesntExist{}.Error()
		os.Exit(2)
	}

	io.Copy(customWriter{}, f)

	f.Seek(0, 0)
	io.Copy(customWriter{}, f)

	f.Seek(0, 0)
	io.Copy(os.Stdout, f)

}

```

Example 3:
```
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	const url = "http://google.com"
	resp, _ := http.Get(url)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	//Reader:
	//resp.Body.Read is a reader function
	// it has some info that it's keeping to itself
	// we can pass a buffer ([]byte) to the readeer fucntion and ask it to fill the buffer with info
	bs := make([]byte, 9)
	resp.Body.Read(bs)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }
	fmt.Println(string(bs))

	// read from resp.body -- READER
	// write to os.stdout -- WRITER
	io.Copy(os.Stdout, resp.Body)

	// Custom writer/custom reader // implement write/read function
}

```



channels & routines:
---------------------
go 
chan


===================================

package, package exports 
----------------------------------
* SCOPE:
All files in a package share the shame scope: Vars, functions etc w/o explicitly mentioning it
Example:
	* pack
		* file1.go -> package p
		* file2.go -> package p
	* file1.go can access all the vars of file2.go since they are in the same package


* Capitalized --> type, variable, func, struct and even struct fields (type S struct{ Name string}) etc --> are exported
i.e can use when a pacakeg is imported
* Otherwise not usable

Notes:

* https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go
* Package visibiity and patterns:
https://www.digitalocean.com/community/tutorials/understanding-package-visibility-in-go
* import ( "directory name of the package") not ("package name in files") or not (file names in packages)

modules:
---------
https://www.digitalocean.com/community/tutorials/how-to-use-go-modules

* modules
	* Inside module dir: go mod init module_dir
		* module_dir
				* go.mod
				* package_dir
					* subpacjage_dir
					* file
						* package package_dir
						  import "fmt"
				* another_package_dir
					* file
						* package another_package_dir
						  import "module_dir(from go mod init)/package_dir"

	* Inside module dir: go get cobra
		* the cobra lib is pulled and stored in GOPATH/src/github.com
		* go.mod is updated with the dependency
		* go.sum is created with hashes (hash(go.mod) = go.sum). Tamper proofing
		* Can be imported as "import "github.com/spf13/cobra"" -> locally stored though in GOPATH
		* go mod tidy : in case you are using 3rd party modules but forgot to do a go get.

Notes:
* GOPATH = go env GOPATH
	* This is where go looks for 3rd party packages and mods
	* pkg, bin, src
	* example: import "github.com/test/asa"
* GOROOT = go env GOROOT
	* This is where go stores and looks for standard packages
	* example: import "fmt" import "greet" (if you create a package greet and move it to GOROOT)
* go.mod: this is where the go tool will look at, to resolve packages of the module. For this to work the import statement should start with 'modulename'
 

To-do:
------

* anonymous functions
* chan and go routines
* defer panic recover
* https://www.digitalocean.com/community/tutorials/understanding-package-visibility-in-go
* type assertions: https://yourbasic.org/golang/type-assertion-switch/
* pointer receiver: https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
* reflections: https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
* tour of go


done:
-----
* import a user written package/file/method/function
* multiple go file package
* scopes local global
* goPATH
https://www.digitalocean.com/community/tutorials/how-to-use-go-modules