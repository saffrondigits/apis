package main

type AuthService interface {
	Login() (string, error)
	Register() (string, error)
	ForgotPassword() (string, error)
	ResetPassword() (string, error)
	Logout(string) (string, error)
	EmailVerification() (string, error)
}

type SystemUser struct {
	Username      string
	Email         string
	Password      string
	AssignedToken string
}

func (su *SystemUser) Login() (string, error) {
	panic("Implement me")
}

func (su *SystemUser) Register() (string, error) {
	panic("Implement me")
}

func (su *SystemUser) ForgotPassword() (string, error) {
	panic("Implement me")
}

func (su *SystemUser) ResetPassword() (string, error) {
	panic("Implement me")
}

func (su *SystemUser) Logout(s string) (string, error) {
	panic("Implement me")
}

func (su *SystemUser) EmailVerification() (string, error) {
	panic("Implement me")
}

func main() {
	su := &SystemUser{
		Username:      "rakshits",
		Password:      "Secure",
		AssignedToken: "sdncvjkfsdbv",
	}

	GetAuthDetails(su)

}

func GetAuthDetails(auth AuthService) {

}
