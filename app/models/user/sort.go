package user

import (
	"strings"
)

// Why is this not a thing....
// Return string slice for chaining.
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

type byEmailDomain []*User

func (s byEmailDomain) Len() int {
	return len(s)
}

func (s byEmailDomain) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byEmailDomain) Less(i, j int) bool {
	iEmail := strings.Join(reverse(strings.Split(s[i].Email, "@")), "")
	jEmail := strings.Join(reverse(strings.Split(s[j].Email, "@")), "")
	return strings.ToLower(iEmail) < strings.ToLower(jEmail)
}

type byCompany []*User

func (s byCompany) Len() int {
	return len(s)
}

func (s byCompany) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byCompany) Less(i, j int) bool {
	return strings.ToLower(s[i].Company) < strings.ToLower(s[j].Company)
}
