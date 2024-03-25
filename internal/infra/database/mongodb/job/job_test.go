package job_test

import (
	"go.mongodb.org/mongo-driver/mongo"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	db *mongo.Database
)

type StackRepositoryTestSuite struct {
	suite.Suite
}

func TestStackRepository(t *testing.T) {
	suite.Run(t, new(StackRepositoryTestSuite))
}

func (suite *StackRepositoryTestSuite) SetupSubTest() {
}
