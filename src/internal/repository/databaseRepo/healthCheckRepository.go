package databaseRepo

func (app *ApplicationRepo) TestDatabase() error {
	conn, _ := app.getDB().DB()
	err := conn.Ping()
	if err != nil {
		//logger.Error.Println("Database failed", err)
		return err
	}

	return nil
}
