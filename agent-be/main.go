package main

import (
	"agent-app/auth0"
	"agent-app/handler"
	"agent-app/model"
	"agent-app/repository"
	"agent-app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

var db *gorm.DB
var err error

func initDB() *gorm.DB {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB")

	connString := fmt.Sprintf("host=localhost port=%s user=%s dbname=%s sslmode=disable password=%s", port, user, dbName, pass)
	db, err = gorm.Open("postgres", connString)

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(model.User{})
	db.AutoMigrate(model.Company{})
	db.AutoMigrate(model.Comment{})
	return db
}

func initUserRepo(database *gorm.DB) *repository.UserRepository {
	return &repository.UserRepository{Database: database}
}

func initCompanyRepo(database *gorm.DB) *repository.CompanyRepository {
	return &repository.CompanyRepository{Database: database}
}

func initAuth0Client() *auth0.Auth0Client {
	domain := os.Getenv("AUTH0_DOMAIN")
	clientId := os.Getenv("AUTH0_CLIENT_ID")
	clientSecret := os.Getenv("AUTH0_CLIENT_SECRET")
	audience := os.Getenv("AUTH0_AUDIENCE")

	client := auth0.NewAuth0Client(domain, clientId, clientSecret, audience)
	return &client
}

func initUserService(userRepo *repository.UserRepository, auth0Client *auth0.Auth0Client) *service.UserService {
	return &service.UserService{UserRepo: userRepo, Auth0Client: *auth0Client}
}

func initCompanyService(companyRepo *repository.CompanyRepository, auth0Client *auth0.Auth0Client) *service.CompanyService {
	return &service.CompanyService{CompanyRepo: companyRepo, Auth0Client: *auth0Client}
}

func initUserHandler(service *service.UserService) *handler.UserHandler {
	return &handler.UserHandler{Service: service}
}

func initCompanyHandler(service *service.CompanyService) *handler.CompanyHandler {
	return &handler.CompanyHandler{Service: service}
}

func handleUserFunc(handler *handler.UserHandler, router *gin.Engine) {
	router.POST("register", handler.Register)
	router.GET("users", handler.GetByEmail)
}

func handleCompanyFunc(handler *handler.CompanyHandler, router *gin.Engine) {
	router.POST("company", handler.Register)
	router.POST("company/approve", handler.Approve)
	router.GET("company", handler.GetAllCompanies)
	router.GET("companyRequests", handler.GetAllCompanyRequests)
}

func initCommentRepo(database *gorm.DB) *repository.CommentRepository {
	return &repository.CommentRepository{Database: database}
}

func initCommentService(commentRepo *repository.CommentRepository) *service.CommentService {
	return &service.CommentService{CommentRepo: commentRepo}
}

func initCommentHandler(service *service.CommentService) *handler.CommentHandler {
	return &handler.CommentHandler{Service: service}
}

func handleCommentFunc(handler *handler.CommentHandler, router *gin.Engine) {
	router.POST("comment", handler.AddComment)
	router.GET("comment/:id", handler.GetCommentByID)
	router.DELETE("comment/:id", handler.DeleteComment)
	router.PUT("comment/:id", handler.UpdateComment)
	router.GET("search/comment/:id/owner", handler.GetCommentByOwnerID)
	router.GET("search/comment/:id/company", handler.GetCommentByCompanyID)
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database := initDB()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	auth0Client := initAuth0Client()

	userRepo := initUserRepo(database)
	userService := initUserService(userRepo, auth0Client)
	userHandler := initUserHandler(userService)

	companyRepo := initCompanyRepo(database)
	companyService := initCompanyService(companyRepo, auth0Client)
	companyHandler := initCompanyHandler(companyService)

	commentRepo := initCommentRepo(database)
	commentService := initCommentService(commentRepo)
	commentHandler := initCommentHandler(commentService)
	/*commentRepo.AddComment(&model.Comment{
		CompanyID:    1234,
		UserOwnerID:  1234,
		Salary:       3667.25,
		Position:     "BOSS",
		CreationDate: time.Now(),
		Description:  "AAAAAAAAAAAAAAAAAAA",
	})*/
	router := gin.Default()

	handleUserFunc(userHandler, router)
	handleCompanyFunc(companyHandler, router)
	handleCommentFunc(commentHandler, router)

	http.ListenAndServe(port, cors.AllowAll().Handler(router))
}
