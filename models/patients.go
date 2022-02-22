package models

type Patient struct {
	Id             string `json:"id"`
	Oid            string `json:"oid"`
	Name           string `json:"name"`
	Firstnames     string `json:"firstnames"`
	Lastname       string `json:"lastname"`
	Birthname      string `json:"birthname"`
	Gender         string `json:"gender"`
	Birthdate      string `json:"birthdate"`
	BirthplaceCode string `json:"birthplace_code"`
	InsMatricule   string `json:"ins_matricule"`
	Nir            string `json:"nir"`
	Nia            string `json:"nia"`
	Address        string `json:"address"`
	City           string `json:"city"`
	Postalcode     string `json:"postalcode"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}
