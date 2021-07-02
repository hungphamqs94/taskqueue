package src

type Worker interface {
	Work(j *Job)
}
