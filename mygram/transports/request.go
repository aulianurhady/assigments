package transports

type RequestUser struct {
	Age      int    `json:"age,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}

type RequestPhoto struct {
	Title    string `json:"title,omitempty"`
	Caption  string `json:"caption,omitempty"`
	PhotoURL string `json:"photo_url,omitempty"`
}
