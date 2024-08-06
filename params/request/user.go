package request

//Registration
// {
//   "user":{
//     "username": "Jacob",
//     "email": "jake@jake.jake",
//     "password": "jakejake"
//   }
// }

type UserRegistrationRequest struct {
	User UserRegistrationBody `json:"user"`
}

type UserRegistrationBody struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//login
// {
//   "user":{
//     "email": "jake@jake.jake",
//     "password": "jakejake"
//   }
// }

type UserLoginRequest struct {
	User UserLoginBody `json:"user"`
}

type UserLoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
