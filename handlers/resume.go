package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"log"
	"net/http"
	"os"

	"github.com/RichDom2185/2023-website-backend/params"
	"github.com/google/go-querystring/query"
)

const (
	// Environment variable keys
	RESUME_PATH_KEY      = "RESUME_PATH"
	HCAPTCHA_SITEKEY_KEY = "HCAPTCHA_SITEKEY"
	HCAPTCHA_SECRET_KEY  = "HCAPTCHA_SECRET"

	// See https://docs.hcaptcha.com/
	HCAPTCHA_VERIFICATION_URL          = "https://hcaptcha.com/siteverify"
	HCAPTCHA_VERIFICATION_CONTENT_TYPE = "application/x-www-form-urlencoded"
)

func HandleResumeForm(w http.ResponseWriter, r *http.Request) {
	var requestParams params.ResumePostRequest
	err := json.NewDecoder(r.Body).Decode(&requestParams)
	if err != nil {
		log.Fatalln(err)
	}

	var captchaParams params.VerifyCaptchaRequest = params.VerifyCaptchaRequest{
		Secret:   os.Getenv(HCAPTCHA_SECRET_KEY),
		Response: requestParams.Token,
		RemoteIP: r.RemoteAddr,
		SiteKey:  os.Getenv(HCAPTCHA_SITEKEY_KEY),
	}
	values, err := query.Values(captchaParams)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(HCAPTCHA_VERIFICATION_URL, HCAPTCHA_VERIFICATION_CONTENT_TYPE, strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	var captchaResponse params.VerifyCaptchaResponse
	err = json.NewDecoder(resp.Body).Decode(&captchaResponse)
	if err != nil {
		log.Fatalln(err)
	}

	if !captchaResponse.Success {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Captcha verification failed: %s", captchaResponse.ErrorCodes)
		return
	}

	resumeFile, err := os.ReadFile(os.Getenv(RESUME_PATH_KEY))
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Add("Content-Type", "application/pdf")
	w.Write(resumeFile)
}
