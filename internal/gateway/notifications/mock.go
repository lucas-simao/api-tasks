package notifications

import (
	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockNotifications struct {
	mock.Mock
}

func (ref *MockNotifications) NotifyManager(t entity.TaskResponse) {
	return
}
