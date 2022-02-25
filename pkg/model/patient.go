package model

type Patient struct {
	Id             *string `json:"id" db:"id"`
	Oid            *string `json:"oid" db:"oid"`
	Name           *string `json:"name" db:"name"`
	Firstnames     *string `json:"firstnames" db:"firstnames"`
	Lastname       *string `json:"lastname" db:"lastname"`
	Birthname      *string `json:"birthname" db:"birthname"`
	Gender         *string `json:"gender" db:"gender"`
	Birthdate      *string `json:"birthdate" db:"birthdate"`
	BirthplaceCode *string `json:"birthplace_code" db:"birthplace_code"`
	InsMatricule   *string `json:"ins_matricule" db:"ins_matricule"`
	Nir            *string `json:"nir" db:"nir"`
	Nia            *string `json:"nia" db:"nia"`
	Address        *string `json:"address" db:"address"`
	City           *string `json:"city" db:"city"`
	Postalcode     *string `json:"postalcode" db:"postalcode"`
	Phone          *string `json:"phone" db:"phone"`
	Email          *string `json:"email" db:"email"`
}