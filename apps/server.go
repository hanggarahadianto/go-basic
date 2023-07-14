package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"origin/database/seeders"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"

	// "github.com/urfave/cli"

	// "github.com/urfave/cli"
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

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig){
	pl("welcome to " + appConfig.AppName)
	server.initalizeRoutes() 

}


func (server *Server) Run(addr string){
	pl("Listening to Port", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}


func (server *Server) initializeDB(dbConfig DBConfig){
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			dbConfig.DBHost,
			dbConfig.DBUser,
			dbConfig.DBPassword,
			dbConfig.DBName,
			dbConfig.DBPort,
		)
	
	server.DB , err =gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}
		pl("succes to connect database")

}

	// seeders.DBSeed(server.DB)


func (server *Server) dbMigrate(){
	for _, model := range RegisterModels(){
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err!= nil{
			log.Fatal(err)
		}

	}
	pl("Database migrated")
}


func (server *Server) initCommands(config AppConfig, dbConfig DBConfig){
	server.initializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context)error{
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}
	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


func getEnv(key, fallback string)string{
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run(){
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err !=nil {
		log.Fatalf("Error geting .env File")
	}




	appConfig.AppName	 = getEnv("APP_NAME", "GO TOKO")
	appConfig.AppEnv 	 = getEnv("APP_ENV", "development")
	appConfig.AppPort 	 = getEnv("APP_PORT", "8080")
	
	dbConfig.DBHost 	= getEnv("DB_HOST", "localhost")
	dbConfig.DBUser		= getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "12345678")
	dbConfig.DBName		= getEnv("DB_NAME", "go")
	dbConfig.DBPort		= getEnv("DB_PORT", "5432")

	// -------------------------------

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(appConfig, dbConfig)
	}else{
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)

	}
	

}