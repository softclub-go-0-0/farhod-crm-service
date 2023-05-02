package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/crm-service/pkg/handlers"
	"github.com/softclub-go-0-0/crm-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func DBInit(user, password, dbname, port string) (*gorm.DB, error) {
	dsn := "host=localhost" +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable" +
		" TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Teacher{},
		&models.Course{},
		&models.Timetable{},
		&models.Group{},
		&models.Student{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	DBName := flag.String("dbname", "crm_service", "Enter the name of DB")
	DBUser := flag.String("dbuser", "postgres", "Enter the name of a DB user")
	DBPassword := flag.String("dbpassword", "postgres", "Enter the password of user")
	DBPort := flag.String("dbport", "5432", "Enter the port of DB")
	flag.Parse()

	db, err := DBInit(*DBUser, *DBPassword, *DBName, *DBPort)
	if err != nil {
		log.Fatal("db connection error:", err)
	}

	log.Println("successfully connected to DB")

	h := handlers.NewHandler(db)

	router := gin.Default()

	//router.Use(AuthMiddleware("somekey"))

	router.GET("/teachers", h.GetAllTeachers)
	router.POST("/teachers", h.CreateTeacher)
	router.GET("/teachers/:teacherID", h.GetOneTeacher)
	router.PUT("/teachers/:teacherID", h.UpdateTeacher)
	router.DELETE("/teachers/:teacherID", h.DeleteTeacher)

	router.GET("/courses", h.GetAllCourses)
	router.POST("/courses", h.CreateCourse)
	router.GET("/courses/:courseID", h.GetOneCourse)
	router.PUT("/courses/:courseID", h.UpdateCourse)
	router.DELETE("/courses/:courseID", h.DeleteCourse)

	router.GET("/timetables", h.GetAllTimetables)
	router.POST("/timetables", h.CreateTimetable)
	router.GET("/timetables/:timetableID", h.GetOneTimetable)
	router.PUT("/timetables/:timetableID", h.UpdateTimetable)
	router.DELETE("/timetables/:timetableID", h.DeleteTimetable)

	groups := router.Group("/groups")
	{
		groups.GET("/", h.GetAllGroups)
		groups.POST("/", h.CreateGroup)
		groups.GET("/:groupID", h.GetOneGroup)
		groups.PUT("/:groupID", h.UpdateGroup)
		groups.DELETE("/:groupID", h.DeleteGroup)

		students := groups.Group("/:groupID/students")
		{
			students.GET("/", h.GetAllTeachers)
			// implemented
			students.POST("/", h.CreateStudent)
			students.GET("/:studentID", h.GetOneTeacher)
			students.PUT("/:studentID", h.UpdateTeacher)
			students.DELETE("/:studentID", h.DeleteTeacher)
		}
	}

	router.Run(":4000")
}

func AuthMiddleware(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bracelet := c.GetHeader("bracelet")
		if bracelet != key {
			log.Println("unauthorized access")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Who are you, warrior?",
			})
			return
		}
		c.Next()
	}
}
