package src

// Job represents a interface of job that can be enqueued into a dispatcher.
type Job struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
