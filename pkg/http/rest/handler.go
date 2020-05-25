package rest

import (
    "encoding/json"
    "net/http"

    "curtaincall.tech/pkg/adding"

    "github.com/bmizerany/pat"
)

func Handler(a adding.Service) http.Handler {
    router := pat.New()
    router.Post("/theaters", http.HandlerFunc(addTheater(a)))
    return router
}

func addTheater(s adding.Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        decoder := json.NewDecoder(r.Body)

        var newTheater adding.Theater
        err := decoder.Decode(&newTheater)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        err = s.AddTheater(newTheater)
        //TODO Error handling

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode("New Theater Added %d")
    }
}
