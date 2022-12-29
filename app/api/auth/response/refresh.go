package response

type RefreshResponse struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenMaxAge uint64 `json:"accessTokenMaxAge"`
}
