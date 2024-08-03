package models

import "fmt"

type EmployeeTable struct {
	EmployeeID int    `json:"id"  gorm:"column:EmployeeID"`
	Name       string `json:"last_name" gorm:"column:Name"`
	// FirstName string `json:"first_name" gorm:"column:FirstName"`
	// Address   string `json:"address" gorm:"column:Address"`
	City   string `json:"city" gorm:"column:City"`
	Gender string `json:"gender" gorm:"column:Gender"`
}

func (v *EmployeeTable) TableName() string {
	return "EmployeeTable"
}

func (v *EmployeeTable) GetEmployeeDetails() string {
	det := fmt.Sprintf("Employee Name is %v ,  gender is %v and  lives in %v", v.Name, v.Gender, v.City)
	return det
}
