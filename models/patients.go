package models

type Patient struct {
	Id             *string `json:"id" sql:"id"`
	Oid            *string `json:"oid" sql:"oid"`
	Name           *string `json:"name" sql:"name"`
	Firstnames     *string `json:"firstnames" sql:"firstnames"`
	Lastname       *string `json:"lastname" sql:"lastname"`
	Birthname      *string `json:"birthname" sql:"birthname"`
	Gender         *string `json:"gender" sql:"gender"`
	Birthdate      *string `json:"birthdate" sql:"birthdate"`
	BirthplaceCode *string `json:"birthplace_code" sql:"birthplace_code"`
	InsMatricule   *string `json:"ins_matricule" sql:"ins_matricule"`
	Nir            *string `json:"nir" sql:"nir"`
	Nia            *string `json:"nia" sql:"nia"`
	Address        *string `json:"address" sql:"address"`
	City           *string `json:"city" sql:"city"`
	Postalcode     *string `json:"postalcode" sql:"postalcode"`
	Phone          *string `json:"phone" sql:"phone"`
	Email          *string `json:"email" sql:"email"`
}
