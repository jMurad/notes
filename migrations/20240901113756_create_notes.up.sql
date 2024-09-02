CREATE TABLE notes (
    id bigserial not null primary key,
    content varchar not null,
    created timestamp not null,
    author_id bigserial not null
);