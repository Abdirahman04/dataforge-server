package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Abdirahman04/dataforge-server/internal/app/models"
)

func DataCraftHandler(w http.ResponseWriter, r *http.Request) {
  var DataSchema models.Blueprint

  if err := json.NewDecoder(r.Body).Decode(&DataSchema);err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  res, err := json.Marshal(DataSchema)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write(res)
}
