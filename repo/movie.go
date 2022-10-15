package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Movie struct {
	Tconst         string `gorm:"primary_key" csv:"tconst"`
	TitleType      string `gorm:"title_type" csv:"titleType"`
	PrimaryTitle   string `gorm:"primary_title" csv:"primaryTitle"`
	RuntimeMinutes int    `gorm:"runtime_minutes" csv:"runtimeMinutes"`
	Genres         string `gorm:"genres" csv:"genres"`
}

func (m *Movie) Validate() error {
	if m.Tconst == "" {
		return fmt.Errorf("parameter 'Tconst' is missing")
	}

	if m.TitleType == "" {
		return fmt.Errorf("parameter 'TitleType' is missing")
	}

	if m.PrimaryTitle == "" {
		return fmt.Errorf("parameter 'PrimaryTitle' is missing")
	}

	if m.RuntimeMinutes < 1 {
		return fmt.Errorf("invalid parameter 'RuntimeMinutes' should be a positive interger")
	}

	if m.Genres == "" {
		return fmt.Errorf("parameter 'Genres' is missing")
	}

	return nil
}

func (m *Movie) Save() error {
	if err := db.Debug().Create(m).Error; err != nil {
		return err
	}

	return nil
}

func (m *Movie) GetByTconst(tconst string) (*Movie, error) {
	if err := db.Debug().First(m, "tconst = ?", tconst).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return m, nil
}

func (m *Movie) GetTopMoviesByField(field, order string, count int) ([]Movie, error) {
	movies := []Movie{}

	if err := db.Debug().Model(m).Order(field + " " + order).Limit(count).Scan(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}
