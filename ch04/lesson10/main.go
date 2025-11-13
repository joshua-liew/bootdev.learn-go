package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	user := User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
			MessageCharLimit: 100,
		},
	}
	if membershipType == "premium" {
		user.MessageCharLimit = 1000
	}
	return user
}
