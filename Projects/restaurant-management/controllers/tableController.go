package controllers

import (
	"github.com/atharvamahamuni/golang/projects/restaurant-management/database"
	"github.com/gin-gonic/gin"
)

var tableCollection = database.DBinstance("table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}