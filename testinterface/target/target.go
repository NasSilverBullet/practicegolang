package target

type UserInterface interface {
	Nickname() string
	IsMale() bool
	Age() int
}

func IsTestTarget(u UserInterface) bool {
	return u.Age() > 25
}
