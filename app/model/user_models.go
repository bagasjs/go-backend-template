package model

type CreateUpdateUserRequest struct {
    Name string `json:"name" form:"name"`
    Email string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
    PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation"`
}

type GeneralUserResponse struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    Created string `json:"created"`
    Updated string `json:"updated"`
}

type LoginUserRequest struct {
    Email string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
}
