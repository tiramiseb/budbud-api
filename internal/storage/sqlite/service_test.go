package sqlite_test

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/tiramiseb/budbud-api/internal/storage/sqlite"
)

type DB struct {
	path    string
	Service *Service
}

func NewDB() DB {
	tmpfile, err := ioutil.TempFile("", "budbud-sqlite-test")
	if err != nil {
		panic(err)
	}
	path := tmpfile.Name()
	tmpfile.Close()
	service, err := New(path)
	if err != nil {
		panic(err)
	}
	return DB{path, service}
}

func (d DB) ExecFile(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return d.Exec(string(content))
}

func (d DB) Exec(query string) error {
	dsn := fmt.Sprintf("file:%s?_foreign_keys=true&cache=shared", d.path)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return db.Close()
}

func (d DB) Close() error {
	if err := d.Service.Close(); err != nil {
		return err
	}
	return os.Remove(d.path)
}

var _ = Describe("Sqlite service", func() {
	It("should open and close a SQLite service with a database in memory", func() {
		s, err := New(":memory:")
		Ω(err).Should(Succeed())
		err = s.Close()
		Ω(err).Should(Succeed())
	})
})
