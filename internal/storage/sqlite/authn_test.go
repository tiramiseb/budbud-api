package sqlite_test

import (
	"database/sql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "github.com/tiramiseb/budbud-api/internal/storage/sqlite"
	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

const tokenSQL = `
INSERT INTO user VALUES ("test@example.com", "X", "X");
`

var _ = Describe("Authentication methods", func() {
	var (
		db DB
	)

	BeforeEach(func() {
		db = NewDB()
		Ω(db.ExecFile("sqlite-init.sql")).Should(Succeed())
	})
	AfterEach(func() {
		Ω(db.Close()).Should(Succeed())
	})

	It("should deal correctly with a token", func() {
		Ω(db.Exec(`INSERT INTO user VALUES ("test@example.com", "HASH", "SALT");`)).Should(Succeed())
		Ω(db.Service.AddToken("the_token", "test@example.com")).Should(Succeed())
		user, err := db.Service.GetUserFromToken("the_token")
		Ω(user).Should(Equal(model.User{ID: "test@example.com", Email: "test@example.com"}))
		Ω(err).Should(BeNil())
		Ω(db.Service.RemoveToken("the_token")).Should(Succeed())
		user, err = db.Service.GetUserFromToken("the_token")
		Ω(user).Should(Equal(model.User{}))
		Ω(err).Should(MatchError(sql.ErrNoRows))
	})

})
