CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE tb_events(
    id_event SERIAL,
    name_event VARCHAR(100) NOT NULL,
    description_event VARCHAR(255),
    location_event VARCHAR(255) NOT NULL ,
    photo_event VARCHAR(255),
    date_event TIMESTAMP NOT NULL,
    embeddings_event VECTOR(1536),
    active_event BOOLEAN NOT NULL DEFAULT TRUE,
    canceled_event BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_event BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (id_event)
);

CREATE TABLE tb_user_events(
    id_user_event SERIAL,
    fk_id_user INTEGER NOT NULL,
    fk_id_event INTEGER NOT NULL,
    participation_status_event INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (id_user_event)
);