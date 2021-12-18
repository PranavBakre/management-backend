package error

type Error struct {
	Code        int    `json:"code"`
	ErrorType   string `json:"error"`
	Description string `json:"description"`
}

func (err Error) Error() string {
	return err.ErrorType + ":" + err.Description
}
