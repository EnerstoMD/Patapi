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

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PatientComment struct {
	Id        *string `json:"id" db:"id"`
	Comment   *string `json:"comment" db:"comment"`
	AddedBy   *string `json:"added_by" db:"added_by"`
	PatientId *string `json:"patient_id" db:"patient_id"`
}

type Disease struct {
	Id          *string `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}

type PatientDisease struct {
	Id         *string `json:"id" db:"id"`
	PatientId  *string `json:"patient_id" db:"patient_id"`
	Disease    Disease `json:"disease"`
	Comment    *string `json:"comment" db:"comment"`
	StartDate  *string `json:"start_date" db:"start_date"`
	EndDate    *string `json:"end_date" db:"end_date"`
	InProgress *string `json:"in_progress" db:"in_progress"`
	AddedBy    *string `json:"added_by" db:"added_by"`
}

type Allergy struct {
	Id          *string `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}

type PatientAllergy struct {
	Id         *string `json:"id" db:"id"`
	PatientId  *string `json:"patient_id" db:"patient_id"`
	Allergy    Allergy `json:"allergy"`
	Comment    *string `json:"comment" db:"comment"`
	StartDate  *string `json:"start_date" db:"start_date"`
	EndDate    *string `json:"end_date" db:"end_date"`
	InProgress *string `json:"in_progress" db:"in_progress"`
	AddedBy    *string `json:"added_by" db:"added_by"`
}

type Treatment struct {
	Id          *string `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	Comment     *string `json:"comment" db:"comment"`
}

type PatientTreatment struct {
	Id         *string   `json:"id" db:"id"`
	PatientId  *string   `json:"patient_id" db:"patient_id"`
	Treatment  Treatment `json:"treatment"`
	Comment    *string   `json:"comment" db:"comment"`
	StartDate  *string   `json:"start_date" db:"start_date"`
	EndDate    *string   `json:"end_date" db:"end_date"`
	InProgress *string   `json:"in_progress" db:"in_progress"`
	AddedBy    *string   `json:"added_by" db:"added_by"`
}

type Hospitalisation struct {
	Id        *string `json:"id" db:"id"`
	Motive    *string `json:"motive" db:"motive"`
	Comment   *string `json:"comment" db:"comment"`
	StartDate *string `json:"start_date" db:"start_date"`
	EndDate   *string `json:"end_date" db:"end_date"`
	PatientId *string `json:"patient_id" db:"patient_id"`
}

type PatientHistory struct {
	Id               *string `json:"id" db:"id"`
	PatientId        *string `json:"patient_id" db:"patient_id"`
	Disease          Disease `json:"disease"`
	Comment          *string `json:"comment" db:"comment"`
	FamilyConnection *string `json:"family_connection" db:"family_connection"`
}
