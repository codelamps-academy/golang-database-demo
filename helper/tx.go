package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic("Error")
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
