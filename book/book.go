package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	Id            int64     `gorm:"column:Id"`
	Name          string    `gorm:"column:Name"`
	Author        string    `gorm:"column:Author"`
	PublishedDate time.Time `gorm:"column:PublishedDate"`
}

type Service struct {
	database *gorm.DB
}

func initializeDatabase() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func main() {
	db, err := initializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	service := &Service{
		database: db,
	}

	service.TruncateTable()

	books := GetSomeBooks()

	for _, book := range books {
		e := service.AddBook(book)
		if e != nil {
			log.Fatal(e)
		}
	}

	booksFor1904, err := service.GetBookByYear(1904)

	if err != nil {
		log.Fatal(err)
	}

	if len(booksFor1904) > 0 {
		fmt.Println("Books published in 1904 are:")
		for i, book := range booksFor1904 {
			fmt.Printf("%d %s by %s\n", i+1, book.Name, book.Author)
		}
	}
}

func (s *Service) AddBook(book Book) error {
	err := s.database.Create(&book).Error
	return err
}

func (s *Service) GetBookByYear(year int) ([]Book, error) {
	var books []Book
	err := s.database.Model(&Book{}).Where("YEAR(PublishedDate) = ?", year).Find(&books).Error
	return books, err
}

func GetSomeBooks() []Book {
	var books []Book

	publishDate, _ := time.Parse("2006-01-02", "1913-01-01")

	books = append(books, Book{
		Name:          "In Search of Lost Time",
		Author:        "Marcel Proust",
		PublishedDate: publishDate,
	})

	publishDate2, _ := time.Parse("2006-01-02", "1904-01-01")

	books = append(books, Book{
		Name:          "Ulysses",
		Author:        "James Joyce",
		PublishedDate: publishDate2,
	})

	publishDate3, _ := time.Parse("2006-01-02", "1599-01-01")

	books = append(books, Book{
		Name:          "Hamlet",
		Author:        "William Shakespeare",
		PublishedDate: publishDate3,
	})

	return books
}

func (s *Service) TruncateTable() error {
	return s.database.Where("1=1").Delete(&Book{}).Error
}
