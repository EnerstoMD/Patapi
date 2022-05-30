package handler

import (
	"lupus/patapi/pkg/auth"
	"lupus/patapi/pkg/middleware"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func InitRoutes(router *gin.Engine, ph PatientHandler, ch CalendarHandler, uh UserHandler, a auth.AuthService) {
	SetupMiddleware(router)
	router.GET("", welcome)
	router.NoRoute(notFound)
	v1 := router.Group("v1")
	{
		user := v1.Group("user")
		{
			user.POST("login", uh.Login)
			user.DELETE("logout", uh.Logout)
		}

		admin := v1.Group("admin")
		admin.Use(middleware.BearerAuth(a))
		{
			admin.POST("register", uh.Register)
			admin.GET("users", uh.GetUsers)
			admin.DELETE("user/:id", uh.DeleteUser)
			admin.PATCH("user/:id/pwd", uh.AdminUpdatePassword)
		}

		userinfo := user.Group("userinfo")
		userinfo.Use(middleware.BearerAuth(a))
		{
			userinfo.GET("", uh.GetUserInfo)
			userinfo.PATCH("pwd", uh.UpdatePassword)
			userinfo.PATCH("", uh.UpdateUserInfo)
			// userinfo.DELETE("/:id", uh.DeleteUserInfo)
		}

		patient := v1.Group("patient")
		patient.Use(middleware.BearerAuth(a))
		{
			patient.GET("", ph.GetAllPatients)
			patient.POST("", ph.CreatePatient)
			patient.GET("search", ph.SearchPatientByName)
			patient.GET("search/count", ph.CountSearchPatientByName)
			patient.GET(":id", ph.GetPatientById)
			patient.PATCH(":id", ph.UpdatePatient)
			patient.GET("ins", ph.SearchPatientByINSMatricule)
			patient.GET("card", ph.ReadCarteVitale)
			patient.POST("csvbatchload", ph.BatchLoadPatients)

			patient.POST(":id/comment", ph.CreatePatientComment)
			patient.GET(":id/comment", ph.GetPatientComments)
			patient.DELETE(":id/comment/:commentid", ph.DeletePatientComment)

			patient.POST(":id/disease", ph.RegisterPatientDisease)
			patient.GET(":id/disease", ph.GetPatientDiseases)
			patient.DELETE(":id/disease/:patdiseaseid", ph.DeletePatientDisease)
			patient.PATCH(":id/disease/:diseaseid", ph.UpdatePatientDisease)

			patient.POST(":id/allergy", ph.RegisterPatientAllergy)
			patient.GET(":id/allergy", ph.GetPatientAllergies)
			patient.DELETE(":id/allergy/:patallergyid", ph.DeletePatientAllergy)
			patient.PATCH(":id/allergy/:allergyid", ph.UpdatePatientAllergy)
		}

		calendar := v1.Group("calendar")
		calendar.Use(middleware.BearerAuth(a))
		{
			calendar.GET("", ch.GetAllEvents)
			calendar.POST("", ch.CreateEvent)
			calendar.PATCH(":id", ch.UpdateEvent)
			calendar.DELETE(":id", ch.DeleteEvent)
			calendar.PATCH(":id/confirm", ch.ConfirmEvent)
			calendar.PATCH(":id/unconfirm", ch.UnconfirmEvent)
		}
	}

}

func SetupMiddleware(router *gin.Engine) {
	creds, validateheaders := false, false
	if os.Getenv("CORS_Credentials") == "true" {
		creds = true
	}
	if os.Getenv("CORS_ValidateHeaders") == "true" {
		validateheaders = true
	}
	router.Use(gin.Recovery(), gin.Logger(), cors.Middleware(cors.Config{
		Origins:         os.Getenv("CORS_Origins"),
		Methods:         os.Getenv("CORS_Methods"),
		RequestHeaders:  os.Getenv("CORS_RequestHeaders"),
		ExposedHeaders:  os.Getenv("CORS_ExposedHeaders"),
		MaxAge:          50 * time.Second,
		Credentials:     creds,
		ValidateHeaders: validateheaders,
	}))
	//cors shouldnot be allowing every orign
	router.SetTrustedProxies(nil)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To PATIENTAPI",
	})
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}
