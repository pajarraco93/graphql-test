package model

type Song struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AppearsIn int    `json:"appersIn"`
}
