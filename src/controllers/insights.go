package controllers

import (
	"encoding/json"
	"net/http"

	"aurora-insights-api/src/models"
)

func GetInsights(w http.ResponseWriter, r *http.Request) {
	insights := []models.Insight{ { Id: "1", Description: "Total Users: 1000" }, { Id: "2", Description: "Total Sessions: 5000" } }
	json.NewEncoder(w).Encode(insights)
}