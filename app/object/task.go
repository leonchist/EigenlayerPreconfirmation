package object

type TaskObjRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TaskUpdateObjRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TaskObjResponse struct {
	ID        int
	Name      string
	Status    string
	CreatedAt string
	UpdatedAt string
}
