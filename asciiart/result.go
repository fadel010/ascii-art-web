package asciiart

import (
	"html/template"
	"net/http"
)

type Response struct {
	Output string
}

func Result(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		// Parse form data for POST requests
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		userInput := r.FormValue("string")
		banner := r.FormValue("banner")

		// Handle empty input or banner
		// if userInput == "" || banner == "" {
		// 	http.Error(w, "Both text and banner are required.", http.StatusBadRequest)
		// 	return
		// }

		asciiartResult := Execute(userInput, banner)

		// Display the result in the home page (option 2 from the prompt)
		str := Response{
			Output: asciiartResult,
		}

		// Execute HTML template and write to the response writer
		err = tmpl.Execute(w, str)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	}
}
