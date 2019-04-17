package handler

import (
	"encoding/json"
	"net/http"

	"github.com/devsolmu/devsolmu-core/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllProjects is getting all projects
func GetAllProjects(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	projects := []model.Project{}
	db.Find(&projects)
	respondJSON(w, http.StatusOK, projects)
}

// GetProject is getting a project
func GetProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(db, title, w, r)
	if project == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// CreateProject is creating a project
func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	project := model.Project{}

	decoder := json.NewDecoder(r.Body)
	if error := decoder.Decode(&project); error != nil {
		respondError(w, http.StatusBadRequest, error.Error())
		return
	}
	defer r.Body.Close()

	if error := db.Save(&project).Error; error != nil {
		respondError(w, http.StatusInternalServerError, error.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

// UpdateProject is updating a project
func UpdateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(db, title, w, r)
	if project == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if error := decoder.Decode(&project); error != nil {
		respondError(w, http.StatusBadRequest, error.Error())
		return
	}
	defer r.Body.Close()

	if error := db.Save(&project).Error; error != nil {
		respondError(w, http.StatusInternalServerError, error.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// DeleteProject is deleting a project
func DeleteProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(db, title, w, r)
	if project == nil {
		return
	}

	if error := db.Delete(&project).Error; error != nil {
		respondError(w, http.StatusInternalServerError, error.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// GetProjectOr404 is getting a project or output 404 error
func GetProjectOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *model.Project {
	project := model.Project{}
	if error := db.First(&project, model.Project{Title: title}).Error; error != nil {
		respondError(w, http.StatusNotFound, error.Error())
		return nil
	}
	return &project
}
