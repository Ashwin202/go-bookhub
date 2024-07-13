package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Ashwin202/go-bookhub/pkg/config"
)

var db *gorm.DB

type Book struct {	
	gorm.Model // Embeds the gorm.Model struct from GORM, which adds fields like ID, CreatedAt, UpdatedAt, and DeletedAt to the Book struct. These fields are commonly used for database operations in GORM.
	Name string `gorm:"" json:"name"` // This is a GORM tag that can be used to specify additional configuration for database operations. In this case, it's empty, meaning there's no additional configuration specified.
// Name        string `gorm:"type:varchar(100);unique_index" json:"title"` ==> pecifies that the Name field in the database should be of type VARCHAR with a maximum length of 100 characters. This overrides the default type inferred by GORM.
	Author string `json:"author"` //This tag specifies that in JSON format, the field should be named author.
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.DB
	db.AutoMigrate(&Book{}) //auto migrates the Book schema (creates the books table if it doesn't exist), and creates a new book record with the specified fields.
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // The NewRecord method in GORM is used to check whether a struct represents a new record that hasn't been persisted to the database yet. 
	db.Create(&b) // The Create method in GORM is used to insert a new record into the database based on the struct 
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("id = ?", Id).Find(&getbook)
	return &getbook, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("id = ?", Id).Delete(&book)
	return book
}