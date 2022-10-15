package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Rating struct {
	Tconst        string  `gorm:"primary_key" csv:"tconst"`
	AverageRating float64 `gorm:"average_rating" csv:"averageRating"`
	NumVotes      int     `gorm:"num_votes" csv:"numVotes"`
}

func (r *Rating) Validate() error {
	if r.Tconst == "" {
		return fmt.Errorf("parameter 'Tconst' is missing")
	}

	if r.AverageRating < 0 {
		return fmt.Errorf("invalid parameter 'AverageRating' should be a positive value")
	}

	if r.NumVotes < 1 {
		return fmt.Errorf("invalid parameter 'NumVotes' should be a positive interger")
	}

	return nil
}

func (r *Rating) Save() error {
	if err := db.Debug().Create(r).Error; err != nil {
		return err
	}

	return nil
}

func (r *Rating) GetByTconst(tconst string) (*Rating, error) {
	if err := db.Debug().First(r, "tconst = ?", tconst).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return r, nil
}
