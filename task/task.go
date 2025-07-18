package task

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
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
	Memory        int64             //will help the system identify no. of resources a task needs
	Disk          int64             //this too
	ExposedPorts  nat.PortSet       //used by Docker to ensure machines allocates the proper network ports for the task
	PortBindings  map[string]string //same as above
	RestartPolicy string            //tells the system what to do in event a task stops or fails unexpectedly
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}

type Config struct {
	Name          string
	AttachStdin   bool
	AttachStdout  bool
	AttachStderr  bool
	ExposedPorts  nat.PortSet
	Cmd           []string
	Image         string //holds name of the image the container will run
	Cpu           float64
	Memory        int64 //scheduler will use Memory and Disk to find anode in cluster capable of running task
	Disk          int64
	Env           []string
	RestartPolicy string //tells docker daemon what to do if container stops unexpectedly
}

func NewConfig(t *Task) *Config {
	return &Config{
		Name:         t.Name,
		ExposedPorts: t.ExposedPorts,
		Image:        t.Image,
		Cpu:          t.CPU,
		Memory:       t.Memory,
	}
}

type Docker struct {
	Client *client.Client
	Config Config
}

type DockerResult struct {
	Error       error
	Action      string
	ContainerID string
	Result      string
}

func (d *Docker) Run() DockerResult {
	ctx := context.Background()

	reader, err := d.Client.ImagePull(ctx, d.Config.Image, types.ImagePullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v\n", d.Config.Image, err) 
		return DockerResult{Error: err}
	}

	io.Copy(os.Stdout, reader)

}
