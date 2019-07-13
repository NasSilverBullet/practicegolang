package user

type User struct {
	FirstName string
	LastName  string
	Gender    string
	Birthday  string
}

func New(f, l, g, b string) *User {
	return &User{
		FirstName: f,
		LastName:  l,
		Gender:    g,
		Birthday:  b,
	}
}

func (u *User) Nickname() string {
	return u.FirstName[:1] + u.LastName[:1]
}
func (u *User) IsMale() bool {
	return u.Gender == "male"
}
func (u *User) Age() int {
	return age(u.Birthday)
}

func age(birthday string) int {
	return 24
}
