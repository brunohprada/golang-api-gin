package main

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"Nome": "Bruno",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos)
	r.Run()
}
