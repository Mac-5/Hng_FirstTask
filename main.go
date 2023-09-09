package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


type Response struct{
	Slack_name string `json:"slack_name"`
	Current_day string `json:"current_day"`
	Utc_time string 	`utc_time:"utc_time"`
	Track string 	`json:"track"`
	Github_file_url string `json:"github_file_url"`
	Github_repo_url string 	`json:"github_repo_url"`
	Status_code int 	`json:"status_code"`

}

func currentDay()(string){
	currentTime := time.Now()

	current_day := currentTime.Weekday()

	day_of_the_week := current_day.String()

	return day_of_the_week


}
func vaildUtcTime()(string){
	currentTime := time.Now().UTC()

	window  := 2 * time.Minute
	lowMark := currentTime.Add(-window)
	upMark := currentTime.Add(window)

	if (currentTime.After(lowMark) && currentTime.Before(upMark)){
		formattedTime := currentTime.Format("2006-01-02T15:04-07:00Z")
		return formattedTime
	}else{
		return ""
	}
}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		 slack_name := r.URL.Query().Get("slack_name")
		 track := r.URL.Query().Get("track")

		data := Response{
			Slack_name: slack_name,
			Current_day: currentDay(),
			Utc_time: vaildUtcTime(),
			Track: track,
			Github_file_url: "",
			Github_repo_url: "",
			Status_code: http.StatusOK,

		}
		
		// param2 := r.URL.Query().Get("param2")

		json, err := json.Marshal(data)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type","application/json")

		w.Write(json)

		
      
    })
    http.ListenAndServe(":3000", r)
}