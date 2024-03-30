package databaseRepo

import (
	"database/sql"
	"movie-journal-api/internal/logger"
)

func (app *ApplicationRepo) TestDatabase() error {
	conn, _ := app.getDB().SQL.DB()
	err := conn.Ping()
	if err != nil {
		logger.Info("test", "ok")
		return err
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			logger.Error("DB_CLOSE_ERROR", "Couldn't close DB", "", err)
		}
	}(conn)
	return nil
}
