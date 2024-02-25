package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagocsz/go-gin-api-rest.git/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAlunos)
	r.POST("/alunos", controllers.StoreAluno)
	r.GET("/alunos/:id", controllers.GetAlunoById)
	r.PATCH("/alunos/:id", controllers.UpdateAluno)
	r.DELETE("/alunos/:id", controllers.DestroyAluno)

	r.Run()
}
