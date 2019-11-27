create table "user"
(
    id            serial                                not null,
    email         varchar                               not null,
    password_hash varchar                               not null,
    name          varchar                               not null,
    avatar_path   varchar default ''::character varying not null
);