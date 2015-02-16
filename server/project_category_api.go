package server

import (
	"errors"
	"net/http"
)

type ProjectCategoryAPI struct {
	db *DB
}

func NewProjectCategoryAPI(db *DB) *ProjectCategoryAPI {
	return &ProjectCategoryAPI{db}
}

func (projectCategoryAPI *ProjectCategoryAPI) GetHandler(w http.ResponseWriter, r *http.Request, user *User) (interface{}, *HandlerError) {
	return nil, &HandlerError{errors.New("not implemented"), "not implemted", http.StatusNotImplemented}
}

func (projectCategoryAPI *ProjectCategoryAPI) GetListHandler(w http.ResponseWriter, r *http.Request, user *User) (interface{}, *HandlerError) {
	projectCategories, err := projectCategoryAPI.getList()
	if err != nil {
		return nil, &HandlerError{err, "couldn't retrieve project categories", http.StatusInternalServerError}
	}
	return projectCategories, nil
}

func (projectCategoryAPI *ProjectCategoryAPI) GetActivityViewListHandler(w http.ResponseWriter, r *http.Request, user *User) (interface{}, *HandlerError) {
	return nil, &HandlerError{errors.New("not implemented"), "not implemted", http.StatusNotImplemented}
}

func (projectCategoryAPI *ProjectCategoryAPI) SaveHandler(w http.ResponseWriter, r *http.Request, user *User) (interface{}, *HandlerError) {
	var projectCategory ProjectCategory
	err := unmarshalJSON(r.Body, &projectCategory)
	if err != nil {
		return nil, &HandlerError{err, err.Error(), http.StatusBadRequest}
	}

	err = projectCategoryAPI.save(&projectCategory)
	if err != nil {
		return nil, &HandlerError{err, "Error: Project category could not be saved.", http.StatusInternalServerError}
	}
	return jsonResultInt(projectCategory.ID)
}

func (projectCategoryAPI *ProjectCategoryAPI) DeleteHandler(w http.ResponseWriter, r *http.Request, user *User) (interface{}, *HandlerError) {
	return nil, &HandlerError{errors.New("not implemented"), "not implemted", http.StatusNotImplemented}
}

func (projectCategoryAPI *ProjectCategoryAPI) get(id int) (*ProjectCategory, error) {
	return nil, errors.New("not implemented")
}

func (projectCategoryAPI *ProjectCategoryAPI) getList() ([]ProjectCategory, error) {
	projectCategories, err := projectCategoryAPI.db.GetProjectCategories(nil)
	if err != nil {
		return nil, err
	}
	return projectCategories, nil
}

func (projectCategoryAPI *ProjectCategoryAPI) save(projectCategory *ProjectCategory) error {
	return projectCategoryAPI.db.SaveProjectCategory(projectCategory)
}

func (projectCategoryAPI *ProjectCategoryAPI) delete(id int) error {
	return errors.New("not implemented")
}
