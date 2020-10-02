package repository

import "gopkg.in/mgo.v2"

func createDatabaseSession(dbHost string) (*mgo.Session, error) {
	session, err := mgo.Dial(dbHost)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func DBExist(dbHost string, dbName string) (bool, error) {
	session, err := createDatabaseSession(dbHost)
	if err != nil {
		return false, err
	}
	defer session.Close()

	dbNames, err := session.DatabaseNames()
	if err != nil {
		return false, err
		//log.Println(backup.BackupMessage)
	}
	for _, name := range dbNames {
		if name == dbName {
			return true, nil
		}
	}
	return false, nil
}
