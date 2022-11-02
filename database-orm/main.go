package main

//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type Product struct {
//	ID    int     `gorm:"primary_key"`
//	Name  string  `gorm:"type:varchar(100)"`
//	Price float64 `gorm:"type:decimal(10,2)"`
//}
//
//func main() {
//	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//
//	db.AutoMigrate(&Product{})
//
//	//db.Create(&Product{
//	//	Name:  "Product 1",
//	//	Price: 10.0,
//	//})
//	//
//	//products := []Product{
//	//	{
//	//		Name:  "Product 2",
//	//		Price: 20.0,
//	//	},
//	//	{
//	//		Name:  "Product 3",
//	//		Price: 30.0,
//	//	},
//	//	{
//	//		Name:  "Product 4",
//	//		Price: 40.0,
//	//	},
//	//}
//	//
//	//db.Create(&products)
//
//	//var product Product
//	////db.First(&product, 1)
//	//db.First(&product, "name = ?", "Product 2")
//	//fmt.Println(product)
//
//	//var products []Product
//	//db.Find(&products)
//	//db.Limit(2).Offset(2).Find(&products)
//	//db.Where("price > ?", 20.0).Find(&products)
//	//db.Where("name LIKE ?", "%Product%").Find(&products)
//	//for _, product := range products {
//	//	fmt.Println(product)
//	//}
//
//	//var p Product
//	//db.First(&p, 1)
//	//p.Name = "New Mouse"
//	//db.Save(&p)
//	//
//	//var p2 Product
//	//db.First(&p2, 1)
//	//fmt.Println(p2.Name)
//	//db.Delete(&p2)
//
//}
