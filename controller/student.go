package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xukai-echo/model"
	"github.com/xukai-echo/storage"
)

func GetStudents(c echo.Context) error {
	students, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Students, error) {
	db := storage.GetDBInstance()

	// students := []model.Students{
	// 	{
	// 		Id:   "1",
	// 		Name: "2",
	// 	},
	// }
	students := []model.Students{}
	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}
