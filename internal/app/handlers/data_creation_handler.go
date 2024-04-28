package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Abdirahman04/dataforge-server/internal/app/models"
	"github.com/Abdirahman04/dataforge-server/internal/app/services"
)

func DataCraftHandler(w http.ResponseWriter, r *http.Request) {
  var DataSchema models.Blueprint

  if err := json.NewDecoder(r.Body).Decode(&DataSchema);err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }
  
  list := services.ListForge(DataSchema)

  res, err := json.Marshal(list)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write(res)
}
