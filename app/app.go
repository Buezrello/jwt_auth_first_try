package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/hexagonal-auth/domain"
	"example.com/hexagonal-auth/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	// sanityCheck()
	router := mux.NewRouter()
	authReposditory := domain.NewAuthRepository(getDbClient())
	ah := AuthHandler{service: service.NewLoginService(authReposditory, domain.GetRolePermissions())}

	router.HandleFunc("/auth/Login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplemented).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Println(fmt.Sprintf("Starting OAuth server on %s:%s", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
