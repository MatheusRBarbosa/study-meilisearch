-- +goose Up
-- +goose StatementBegin
CREATE TABLE confirmations (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    confirmed boolean default false,
    "userId" int NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,

    FOREIGN KEY ("userId") REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE confirmations;
-- +goose StatementEnd
