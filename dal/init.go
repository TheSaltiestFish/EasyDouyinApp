package dal

import "github.com/TheSaltiestFish/EasyDouyinApp/dal/db"

// Init init dal
func Init() {
	db.InitDB() // mysql
}
