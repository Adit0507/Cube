package task

type State int 

const (
	Pending State = iota
	Scheduled 	//means system has figured out where to run the task
	Running
	Completed
	Failed
)

