package main

import(
	"database/sql"
	"fmt"
	"log"
	"userreg/api"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	//dsn := "root:123456@tcp(1c4ce498c4e7:3306)/user"
	dsn := "root:123456@tcp(localhost:3306)/user"

	db, err := sql.Open("mysql",dsn)
	if err != nil{
		log.Fatal("Error opening database", err)
	}
	defer db.Close()

	if err:=db.Ping(); err!=nil{
		log.Fatal("Error pingigng database:",err)
	}

	fmt.Println("Connected to database successfully!")

	//Register routes
	api.RegisterRoutes(db)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}