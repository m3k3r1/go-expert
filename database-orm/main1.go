package main

//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type Category struct {
//	ID   int    `gorm:"primary_key"`
//	Name string `gorm:"type:varchar(100)"`
//}
//
//type Product struct {
//	ID         int     `gorm:"primary_key"`
//	Name       string  `gorm:"type:varchar(100)"`
//	Price      float64 `gorm:"type:decimal(10,2)"`
//	CategoryID int
//	Category   Category
//	gorm.Model
//}
//
//func main() {
//	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//
//	db.AutoMigrate(&Product{}, &Category{})
//
//	category := Category{Name: "Category 1"}
//	db.Create(&category)
//
//	db.Create(&Product{
//		Name:       "Product 1",
//		Price:      10.0,
//		CategoryID: category.ID,
//	})
//
//	var products []Product
//	db.Preload("Category").Find(&products)
//	for _, product := range products {
//		fmt.Println(product.Name, product.Category.Name)
//	}
//}
