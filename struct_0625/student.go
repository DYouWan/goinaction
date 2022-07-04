package main

type student struct {
	id int
	age int
	name string
}

func NewStudent(id int,age int, name string) *student {
	s := &student{
		id:   id,
		age:  age,
		name: name,
	}
	return s
}

func (s student) SetName(name string) {
	s.name = name
}
func (s *student) SetAge(age int)  {
	s.age = age
}