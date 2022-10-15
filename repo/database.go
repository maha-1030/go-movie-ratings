package repo

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func init() {
	var (
		err           error
		connectionURL string = fmt.Sprintf("root:password@tcp(localhost:3306)/movies?charset=utf8&parseTime=True&loc=Local")
	)

	defer panicIfHasError(err, "Connecting Database")

	if db, err = gorm.Open("mysql", connectionURL); err != nil {
		fmt.Printf("\nError occured while connecting to database, err: %v\n", err)

		return
	}

	fmt.Println("\nSuccessfully connected to the movies database in mysql server")
	autoMigrate()
}

func autoMigrate() {
	var err error

	defer panicIfHasError(err, "Automigrating Models")

	if err = db.Debug().AutoMigrate(&Movie{}, &Rating{}).Error; err != nil {
		fmt.Printf("\nError occured while automigrating models into the database, err: %v\n", err)

		return
	}
}

func PopulateData() {
	var (
		err                   error
		moviesCSV, ratingsCSV *os.File
		movies                []*Movie
		ratings               []*Rating
	)

	defer panicIfHasError(err, "Populating Data")

	if moviesCSV, err = os.Open("./data/movies.csv"); err != nil {
		fmt.Printf("\nError occured while opening data file of movies, err: %v\n", err)

		return
	}

	if ratingsCSV, err = os.Open("./data/ratings.csv"); err != nil {
		fmt.Printf("\nError occured while opening data file of ratings, err: %v\n", err)

		return
	}

	if err = gocsv.Unmarshal(moviesCSV, &movies); err != nil {
		fmt.Printf("\nError occured while unmarshaling movies data into Movie struct format, err: %v\n", err)

		return
	}

	if err = gocsv.Unmarshal(ratingsCSV, &ratings); err != nil {
		fmt.Printf("\nError occured while unmarshaling ratings data into Rating struct format, err: %v\n", err)

		return
	}

	tx := db.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Debug().Delete(&Movie{}).Error; err != nil {
		fmt.Printf("\nError occured while deleting existing data from movies table, err: %v\n", err)

		return
	}

	if err = tx.Debug().Delete(&Rating{}).Error; err != nil {
		fmt.Printf("\nError occured while deleting existing data from ratings table, err: %v\n", err)

		return
	}

	for i, m := range movies {
		if err = tx.Create(m).Error; err != nil {
			fmt.Printf("\nError occured while inserting records into movies table, record no: %v, err: %v\n", i+1, err)

			return
		}
	}

	for i, r := range ratings {
		if err = tx.Create(r).Error; err != nil {
			fmt.Printf("\nError occured while inserting records into ratings table, record no: %v, err: %v\n", i+1, err)

			return
		}
	}

	tx.Commit()
	fmt.Printf("\nSuccessfully truncated tables and inserted %v records into movies table & %v records into ratings table\n", len(movies), len(ratings))
}

func panicIfHasError(err error, functionalityWithError string) {
	if err != nil {
		panic("\nERROR!!! " + functionalityWithError + "\n")
	}
}
