package handlers

import (
	"log"
	"net/http"
	"os"
)

const (
	RESUME_PATH_KEY = "RESUME_PATH"
)

func HandleResumeForm(w http.ResponseWriter, r *http.Request) {
	resumeFile, err := os.ReadFile(os.Getenv(RESUME_PATH_KEY))
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "application/pdf")
	w.Write(resumeFile)
}
