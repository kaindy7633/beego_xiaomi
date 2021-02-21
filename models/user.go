package models

// User is ...
type User struct {
	ID       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// TableName is ....
func (User) TableName() string {
	return "user"
}
