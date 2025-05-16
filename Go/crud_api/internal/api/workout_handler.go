package api

import (
    "net/http"
    "fmt"

    "github.com/go-chi/chi/v5"
    "strconv"
)

type WorkoutHandler struct {}

func NewWorkoutHandler() *WorkoutHandler {
    return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
    paramsWorkoutID := chi.URLParam(r, "id")
    if paramsWorkoutID == "" {
        http.NotFound(w, r)
        return
    }

    workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "this is the workout id %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "created workout\n")
}


