package database

func (db *appdbimpl) GetListLike(photoid int) (arrayoflike []Like, err error) {

	rows, err1 := db.c.Query("SELECT * FROM Like WHERE PhotoID = ? ", photoid)
	if err1 != nil {
		return arrayoflike, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var like Like
		err2 := rows.Scan(&like.LikeID, &like.PhotoID, &like.UserRec)
		if err2 != nil {

			return arrayoflike, err
		}

		arrayoflike = append(arrayoflike, like)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return arrayoflike, nil
}
