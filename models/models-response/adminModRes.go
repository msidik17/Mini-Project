package modelsresponse

type AdminCreateResponse struct {
	Email string `json:"email"`
	Password string `json:"password"`
    Token string `json:"token"`
}

type AdminReponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	Email string `json:"email"`
	Password string `json:"password"`
    Token string `json:"token"`
}