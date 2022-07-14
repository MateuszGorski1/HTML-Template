package person

type Person struct {
	Name    string
	Surname string
	Age     int
}

func CreatePerson(name string, surname string, age int) Person {
	Person := Person{
		Name:    name,
		Surname: surname,
		Age:     age,
	}
	return Person
}
