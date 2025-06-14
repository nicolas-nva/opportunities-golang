package schemas

import (
	"gorm.io/gorm"
)

type Opening struct{
	gorm.Model
	Role 		string
	Company 	string
	Location 	string
	TypeOfWork 	string
	Link 		string
	Salary 		int64
}