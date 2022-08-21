package notifications

import (
	"fmt"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

func NotifyManager(t entity.TaskResponse) {
	fmt.Printf("\nThe tech %v performed the task %d - (%v), on date %v\n", t.CreatedBy.Name, t.Id, t.Title, t.FinishedAt)
}
