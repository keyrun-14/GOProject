package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	// RefId    
	Name             string             `json:"name,omitempty" validate:"required"`
	Model            string             `json:"model,omitempty" validate:"required"`
	Specification    string             `json:"specification,omitempty" validate:"required"`
	Price            int                `json:"price,string,omitempty" validate:"required"`
}

type Register_Users struct{
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Email       string             `json:"email,omitempty" validate:"required"`
	Password    string             `json:"password,omitempty" validate:"required"`
}


type Login_Users struct{

	Email       string             `json:"email,omitempty" validate:"required"`
	Password    string             `json:"password,omitempty" validate:"required"`
}