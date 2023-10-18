package database

import (
	"database/sql"
	"testing"

	"github.com/CarlosHenriqueDamasceno/wallet-ms-fc/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	ClientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.ClientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client, _ := entity.NewClient("Carlos", "carlos@teste.com")
	err := s.ClientDB.Save(*client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("Carlos", "carlos@teste.com")
	s.ClientDB.Save(*client)
	persistedClient, err := s.ClientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, persistedClient.ID)
	s.Equal(client.Name, persistedClient.Name)
	s.Equal(client.Email, persistedClient.Email)
}
