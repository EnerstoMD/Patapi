CREATE DATABASE lupusdb;

CREATE TABLE patient (
    id SERIAL PRIMARY KEY,
    oid VARCHAR(255), 
	name VARCHAR(255),      
	firstnames VARCHAR(255),     
	lastname VARCHAR(255),     
	birthname VARCHAR(255),     
	gender VARCHAR(255),        
	birthdate VARCHAR(255),     
	birthplace_code VARCHAR(255),
	ins_matricule VARCHAR(255),  
	nir VARCHAR(255),           
	nia VARCHAR(255),           
	address VARCHAR(255),       
	city VARCHAR(255),          
	postalcode VARCHAR(255),    
	phone VARCHAR(255),         
	email VARCHAR(255)         
);

CREATE TABLE user (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
	email VARCHAR(255),
	password VARCHAR(255)
);
ALTER TABLE public.user ADD CONSTRAINT "identifier" UNIQUE (email);

CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description VARCHAR(255),
    startdate VARCHAR(255),
    enddate VARCHAR(255),
    is_confirmed VARCHAR(5),
	created_by INTEGER,
);
ALTER TABLE "event"
	ADD CONSTRAINT "FK_event_user" FOREIGN KEY ("created_by") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

ALTER TABLE public.patient ADD CONSTRAINT "patientIdentifier" UNIQUE (name,lastname,ins_matricule);