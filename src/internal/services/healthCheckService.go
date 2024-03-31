package services

import (
	"movie-journal-api/internal/repository/databaseRepo"
)

func (as *ApplicationService) CheckDBConnection() error {
	repo := databaseRepo.NewApplicationRepo(as.Application)
	return repo.TestDatabase()
}
