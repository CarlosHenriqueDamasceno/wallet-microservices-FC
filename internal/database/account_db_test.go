package database

import (
	"database/sql"
	"testing"

	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	s.accountDB = *NewAccountDB(db)
	s.client, _ = entity.NewClient("Carlos", "carlos@teste.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestGet() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at) values (?,?,?,?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
	persistedAccount, err := s.accountDB.Get(account.ID)
	s.Nil(err)
	s.Equal(account.ID, persistedAccount.ID)
	s.Equal(account.Balance, persistedAccount.Balance)
	s.Equal(account.Client.ID, persistedAccount.Client.ID)
	s.Equal(account.Client.Name, persistedAccount.Client.Name)
	s.Equal(account.Client.Email, persistedAccount.Client.Email)

}
