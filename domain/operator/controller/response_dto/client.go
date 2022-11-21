package response_dto

type Client struct {
	ID       uint   `json:"id"`
	FileName string `json:"fileName"`
	Config   string `json:"config"`
}
