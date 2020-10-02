package backup

import "gopkg.in/mgo.v2"

const (
	gamesFilePath = "user_games.json"
	usersFilePath = "users.json"

	BackupMessage = "backup: Restoring DB has started"
)

func CreateDB() (*mgo.Database, error) {

	return nil, nil
}

func checkBackupFiles() error {

	return nil
}

func readUsersFromFile() {

}

func readGamesFromFile() {

}
