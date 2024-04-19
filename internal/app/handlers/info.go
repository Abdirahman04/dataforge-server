package handlers

import (
	"encoding/json"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)

  myData := map[string]interface{}{
    "project": "dataforge",
    "name": "Abdirahman Hassan",
    "age": 21,
    "email": "abdixcom@gmai.com",
    "github": "github.com/Abdirahman04",
    "linkedin": "linkedin.com/in/abdirahman-hassan-a12ab42a6",
  }

  myJsonData, err := json.Marshal(myData)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  w.Write(myJsonData)
}
