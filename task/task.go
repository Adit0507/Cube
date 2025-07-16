package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Pending   State = iota
	Scheduled       //means system has figured out where to run the task
	Running
	Completed
	Failed
)

type Task struct {
	ID            uuid.UUID
	ContainerID   string
	Name          string
	State         State
	Image         string
	CPU           float64
	Memory        int               //will help the system identify no. of resources a task needs
	Disk          int               //this too
	ExposedPorts  nat.PortSet       //used by Docker to ensure machines allocates the proper network ports for the task
	PortBindings  map[string]string //same as above
	RestartPolicy string            //tells the system what to do in event a task stops or fails unexpectedly
	StartTime     time.Time
	FinishTime    time.Time
}
