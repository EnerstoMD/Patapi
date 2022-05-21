ALTER TABLE patient_disease
	DROP FOREIGN KEY "FK_add_disease_user";
ALTER TABLE patient_treatment
    DROP FOREIGN KEY "FK_add_treatment_user";
ALTER TABLE patient_allergy
    DROP FOREIGN KEY "FK_add_allergy_user";
ALTER TABLE patient_comment
    DROP FOREIGN KEY "FK_add_patientcomment_user";

ALTER TABLE patient_disease
    DROP COLUMN added_by;
ALTER TABLE patient_treatment
    DROP COLUMN added_by;
ALTER TABLE patient_allergy
    DROP COLUMN added_by;
ALTER TABLE patient_comment
    DROP COLUMN added_by;
    