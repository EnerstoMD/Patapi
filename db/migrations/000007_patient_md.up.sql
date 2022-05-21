CREATE TABLE md (
  id serial PRIMARY KEY,
  name VARCHAR(255),
  lastname VARCHAR(255) ,
  speciality VARCHAR(255),
  address VARCHAR(255),
  phone VARCHAR(255),
  email VARCHAR(255)
);

CREATE TABLE patient_md (
    id serial PRIMARY KEY,
    patient_id integer NOT NULL,
    md_id integer,
    added_by integer
);
ALTER TABLE patient_md
    ADD CONSTRAINT "FK_patient_md_patient" FOREIGN KEY (patient_id) REFERENCES "patient" (id) ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE patient_md
    ADD CONSTRAINT "FK_patient_md_md" FOREIGN KEY (md_id) REFERENCES "md" (id) ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE patient_md
    ADD CONSTRAINT "FK_patient_md_user" FOREIGN KEY (added_by) REFERENCES "user" (id) ON UPDATE NO ACTION ON DELETE NO ACTION;