package user

type userRepository interface {
	fetch(userId string) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}
