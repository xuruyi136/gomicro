package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//http://v1.gorm.io/docs/connecting_to_the_database.html
var (
	db *gorm.DB
)

//https://blog.csdn.net/qq_37767455/article/details/104805762 一个有关Golang变量作用域的坑
/*
select User,authentication_string,Host from user;
flush privileges;
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '123456';
*/
func init() {
	db, _ = gorm.Open("mysql", "root:123456@(121.196.195.242:3306)/go_micro?charset=utf8&parseTime=True&loc=Local")

}
func GetDB() *gorm.DB {

	return db
}
