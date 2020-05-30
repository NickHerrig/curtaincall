package rest

import (
    "encoding/json"
    "errors"
    "net/http"
    "strconv"

    "curtaincall.tech/pkg/creating"
    "curtaincall.tech/pkg/retrieving"
    "curtaincall.tech/pkg/deleting"
    "curtaincall.tech/pkg/storage/sqlite"

    "github.com/bmizerany/pat"
)

func Handler(c creating.Service, r retrieving.Service, d deleting.Service) http.Handler {
    router := pat.New()

    router.Post("/theaters", http.HandlerFunc(createTheater(c)))
    router.Get("/theaters", http.HandlerFunc(retrieveAllTheaters(r)))
    router.Get("/theaters/:id", http.HandlerFunc(retrieveTheater(r)))
    router.Del("/theaters/:id", http.HandlerFunc(deleteTheater(d)))

    router.Get("/theaters/:id/shows", http.HandlerFunc(retrieveAllShows(r)))
    // router.Post("/shows", http.HandlerFunc(createShow(c)))

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

func retrieveTheater(s retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

      id, err := strconv.Atoi(r.URL.Query().Get(":id"))
      if err != nil {
          http.Error(w, err.Error(), http.StatusNotFound)
          return
      }

      theater, err := s.RetrieveTheater(id)
      if err != nil {
          if errors.Is(err, sqlite.ErrNoRecord) {
              http.Error(w, err.Error(), http.StatusNotFound)
          } else {
              http.Error(w, err.Error(), http.StatusBadRequest)
          }
          return
      }
        
      //SET HEADERS IN MIDDLEWARE
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(theater)
    }
}

func deleteTheater(s deleting.Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

      id, err := strconv.Atoi(r.URL.Query().Get(":id"))
      if err != nil {
          http.Error(w, err.Error(), http.StatusNotFound)
          return
      }

      err = s.DeleteTheater(id)
      if err != nil {
          if errors.Is(err, sqlite.ErrNoRecord) {
              http.Error(w, err.Error(), http.StatusNotFound)
          } else {
              http.Error(w, err.Error(), http.StatusBadRequest)
          }
          return
      }
        
      //SET HEADERS IN MIDDLEWARE
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode("Theater Deleted!")
    }
}

func retrieveAllShows(s retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
      id, err := strconv.Atoi(r.URL.Query().Get(":id"))
      if err != nil {
          http.Error(w, err.Error(), http.StatusNotFound)
          return
      }

      shows, err := s.RetrieveAllShows(id)
      if err != nil {
          if errors.Is(err, sqlite.ErrNoRecord) {
              http.Error(w, err.Error(), http.StatusNotFound)
          } else {
              http.Error(w, err.Error(), http.StatusBadRequest)
          }
          return
      }
        
      //SET HEADERS IN MIDDLEWARE
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(shows)
    }
}
