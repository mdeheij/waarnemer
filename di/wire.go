// +build wireinject

package di

import (
	"github.com/google/wire"
	"waarnemer/repository"
)

func InitializeCheckRepository() *repository.CheckRepository {
	panic(wire.Build(repository.CheckRepositorySet))
}