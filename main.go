package main

import (
	"GoCrud/models"
	"fmt"
	//"github.com/gin-gonic/gin"
)

func main() {
	//	router := gin.Default()

	models.ConnectDatabase() // new!

	// ...

	//router.Run("localhost:8080")

	conn, err := models.DB.DB()
	if err != nil {
		fmt.Print("error")
	} else {
		defer conn.Close()
	}

	// Creating the first person
	// person3 := models.Person{
	// 	PersonID:  3,
	// 	LastName:  "Kumar",
	// 	FirstName: "kailash",
	// 	Address:   "Pali",
	// 	City:      "Jaipur",
	// }

	// Creating the second person
	// person4 := models.Person{
	// 	PersonID:  4,
	// 	LastName:  "padia",
	// 	FirstName: "Ronak",
	// 	Address:   "kandivali",
	// 	City:      "Mumbai",
	// }

	EmployeeTable1 := models.EmployeeTable{
		EmployeeID: 3735637,
		Name:       "kailash Kumar",
		City:       "Jaipur",
		Gender:     "Male",
	}

	EmployeeTable2 := models.EmployeeTable{
		EmployeeID: 8928468,
		Name:       "Sahitya Singh",
		City:       "Lucknow",
		Gender:     "Male",
	}

	EmployeeTable3 := models.EmployeeTable{
		EmployeeID: 2134242,
		Name:       "Abhay Kumar",
		City:       "Jhasi",
		Gender:     "Male",
	}

	EmployeeTable4 := models.EmployeeTable{
		EmployeeID: 9872522,
		Name:       "Jaya",
		City:       "Thane",
		Gender:     "Female",
	}
	// Inserting the persons into the database
	// models.DB.Create(&person3)
	// models.DB.Create(&person4)
	// models.DB.Create(&EmployeeTable1)
	// models.DB.Create(&EmployeeTable2)

	// models.DB.Create(&EmployeeTable3)
	// models.DB.Create(&EmployeeTable4)

	fmt.Println(EmployeeTable1.GetEmployeeDetails())
	fmt.Println(EmployeeTable2.GetEmployeeDetails())
	fmt.Println(EmployeeTable3.GetEmployeeDetails())
	fmt.Println(EmployeeTable4.GetEmployeeDetails())

	//Update the persons into the database
	// models.DB.Model(&models.Person{}).Where("LastName= ?", "padia").Update("LastName", "Office")
	// models.DB.Model(&models.Person{}).Where("LastName= ?", "Singh").Update("LastName", "Rajput")

	//Delete the persons into the database
	// models.DB.Where("LastName = ?", "Kumar").Delete(&models.Person{})
	//models.DB.Delete(&person2.City, "Mayur Vihar")
	// models.DB.Where("FirstName = ?", "Vijay").Delete(&person1)

}
