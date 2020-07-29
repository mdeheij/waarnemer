package repository

import "waarnemer/model"

type CheckRepository struct {
}

// FindAllChecks is currently a stub to return all known checks
func (CheckRepository) FindAllChecks() []model.Check {
	return []model.Check{}
}
