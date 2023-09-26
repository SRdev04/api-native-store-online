package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx) {

	err := recover()
	if err != nil {
		//jika error akan kita rollback
		MenghasilkanError := tx.Rollback()
		if err != nil {

			panic(MenghasilkanError)
		}
		panic(err)
	} else {

		err := tx.Commit()
		IfError(err)
		log.Println("Commit")

	}

}
