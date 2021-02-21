package models

// Manager ...
type Manager struct {
	ID       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleID   int
	AddTime  int
	isSuper  int
}

// TableName ...
func (Manager) TableName() string {
	return "manager"
}
