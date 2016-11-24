package selecter

import (
	"errors"

	"github.com/khezen/darwin/environment/population"
)

type proportionalToRankingSelecter struct{}

func (s proportionalToRankingSelecter) Select(pop *population.Population, survivorsSize uint) (*population.Population, error) {
	err := checkArgs(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	return nil, errors.New("unsupported operation")
}

// NewProportionalToRankingSelecter is the constrctor for truncation selecter
func NewProportionalToRankingSelecter() Interface {
	return proportionalToRankingSelecter{}
}