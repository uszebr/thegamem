package entity

type UserAuth struct {
	Email      string
	Role       string
	IsAnonymus bool
	UserId     string
}

type AuthDetails struct {
	AccessToken  string
	ExpiresIn    int
	RefreshToken string
}
