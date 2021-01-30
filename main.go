package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/role", Get)

	router.GET("/role/:id", GetOne)

	router.POST("/role", Post)

	router.PUT("/role/:id", Put)

	router.DELETE("/role/:id", Delete)

	router.Run(":8080")
}

// 取得全部資料
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

// 取得單一筆資料
func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var targetRole *Role
	for _, value := range Data {
		if value.ID == uint(id) {
			targetRole = &value
			break
		}
	}

	if targetRole == nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "role not found"})
		return
	}

	c.JSON(http.StatusOK, targetRole)
}

// 新增資料
func Post(c *gin.Context) {
	var newRole Role
	if err := c.ShouldBind(&newRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "123"})
		return
	}

	var maxId uint = 0
	for _, role := range Data {
		if role.ID > maxId {
			maxId = role.ID
		}
	}

	newRole.ID = maxId + 1
	Data = append(Data, newRole)
	c.JSON(http.StatusOK, newRole)
}

// 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var existingRole Role
	if err := c.ShouldBind(&existingRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i := 0; i < len(Data); i++ {
		if Data[i].ID == uint(id) {
			Data[i].Name = existingRole.Name
			Data[i].Summary = existingRole.Summary
			Data[i].Skills = existingRole.Skills
			break
		}
	}
	c.Status(http.StatusNoContent)
}

// 刪除資料
func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i, role := range Data {
		if role.ID == uint(id) {
			Data = append(Data[0:i], Data[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
