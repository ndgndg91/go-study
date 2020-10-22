package syntax

// CanIDrink : return parameter age check 18 over ?
func CanIDrink(age int) bool {
	if koreanAge := age - 2; koreanAge < 18 { // If statement with initialization
		return false
	}

	return true
}
