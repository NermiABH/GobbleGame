package vars

type Response struct {
	Type          string `json:"type"`
	Id            int    `json:"id,omitempty"`
	Attributes    any    `json:"attributes"`
	Relationships any    `json:"relationships,omitempty"`
}

type UserAttributes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
