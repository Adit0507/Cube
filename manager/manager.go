package manager

import (
	"fmt"

	"github.com/Adit0507/cube/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorkers() {
	fmt.Println("I will select an appropriate worker")
}

func (m *Manager) UpdateTasks() {
 fmt.Println("I will update tasks")
}

func (m *Manager) SendWork() {
 fmt.Println("I will send work to workers")
}
