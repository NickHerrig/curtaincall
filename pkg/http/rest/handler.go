package rest

import (
    "encoding/json"
    "net/http"

    "curtaincall.tech/pkg/creating"
    "curtaincall.tech/pkg/retrieving"

    "github.com/bmizerany/pat"
)

func Handler(c creating.Service, r retrieving.Service) http.Handler {
    router := pat.New()

    router.Post("/theaters", http.HandlerFunc(createTheater(c)))
    router.Post("/shows", http.HandlerFunc(createShow(c)))

    router.Get("/theaters", http.HandlerFunc(retrieveAllTheaters(r)))

    return router
}

func createTheater(s creating.Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        //var id int
        decoder := json.NewDecoder(r.Body)

        var newTheater creating.Theater
        err := decoder.Decode(&newTheater)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        _, err = s.CreateTheater(newTheater)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        //SET HEADERS IN MIDDLEWARE
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode("New Theater Created!")
    }
}

func createShow(s creating.Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        //var id int
        decoder := json.NewDecoder(r.Body)

        var newShow creating.Show

        err := decoder.Decode(&newShow)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        _, err = s.CreateShow(newShow)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        //SET HEADERS IN MIDDLEWARE
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode("New Show Created")
    }
}

func retrieveAllTheaters(s retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

        theaters, err := s.RetrieveAllTheaters()
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        //SET HEADERS IN MIDDLEWARE
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(theaters)
    }
}
