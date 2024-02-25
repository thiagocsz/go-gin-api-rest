package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagocsz/go-gin-api-rest.git/database"
	"github.com/thiagocsz/go-gin-api-rest.git/models"
)

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func GetAlunoById(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Aluno nao encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func StoreAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func UpdateAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Aluno nao encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func DestroyAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Aluno nao encontrado"})
		return
	}
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, aluno)
}
