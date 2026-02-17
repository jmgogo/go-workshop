package types

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) PrettyPrint() {
	println("User Information:")
	println("Name: ", u.Name)
	println("Email: ", u.Email)
	println("Age: ", u.Age, "\n")
}
