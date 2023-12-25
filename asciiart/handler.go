package asciiart

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Response struct {
	Output string
}

const indexTemplatePath  = "templates/index.html"

var asciiArtText string


func handleError(w http.ResponseWriter, statusCode int, message string) {
	http.Error(w, fmt.Sprintf("%d %s: %s", statusCode, http.StatusText(statusCode), message), statusCode)
}

func loadTemplate(filename string) (*template.Template, error) {
    return template.ParseFiles(filename)
}


func RenderErrorPage(w http.ResponseWriter, statusCode int, message string) {
	tmpl, err := loadTemplate(fmt.Sprintf("templates/%d.html", statusCode))
	if err != nil {
		handleError(w, statusCode, message)
		return
	}

	w.WriteHeader(statusCode)
	err = tmpl.Execute(w, nil)
	if err != nil {
		handleError(w, statusCode, message)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderErrorPage(w, http.StatusNotFound, "Page not found")
		return
	}

	tmpl, err := loadTemplate(indexTemplatePath)
	if err != nil {
		RenderErrorPage(w, http.StatusInternalServerError, "Failed to load template")
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		RenderErrorPage(w, http.StatusInternalServerError, "Failed to render template")
		return
	}
}

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderErrorPage(w, http.StatusBadRequest, "Invalid request method")
		return
	}

	tmpl, err := loadTemplate(indexTemplatePath)
	if err != nil {
		RenderErrorPage(w, http.StatusInternalServerError, "Failed to load template")
		return
	}

	err = r.ParseForm()
	if err != nil {
		RenderErrorPage(w, http.StatusInternalServerError, "Failed to parse form data")
		return
	}

	userInput := r.FormValue("string")
	banner := r.FormValue("banner")

	if len(banner) == 0 {
		RenderErrorPage(w, http.StatusBadRequest, "Banner cannot be empty")
		return
	}

	asciiartResult, statusCode := Execute(userInput, banner)

	asciiArtText = asciiartResult

	str := Response{
		Output: asciiartResult,
	}

	if statusCode != http.StatusOK {
		RenderErrorPage(w, statusCode, "Failed to generate ASCII art")
		return
	}

	err = tmpl.Execute(w, str)
	if err != nil {
		RenderErrorPage(w, http.StatusInternalServerError, "Failed to render template")
		return
	}
}

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for file download
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(asciiArtText)))
	w.Header().Set("Content-Disposition", "attachment; filename=asciiart.txt")
	w.Write([]byte(asciiArtText))
}
