package params

type VerifyCaptchaRequest struct {
	Secret   string `url:"secret"`
	Response string `url:"response"`
	RemoteIP string `url:"remoteip"`
	SiteKey  string `url:"sitekey"`
}

type VerifyCaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}
