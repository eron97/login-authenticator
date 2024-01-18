package models

type CreateUser struct {
	First_Name string `json:"first_name" binding:"required,min=4,max=100" example:""`
	Last_Name  string `json:"last_name" binding:"required,min=4,max=100" example:""`
	Email      string `json:"email" binding:"required,email" example:"test@test.com"`
	Password   string `json:"password" binding:"required,min=6,containsany=!@#$%*" example:"password#@#@!2121"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUsers struct {
	First_name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
}
