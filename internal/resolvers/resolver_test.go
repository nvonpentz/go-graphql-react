package resolvers

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	*Resolver
}

func (suite *TestSuite) SetupSuite() {
	resolver, err := NewWithDefaults()
	suite.Require().NoError(err, "Failed to setup test suite.")
	suite.Resolver = resolver
}

func (suite *TestSuite) SetupTest() {
	suite.Resolver.Service.Postgres.ClearTables()
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
