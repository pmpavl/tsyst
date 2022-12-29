package response

type AuthenticationResponse struct {
	AccessToken        string `json:"accessToken"`
	AccessTokenMaxAge  uint64 `json:"accessTokenMaxAge"`
	RefreshToken       string `json:"refreshToken"`
	RefreshTokenMaxAge uint64 `json:"refreshTokenMaxAge"`
}
