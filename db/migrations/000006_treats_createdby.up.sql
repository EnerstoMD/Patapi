ALTER TABLE patient_disease ADD COLUMN added_by INTEGER;
ALTER TABLE patient_treatment ADD COLUMN added_by INTEGER;
ALTER TABLE patient_allergy ADD COLUMN added_by INTEGER;
ALTER TABLE patient_comment ADD COLUMN added_by INTEGER;

ALTER TABLE public.patient_treatment
	ADD CONSTRAINT "FK_add_treatment_user" FOREIGN KEY (added_by) REFERENCES "user" (id) ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE public.patient_allergy
	ADD CONSTRAINT "FK_add_allergy_user" FOREIGN KEY (added_by) REFERENCES "user" (id) ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE public.patient_disease
	ADD CONSTRAINT "FK_add_disease_user" FOREIGN KEY (added_by) REFERENCES "user" (id) ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE public.patient_comment
	ADD CONSTRAINT "FK_add_patientcomment_user" FOREIGN KEY (added_by) REFERENCES "user" (id) ON UPDATE NO ACTION ON DELETE NO ACTION;