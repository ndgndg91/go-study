package syntax

import "fmt"

// NewMap : initialize map with make()
func NewMap() map[int]string {
	var users map[int]string = make(map[int]string)
	AddToMap(users, 1, "남동길")
	AddToMap(users, 2, "김철수")
	AddToMap(users, 3, "김영희")
	fmt.Println(users)
	return users
}

// AddToMap : add key,value paire to map
func AddToMap(users map[int]string, key int, value string) {
	users[key] = value
}

// NewMapWithLiteral : initialize map with literal
func NewMapWithLiteral() map[int]string {
	users := map[int]string{1: "남동길", 2: "김철수"}
	return users
}

// PrintMapUsingRange : print map using range()
func PrintMapUsingRange(myMap map[int]string) {
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
