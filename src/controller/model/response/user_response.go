package response

type UserResponse struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	IsActive  bool   `json:"is_active"`
	IsOver18  bool   `json:"is_over_18"`
	Documents []struct {
		Type   string `json:"type"`
		Number string `json:"number"`
	} `json:"documents"`
}
