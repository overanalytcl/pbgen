package pbgen

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProblemDetails struct {
	ID             int    `json:"id"`
	ProgLang       string `json:"limbaj_de_programare"`
	Name           string `json:"denumire"`
	Difficulty     int    `json:"dificultate"`
	Grade          int    `json:"clasa"`
	TimeLimit      string `json:"limita_timp"`
	MemoryLimit    string `json:"limita_memorie"`
	StackLimit     string `json:"limita_stiva"`
	UseConsole     string `json:"folosesc_consola"`
	InFile         string `json:"fisier_in"`
	OutFile        string `json:"fisier_out"`
	OkFile         string `json:"fisier_ok"`
	IDUser         int    `json:"id_user"`
	Visible        int    `json:"vizibila"`
	Approved       int    `json:"aprobata"`
	Author         string `json:"autor"`
	ProblemSource  string `json:"sursa_problema"`
	ContestID      int    `json:"id_concurs"`
	ContestLevelID int    `json:"id_nivel_concurs"`
	Statement      string `json:"enunt_html"`
	Summary        string `json:"rezumat_html"`
	Solution       string `json:"solutie_html"`
	Etichete       []any  `json:"etichete"`
	Tags           []struct {
		Tag struct {
			ID       string `json:"id"`
			Denumire string `json:"denumire"`
			Clasa    string `json:"clasa"`
		} `json:"tag"`
		Subtag struct {
			ID       string `json:"id"`
			Denumire string `json:"denumire"`
			Clasa    string `json:"clasa"`
		} `json:"subtag"`
	} `json:"taguri"`
	User struct {
		User    string `json:"user"`
		Nume    string `json:"nume"`
		Prenume string `json:"prenume"`
	} `json:"user"`
	SourceName string `json:"nume_sursa"`
}

type apiResponse struct {
	Status   string          `json:"stare"`
	Response string          `json:"raspuns"`
	IsAuthed bool            `json:"user_autentificat"`
	Problem  json.RawMessage `json:"problema"`
}

func NewProblemDetails(id int) (*ProblemDetails, error) {
	url := fmt.Sprintf("https://new.pbinfo.ro/json/probleme/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	if string(data.Problem) == "false" {
		return nil, fmt.Errorf("problem doesn't exist")
	}

	var problem ProblemDetails
	if err := json.Unmarshal(data.Problem, &problem); err != nil {
		return nil, fmt.Errorf("failed to unmarshal problem details: %w", err)
	}

	return &problem, nil
}
