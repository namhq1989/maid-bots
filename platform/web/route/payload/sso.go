package routepayload

type SSOLoginWithGoogleBody struct {
	Token string `json:"token" validate:"required" message:"Google token is required"`
}

type SSOLoginWithGitHubBody struct {
	Code string `json:"code" validate:"required" message:"GitHub authorization code is required"`
}
