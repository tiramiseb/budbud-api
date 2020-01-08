package model

// User represents a user
type User struct {
	ID    string
	Email string
}

// Workspace represents a workspace
type Workspace struct {
	ID    string
	Name  string
	Owner User
}

// SuperCategory represents a supercategory
type SuperCategory struct {
	ID         string
	Name       string
	Categories []Category
}

// Category represents a category
type Category struct {
	ID   string
	Name string
}
