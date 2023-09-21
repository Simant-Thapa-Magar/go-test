package main

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDBObj(mockDB *sql.DB) *Service {
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	service := Service{
		database: db,
	}

	return &service
}

func TestAddBook(t *testing.T) {
	mockDB, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	service := getDBObj(mockDB)

	bookSample := Book{
		Id:            1,
		Name:          "Sample Book",
		Author:        "Sample Author",
		PublishedDate: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT").WithArgs(bookSample.Id, bookSample.Name, bookSample.Author, bookSample.PublishedDate).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	addErr := service.AddBook(bookSample)

	if addErr != nil {
		t.Errorf("Got error while inserting book: %s", addErr)
	}
}

func TestGetBookByYear(t *testing.T) {

	mockDB, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	service := getDBObj(mockDB)

	bookSample := Book{
		Id:            1,
		Name:          "Sample Book",
		Author:        "Sample Author",
		PublishedDate: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"Name", "Author", "PublishedDate"}).AddRow(bookSample.Name, bookSample.Author, bookSample.PublishedDate)

	mock.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

	service.GetBookByYear(2023)
}

func TestGetBookByYear2(t *testing.T) {

	mockDB, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	service := getDBObj(mockDB)

	emptyRow := sqlmock.NewRows([]string{"Name", "Author", "PublishedDate"})

	mock.ExpectQuery("SELECT(.*)").WithArgs(2020).WillReturnRows(emptyRow)

	service.GetBookByYear(2020)
}
