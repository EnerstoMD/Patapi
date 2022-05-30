CREATE TABLE hospitalisation (
  id serial PRIMARY KEY,
    patient_id integer,
  motive VARCHAR(255),
  start_date DATE,
    end_date DATE,
  comment VARCHAR(255)
);

CREATE TABLE patient_history (
  id serial PRIMARY KEY,
  disease_id integer,
    patient_id integer,
    family_connection VARCHAR(255),
  comment VARCHAR(255)
);