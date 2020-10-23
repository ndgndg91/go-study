package syntax

import "fmt"

// GetNames : return 5 length string array contains 5 names
func GetNames() [5]string {
	names := [5]string{"남동길", "김철수", "김영희"}
	names[3] = "김진수"
	names[4] = "박진희"
	fmt.Println(names)
	return names
}

// SliceTest : without length
func SliceTest() []string {
	array := []string{}
	array = append(array, "남동길")
	return array
}
