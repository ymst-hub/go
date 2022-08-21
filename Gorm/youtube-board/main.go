package main

import (
	"mymodule/my"
	_ "gorm.io/driver/sqlite"
)

//main program
func main(){
	my.Migrate()
}