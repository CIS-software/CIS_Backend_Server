CREATE TABLE news (
    id bigserial not null unique,
    title varchar(40) not null,
    description varchar(1000) not null,
    photo text,
    time_date timestamp(0) default now()
);
