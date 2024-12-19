package application

import (
	"CALC_GO/pkg/calculation"
	"encoding/json"
	"net/http"
	"os"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}


type Result struct {
	Result float64 `json:"result"`
}
type Error struct {
	Message string `json:"error"`
}
func CalcHandler(w http.ResponseWriter, r *http.Request) {
    request := new(Request)
    defer r.Body.Close()
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Ошибки связанные с некоректностью выражения
    errs := []error{
		calculation.ErrEmptyExpression,
		calculation.ErrIncorrectBracketPlacement,
		calculation.ErrBracketMismatch,
		calculation.ErrInvalidCharacter,
		calculation.ErrNotEnoughOperands,
		calculation.ErrIncorrectExpression,
		calculation.ErrDivisionByZero,
    }
    
    res, err := calculation.Calc(request.Expression)
    if err != nil {
        for _, ERROR := range errs{
            if err == ERROR{
                result := Error{Message: "Expression is not valid"}
                w.WriteHeader(http.StatusUnprocessableEntity)
                json.NewEncoder(w).Encode(result)
                return 
            }
        }
        result := Error{Message: "Internal server error"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(result)
    } else {
        result := Result{Result: res}
        json.NewEncoder(w).Encode(result)
    }
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}