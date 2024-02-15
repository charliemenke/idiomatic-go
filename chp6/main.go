package main

import (
	"fmt"
	"time"
)

type Person struct {
    FirstName string
    LastName string
    Age int
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
    return &Person{
        FirstName: firstName,
        LastName: lastName,
        Age: age,
    }
}
func MakePerson(firstName, lastName string, age int) Person {
    return Person{
        FirstName: firstName,
        LastName: lastName,
        Age: age,
    }
}
func (p *Person) SetAge(age int) {
	p.Age = age
}
func (p *Person) SetAgeAndReturn(age int) *Person {
	if p == nil {
		return &Person{"placeholder", "placeholder", age}
	}
	p.Age = age
	return p
}

func UpdateSlice(slice []string, str string) {
    slice[len(slice)-1] = str
    fmt.Println(slice)
}
func GrowSlice(slice []string, str string) {
    slice = append(slice, str)
    fmt.Println(slice)
}

func main() {
    // 1
    personValue := MakePerson("charlie", "menke", 26)
    personPointer := MakePerson("charlie", "menke", 26)
    fmt.Printf("%s %s\n", personValue.FirstName, personPointer.LastName) 
    
    //2
    slice := []string{"a", "b", "c"}
    fmt.Println(slice)
    UpdateSlice(slice, "d")
    fmt.Println(slice)
    GrowSlice(slice, "e")
    fmt.Println(slice)

    //3
    start := time.Now()
    personSlice := []Person{}
    //personSlice := make([]Person, 0, 10000000)
    for i := range 10000000 {
        personSlice = append(personSlice, Person{"charlie", "menke", i})
    }
    dur := time.Since(start)
    fmt.Printf("took: %f\n", dur.Seconds())

    var p *Person
    p = p.SetAgeAndReturn(25)
    fmt.Println(p.Age)
}
