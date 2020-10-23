package syntax

// struct has no constructor
type user struct {
	name      string
	age       int
	interests []string
}

// CreateUsers : create user and return
func CreateUsers(name string, age int, interests []string) interface{} {
	newUser := user{name: name, age: age, interests: interests}
	return newUser
}
