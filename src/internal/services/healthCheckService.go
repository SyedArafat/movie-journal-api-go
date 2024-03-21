package services

import (
	"movie-journal-api/internal/repository/databaseRepo"
)

func (aS *ApplicationService) CheckDBConnection() error {
	repo := databaseRepo.NewApplicationRepo(aS.Application)
	return repo.TestDatabase()
}
