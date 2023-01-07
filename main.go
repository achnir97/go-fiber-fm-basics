package main 
import (
	"github.com/achnir97/go-fiber-crm-basics/database"
	"github.com/achnir97/go-fiber-crm-basics/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"fmt")

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
	
}


func initDatabase() {
var err error 
database.DBConn, err= gorm.Open("sqlite3", "leads.db") //gorm is opening the databse called leads.db 
if err!=nil {
	panic("Faileed to connect database")
}
fmt.Println("Connection opened to database")
database.DBConn.AutoMigrate(&lead.Lead{})
}

func main () {
	app:=fiber.New()
	initDatabase() 
	setupRoutes(app)
	app.Listen(5000)
	defer database.DBConn.Close() // it will be run after everything inside is execcute

}