-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    description varchar(255),
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL
);

INSERT INTO roles (name, description, "created_at", "updated_at") VALUES 
('sadmin', 'Super Administrador', now(), now()),
('admin', 'Administrador', now(), now()),
('user', 'Usuário', now(), now());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE roles;
-- +goose StatementEnd
