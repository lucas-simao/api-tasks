package notifications

import (
	"fmt"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

type Notifications interface {
	NotifyManager(t entity.TaskResponse)
}

func New() Notifications {
	return notifications{}
}

type notifications struct{}

func (n notifications) NotifyManager(t entity.TaskResponse) {
	fmt.Printf("\nThe tech %v performed the task %d - (%v), on date %v\n", t.CreatedBy.Name, t.Id, t.Title, t.FinishedAt)
}
