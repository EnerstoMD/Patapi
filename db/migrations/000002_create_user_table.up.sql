BEGIN;
CREATE TABLE public.user (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
	email VARCHAR(255),
	password VARCHAR(255)
);
ALTER TABLE public.user ADD CONSTRAINT "identifier" UNIQUE (email);

COMMIT;