package models

type UserModel struct {
	BaseModel
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (user *UserModel) TableName() string {
	return "users"
}
