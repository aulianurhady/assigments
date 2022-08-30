package service

type User struct {
	Name string
}

type userService struct {
	db []*User
}

type UserInterface interface {
	Register(u *User) string
	GetUser() []*User
}

func NewUserService(db []*User) UserInterface {
	return &userService{
		db: db,
	}
}

func (u *userService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Name + " berhasil didaftarkan"
}

func (u *userService) GetUser() []*User {
	return u.db
}
