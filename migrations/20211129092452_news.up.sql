CREATE TABLE news (
    id bigserial not null unique,
    title varchar(40) not null,
    description varchar(1000) not null,
    photo text,
    time_date timestamp(0) default now()
);

CREATE TABLE users (
    id bigserial not null unique,
    name varchar(20) not null,
    surname varchar(20) not null,
    patronymic varchar(20) default '',
    town varchar(20) not null,
    age smallint not null,
    belt varchar(15) default '',
    weight numeric not null,
    id_iko varchar(20) default ''
);