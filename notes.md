- error is last var returned in function
- defer functions can modify or take action on named returned values, usually with a closure
- try to allocate slice capacity if you have an idea of how many elements it will contain
- `var person Person && person := Person{}` is the same thing unlike slices/maps. both assign a empty literal and create a struct with "zeroed" values for the key:values
- Golang is "Call By Value" meaning paramenters to a function are always a copy of the input value (expect for slices and maps)
- golang switch breaks at end of case automatically, no fallthrough unless specified
- code your struct methods to expect nil instances
	- with a pointer receiver method, you cannot have it both handle nil and make the original pointer non-nil
```golang
type Person struct {
	Age int
	Name string
}
func (p *Person) SetAge(age int) {
	p.Age = age // this will panic as pointer `p` is still nil
}
func (p *Person) SetAgeAndReturn(age int) *Person {
	if p == nil {
		return Person{age, "placeholder"}
	}
	p.Age = age
	return p
}

var p *Person
p.SetAge(26)
fmt.Println(p.Age) 
```
- a pointer instance's method set includes both pointer receiver and value receiver methods
- be careful using comparison with interfaces as the underlying type of the interface may not be comparable
- function types can have methods that can also implement interfaces
- use function types to bridge between interfaces
	- see `net/http`'s `HandlerFunc` type
- ACCEPT interfaces, return concrete types (structs)
- interfaces encourage the decorator pattern. [[link1https://www.henrydu.com/2022/01/05/golang-decorator-pattern/]]
	- either function that takes in interface and returns another type that implements same interface, ex: `gzip.NewReader` takes in `io.Reader` and returns a concrete type of `Reader` that implements `io.Reader`
	- or 
