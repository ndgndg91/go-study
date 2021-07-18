package syntax

// WhatIsGrade : return Grade
func WhatIsGrade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case 80 <= score && score < 90:
		return "B"
	case 70 <= score && score < 80:
		return "C"
	case 60 <= score && score < 70:
		return "D"
	default:
		return "F"
	}
}
