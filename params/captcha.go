package params

type VerifyCaptchaRequest struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
	RemoteIP string `json:"remoteip"`
	SiteKey  string `json:"sitekey"`
}

type VerifyCaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}
