package model

type Community struct {
	// Name is the name of the community
	Name string `json:"name"`
}

// Communities slice of communities
type Communities []*Community

// Person struct for a person
type Person struct {
	// Name is the name of a person
	Name string `json:"name"`
	// Age is the age of the person
	Age uint8 `json:"age"`
	// communities user belongs to some communities
	Communities Communities `json:"communities"`
}

// People slice of a lot of person
type People []*Person
