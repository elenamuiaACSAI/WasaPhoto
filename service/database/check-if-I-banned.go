package database

func (db *appdbimpl) CheckIfIBanned(userid string, idget string) (bool, error) {
	rows, err := db.c.Query(`SELECT Banned,Banning FROM Banned WHERE Banned = ? AND Banning = ? `, idget, userid)
	if err != nil {
		return true, err
	}
	defer func() { _ = rows.Close() }()

	if rows.Next() {
		return true, nil
	}

	if err = rows.Err(); err != nil {
		return false, err
	}
	return false, nil
}
