package models

type Secret struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
}
