package database

import (
	"database/sql"
	"testing"

	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_from_id varchar(255), account_to_id varchar(255), amount float, created_at date)")

	client, err := entity.NewClient("Carlos", "carlos@teste.com")
	s.Nil(err)
	s.client = client
	client2, err := entity.NewClient("Carlos 2", "carlos2@teste.com")
	s.Nil(err)
	s.client = client2

	accountFrom := entity.NewAccount(client)
	accountFrom.Credit(1000)
	s.accountFrom = accountFrom
	accountTo := entity.NewAccount(client2)
	accountTo.Credit(1000)
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestSave() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Save(transaction)
	s.Nil(err)
}