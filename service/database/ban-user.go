package database

func (db *appdbimpl) BanUser(banning string, banned string) (string, error) {

	res, err := db.c.Exec(`INSERT INTO Banned(Banned, Banning) VALUES (?,?)`, banned, banning)

	if err != nil {

		return "", err
	}
	_, err2 := res.RowsAffected()
	if err2 != nil {

		return "", err2
	}

	return banned, nil
}
