package apifuncs

type PrintData struct {
	FilePath string `json:"file_path,omitempty"`
}

type PrintStatus struct {
	IsFinished bool `json:"is_finished,omitempty"`
}
