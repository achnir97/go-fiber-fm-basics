package database 
import (
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/sqlite" // we are using the language sqlite in gorm 
)

var ( DBConn *gorm.DB)