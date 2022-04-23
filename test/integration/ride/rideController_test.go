package ride

import (
	"os"
	"reby/cmd/server"
	"syscall"
	"testing"

	"github.com/golang-migrate/migrate"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:toor@tcp(localhost:3333)/prueba-tecnica-reby?charset=utf8mb4&parseTime=True&loc=Local"
const httpPort = 8080

type rideControllerTestSuite struct {
	suite.Suite
	dbConnectionStr string
	port            int
	dbConn          *gorm.DB
	dbMigration     *migrate.Migrate
}

func RideControllerTestSuite(t *testing.T) {
	suite.Run(t, &rideControllerTestSuite{})
}

func (s *rideControllerTestSuite) SetupSuite() {
	s.port = httpPort
	s.dbConnectionStr = dsn

	s.dbConn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	s.dbMigration, _ = migrate.New("file://../repositories/sql/migration", s.dbConnectionStr)

	s.dbMigration.Up()

	server := server.Server{
		Port:   httpPort,
		DBConn: s.dbConn,
	}

	go server.Start()
}

func (s *rideControllerTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *rideControllerTestSuite) SetupTest() {
	s.dbMigration.Up()
}

func (s *rideControllerTestSuite) TearDownTest() {
	s.dbMigration.Down()
}
