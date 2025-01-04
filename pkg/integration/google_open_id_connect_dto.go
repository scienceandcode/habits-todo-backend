package integration

type GoogleOpenIdConnectConfigurationDTO struct {
	UserInfoEndpoint string `json:"userinfo_endpoint"`
}

func NewGoogleOpenIdConnectConfigurationDTO() *GoogleOpenIdConnectConfigurationDTO {
	return &GoogleOpenIdConnectConfigurationDTO{}
}
