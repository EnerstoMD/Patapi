BEGIN;
CREATE TABLE public.event (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description VARCHAR(255),
    startdate VARCHAR(255),
    enddate VARCHAR(255),
    is_confirmed VARCHAR(5),
	created_by INTEGER
);
ALTER TABLE public.event
	ADD CONSTRAINT "FK_event_user" FOREIGN KEY (created_by) REFERENCES "user" (id) ON UPDATE NO ACTION ON DELETE NO ACTION;

COMMIT;