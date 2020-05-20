package sqlite

import (
	"database/sql"
	"io/ioutil"
	"testing"

    _ "github.com/mattn/go-sqlite3"
)

func newTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("sqlite3", "./testing_db.db")
	if err != nil {
		t.Fatal(err)
	}

	script, err := ioutil.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {

		script, err := ioutil.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	}
}
