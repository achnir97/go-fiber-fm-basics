package lead

import (
	"github.com/achnir97/go-fiber-crm-basics/database"
	"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gofiber/fiber"
)


type Lead struct {
	gorm.Model 
	Name      string `json:"name"`
	Company   string  `json:"company"`
	Email     string  `json:"email"`
	Phone     int     `json:"phone"`
}

func GetLeads(c *fiber.Ctx){
	db:=database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)

}
//slice of lead means mulitle lead 
// its an arry of objects

func GetLead(c *fiber.Ctx) {
 id :=c.Params("id ")
 db:=database.DBConn
 var lead Lead
 db.Find(&lead, id)
c.JSON(lead) // go lang doesnot understand json so it has to convert what erver it receive in json through c
}

func NewLead(c *fiber.Ctx){
	db:=database.DBConn
	lead:= new(Lead)
	if err:= c.BodyParser(lead); err!=nil{
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx){
	id:=c.Params("id") // get the id 
	db:=database.DBConn // connnect to the 
	var lead Lead  // 
	db.First(&lead, id)
	if lead.Name==""{
		c.Status(500).Send("No lead found with ID")
		return 
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted ")
}
