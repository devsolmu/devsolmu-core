package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllProjects is getting all projects
func (app *App) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := []Project{}
	app.DB.Find(&projects)
	respondJSON(w, http.StatusOK, projects)
}

// GetProject is getting a project
func (app *App) GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(app.DB, title, w, r)
	if project == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// CreateProject is creating a project
func (app *App) CreateProject(w http.ResponseWriter, r *http.Request) {
	project := Project{}

	decoder := json.NewDecoder(r.Body)
	if error := decoder.Decode(&project); error != nil {
		respondError(w, http.StatusBadRequest, error.Error())
		return
	}
	defer r.Body.Close()

	if error := app.DB.Save(&project).Error; error != nil {
		respondError(w, http.StatusInternalServerError, error.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

// UpdateProject is updating a project
func (app *App) UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(app.DB, title, w, r)
	if project == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if error := decoder.Decode(&project); error != nil {
		respondError(w, http.StatusBadRequest, error.Error())
		return
	}
	defer r.Body.Close()

	if error := app.DB.Save(&project).Error; error != nil {
		respondError(w, http.StatusInternalServerError, error.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// DeleteProject is deleting a project
func (app *App) DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(app.DB, title, w, r)
	if project == nil {
		return
	}

	if error := app.DB.Delete(&project).Error; error != nil {
		respondError(w, http.StatusInternalServerError, error.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// GetProjectOr404 is getting a project or output 404 error
func GetProjectOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *Project {
	project := Project{}
	if error := db.First(&project, Project{Title: title}).Error; error != nil {
		respondError(w, http.StatusNotFound, error.Error())
		return nil
	}
	return &project
}
