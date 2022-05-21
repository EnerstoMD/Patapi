CREATE TABLE disease(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255)
);
CREATE TABLE patient_disease(
    id SERIAL PRIMARY KEY,
    patient_id INTEGER,
    disease_id INTEGER,
    start_date DATE,
    end_date DATE,
    comment VARCHAR(255),
    in_progress BOOLEAN
);
CREATE TABLE treatment(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255),
    comment VARCHAR(255)
);
CREATE TABLE patient_treatment(
    id SERIAL PRIMARY KEY,
    patient_id INTEGER,
    treatment_id INTEGER,
    start_date DATE,
    end_date DATE,
    comment VARCHAR(255),
    in_progress BOOLEAN
);
CREATE TABLE allergy(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255)
);
CREATE TABLE patient_allergy(
    id SERIAL PRIMARY KEY,
    patient_id INTEGER,
    allergy_id INTEGER,
    start_date DATE,
    end_date DATE,
    comment VARCHAR(255),
    in_progress BOOLEAN
);
CREATE TABLE patient_comment(
    id SERIAL PRIMARY KEY,
    patient_id INTEGER,
    comment VARCHAR(255)
);