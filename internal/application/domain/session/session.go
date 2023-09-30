package dmsession

type Session struct {
	Id                    string
	UserId                string
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  string
	RefreshTokenExpiresAt string
}
