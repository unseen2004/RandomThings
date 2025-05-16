package app

import(
    "log"
    "os"
    "fmt"
    "net/http"

    "github.com/unseen2004/crudApi/internal/api"
)

type Application struct {
    Logger          *log.Logger
    WorkoutHandler  *api.WorkoutHandler
}

func NewApplication() (*Application, error){
    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

    
    workoutHandler := api.NewWorkoutHandler()

    app := &Application{
        Logger: logger,
        WorkoutHandler: workoutHandler,
    }

    return app, nil
}


func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Status is available\n")
}
