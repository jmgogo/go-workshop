package types

type Instructor struct {
	User
	Subject   string
	Classroom string
}

func NewInstructor(user User, subject, classroom string) Instructor {
	return Instructor{
		User:      user,
		Subject:   subject,
		Classroom: classroom,
	}
}
