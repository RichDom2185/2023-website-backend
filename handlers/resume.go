package handlers

import (
	"log"
	"net/http"
	"os"
)

const (
	RESUME_PATH = "./resume-latest.pdf"
)

func HandleResumeForm(w http.ResponseWriter, r *http.Request) {
	resumeFile, err := os.ReadFile(RESUME_PATH)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "application/pdf")
	w.Write(resumeFile)
}
