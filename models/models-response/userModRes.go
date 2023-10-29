package modelsresponse

type UserCreateResponse struct {
	Email string `json:"email"`
	Password string `json:"password"`
    Token string `json:"token"`
}

type UserReponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Email string `json:"email"`
	Password string `json:"password"`
    Token string `json:"token"`
}