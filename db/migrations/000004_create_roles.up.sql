BEGIN;
CREATE TABLE role(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255)
);

CREATE TABLE user_roles(
    user_id INTEGER,
    role_id INTEGER
);
COMMIT;