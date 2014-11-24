package main

import (
	"fmt"
	"strconv"
)

type Say func(string) string

type Person struct {
	Name string
	Talk Say
}

func (p *Person) Introduce() {
	p.Talk(fmt.Sprintf("Hi, I'm %s", p.Name))
}

type Student struct {
	*Person
	Gpa float64
}

func NewStudent(name string, gpa float64) *Student {
	return &Student{
		Person: &Person{"Bob", func(words string) string {
			fmt.Println(words)
			return words
		}},
		Gpa: gpa,
	}
}

func main() {
	user := NewStudent("Bob", 3.14)
	fmt.Printf("%s's GPA is %f.\n", user.Name, user.Gpa)
	user.Introduce()
	fmt.Printf("Bob said: '%s'\n", user.Talk("My GPA is "+strconv.FormatFloat(user.Gpa, 'f', -1, 32)))

	syncher := synchronizer{}
	syncher.Increment()
}
