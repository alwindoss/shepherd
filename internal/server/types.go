package server

import "net/http"

type ResourceHandler interface {
	// Index handler lists all the resources
	Index(w http.ResponseWriter, r *http.Request)

	// New Handler returns a form that is used to add new resource
	New(w http.ResponseWriter, r *http.Request)

	// Create handler is called when the new resource form is submitted to save the resource
	Create(w http.ResponseWriter, r *http.Request)

	// Show handler returns the resource details whose id passed as path parameter
	Show(w http.ResponseWriter, r *http.Request)

	// Edit Handler returns a form that is used to edit the resource with id
	Edit(w http.ResponseWriter, r *http.Request)

	// UpdateFull handler is called to update the resource fully
	UpdateFull(w http.ResponseWriter, r *http.Request)

	// UpdatePartial handler is called to update the resource partially
	UpdatePartial(w http.ResponseWriter, r *http.Request)

	// Delete handler is called with delete the resource with id
	Delete(w http.ResponseWriter, r *http.Request)
}
