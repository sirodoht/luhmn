CREATE TABLE documents (
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(300) NOT NULL,
    body TEXT
);
