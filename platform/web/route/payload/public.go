package routepayload

type PublicCheckDomain struct {
	Domain string `query:"domain" validate:"required" message:"Domain name is required"`
}
