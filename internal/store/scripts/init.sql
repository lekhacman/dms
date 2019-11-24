create database dms;

create table objects
(
    id          uuid         not null unique,
    name        varchar(255) not null,
    description varchar(255),
    size        numeric(10, 0),        -- unit32 10 digits (precision), 0 decimal point (scale)
    created_at  time         not null, -- UTC
    updated_at  time         not null, -- UTC
    primary key (id)
);
