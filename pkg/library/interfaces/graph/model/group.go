package model

type Group struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Genre     *string `json:"genre"`
	GroupInfo *string `json:"groupInfo"`
}
