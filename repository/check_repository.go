package repository

import (
	"waarnemer/model"

	"github.com/google/wire"
)

var CheckRepositorySet = wire.NewSet(NewCheckRepository)

type CheckRepository struct {
}

func NewCheckRepository() *CheckRepository {
	return &CheckRepository{}
}

// FindAllChecks is currently a stub to return all known checks
func (CheckRepository) FindAllChecks() []model.Check {
	return []model.Check{
		model.Check{
			Identifier: "sample_1",
			Type:       "DUMMY",
		},
		model.Check{
			Identifier: "sample_2",
			Type:       "DUMMY",
		},
	}
}
