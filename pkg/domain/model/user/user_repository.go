package user

type userRepository interface {
	find(userId string) (err error)
	findAll() (err error)
	save() (err error)
	update() (err error)
	delete() (err error)
}
