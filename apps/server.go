package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pl = fmt.Println

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName 	string
	AppEnv 		string
	AppPort 	string
}

type DBConfig struct {
	DBHost		string
	DBUser		string
	DBPassword 	string
	DBName 		string
	DBPort		string
}

func (server *Server) Initialize(AppConfig AppConfig, DBConfig DBConfig){
	pl("welcome to " + AppConfig.AppName)


	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			DBConfig.DBHost,
			DBConfig.DBUser,
			DBConfig.DBPassword,
			DBConfig.DBName,
			DBConfig.DBPort,
		)
	
	server.DB , err =gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}
		pl("succes to connect database")

	server.initalizeRoutes() 

	for _, model := range RegisterModels(){
		err = server.DB.Debug().AutoMigrate(model.Model)

		if err!= nil{
			log.Fatal(err)
		}

	}
	pl("Database migrated")

}

func (server *Server) Run(addr string){
	pl("Listening to Port", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string)string{
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run(){
	var server = Server{}
	var AppConfig = AppConfig{}
	var DBConfig = DBConfig{}

	err := godotenv.Load()
	if err !=nil {
		log.Fatalf("Error geting .env File")
	}


	AppConfig.AppName	 = getEnv("APP_NAME", "GO TOKO")
	AppConfig.AppEnv 	 = getEnv("APP_ENV", "development")
	AppConfig.AppPort 	 = getEnv("APP_PORT", "8080")
	
	DBConfig.DBHost 	= getEnv("DB_HOST", "localhost")
	DBConfig.DBUser		= getEnv("DB_USER", "postgres")
	DBConfig.DBPassword = getEnv("DB_PASSWORD", "12345678")
	DBConfig.DBName		= getEnv("DB_NAME", "go")
	DBConfig.DBPort		= getEnv("DB_PORT", "5432")


	server.Initialize(AppConfig, DBConfig)
	server.Run(":" + AppConfig.AppPort)

}