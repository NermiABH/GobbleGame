package vars

type RequestUserRegister struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestJwtLogin struct {
	Username string `json:"username" validate:"required_if=Email ''"`
	Email    string `json:"email" validate:"required_if=Username ''"`
	Password string `json:"password" validate:"required"`
}

type RequestJwtRefresh struct {
	Refresh string `json:"refresh" validate:"required"`
}

type RequestUserUniqueUsername struct {
	Username string `json:"username" validate:"required"`
}

type RequestUserUniqueEmail struct {
	Email string `json:"email" validate:"required"`
}
