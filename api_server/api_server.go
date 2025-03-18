package apiserver

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	auth_controllers "gitlab.com/hospital-web/hospital-backend/auth/infra/controller"
	auth_middlewares "gitlab.com/hospital-web/hospital-backend/auth/infra/middleware"
	consulta_controllers "gitlab.com/hospital-web/hospital-backend/consulta/infra/controller"
	doctor_controllers "gitlab.com/hospital-web/hospital-backend/doctor/infra/controller"
	leve_controllers "gitlab.com/hospital-web/hospital-backend/leves/infra/controller"
	paciente_controllers "gitlab.com/hospital-web/hospital-backend/paciente/infra/controller"
	user_controllers "gitlab.com/hospital-web/hospital-backend/user/infra/controller"
	user_middlewares "gitlab.com/hospital-web/hospital-backend/user/infra/middlewares"
)

var apiServer *APIServer

type APIServer struct {
	ginInstance *gin.Engine
}

func NewApiServer() *APIServer {

	ginInstance := gin.New()

	ginInstance.Use(gin.Recovery())

	ginInstance.Use(cors.Default())

	registerRoutes(ginInstance)

	apiServer = &APIServer{
		ginInstance: ginInstance,
	}
	return apiServer
}

func (s *APIServer) Start() {
	url := fmt.Sprintf(":%d", 3000)
	if err := s.ginInstance.Run(url); err != nil {
		log.Fatal("can't start api server on", url)
	}
}

func registerRoutes(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(200, "text/html; charset=utf-8", []byte("Welcome to Guanero's API"))
	})

	v1 := r.Group("/v1")

	auth := v1.Group("/auth")

	auth.POST("/signup", auth_controllers.Signup)
	auth.POST("/login", auth_controllers.Login)
	auth.POST("/refresh-token", auth_controllers.RefreshToken)

	user := v1.Group("/user").Use(auth_middlewares.VerifyLoginToken)
	user.GET("/rol", user_controllers.GetRole)
	user.PUT("/change-password", user_controllers.ChangePassword)
	user.GET("/info", user_controllers.GetMyInfo)

	admin := v1.Group("/admin").Use(auth_middlewares.VerifyLoginToken, user_middlewares.VerifyAdmin)
	admin.GET("/users", user_controllers.GetUsers)
	admin.POST("/create-user", user_controllers.CreateUser)
	admin.DELETE("/delete-user", user_controllers.DeleteUser)

	consulta := v1.Group("/consulta").Use(auth_middlewares.VerifyLoginToken)
	consulta.POST("/create-consulta", consulta_controllers.CreateConsulta)
	consulta.GET("/leve", consulta_controllers.GetLeve)
	consulta.GET("/garve", consulta_controllers.GetGrave)
	consulta.GET("/high-pressure-percentage", consulta_controllers.GetHighPressurePercentage)
	consulta.GET("/analysis-percentage", consulta_controllers.GetAnalysisPercentage)
	consulta.GET("/all", consulta_controllers.GetAllConsultas)
	consulta.GET("/", consulta_controllers.GetConsultaInfo)
	consulta.GET("/grade", consulta_controllers.GetConsultaGrade)

	leve := v1.Group("/leve").Use(auth_middlewares.VerifyLoginToken)
	leve.GET("/oldest-leve-patient", leve_controllers.GetOldestLevePatientWithDiagnosis)

	paciente := v1.Group("/paciente").Use(auth_middlewares.VerifyLoginToken)
	paciente.POST("/", paciente_controllers.CreatePaciente)
	paciente.GET("/all", paciente_controllers.GetAllPacientes)

	doctors := v1.Group("/doctor").Use(auth_middlewares.VerifyLoginToken)
	doctors.POST("/", doctor_controllers.CreateDoctor)
	doctors.GET("/all", doctor_controllers.GetAllDoctors)

}
