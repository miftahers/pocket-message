package services

import (
	m "pocket-message/services/mock"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PocketMessageSuite struct {
	suite.Suite
	service PocketMessageServices
}

func TestSuitePocketMessage(t *testing.T) {
	suite.Run(t, new(PocketMessageSuite))
}

func (s *PocketMessageSuite) SetupSuite() {
	service := NewPocketMessageServices(&m.MockGorm{})
	s.service = service
}

func (s *PocketMessageSuite) TearDownSuite() {}
