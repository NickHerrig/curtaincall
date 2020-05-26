package rest

import (
    "encoding/json"
    "net/http"
    "strconv"

    "curtaincall.tech/pkg/adding"

    "github.com/bmizerany/pat"
)

func Handler(a adding.Service) http.Handler {
    router := pat.New()

    router.Post("/theaters", http.HandlerFunc(addTheater(a)))
    router.Post("/theaters/:id/shows", http.HandlerFunc(addShow(a)))

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
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode("New Theater Added")
    }
}

func addShow(s adding.Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
	    id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	    if err != nil || id < 1 {
            http.Error(w, err.Error(), http.StatusBadRequest)
	        return
	    }

        decoder := json.NewDecoder(r.Body)

        var newShow adding.Show
        newShow.TheaterID = id

        err = decoder.Decode(&newShow)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        err = s.AddShow(newShow)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode("New Show Added")
    }
}
