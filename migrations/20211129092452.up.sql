CREATE TABLE news (
    id bigserial not null unique,
    title varchar(40) not null,
    description varchar(1000) not null,
    photo text,
    time_date timestamp(0) default now()
);
CREATE TABLE user_auth (
    id bigserial not null unique primary key,
    email varchar(50) not null,
    user_type varchar(20) default 'user',
    encrypted_password varchar(100) not null,
    access_token varchar(200),
    refresh_token varchar(200)
);
CREATE TABLE user_profile (
    user_id bigint not null unique primary key,
    name varchar(20) not null,
    surname varchar(20) not null,
    town varchar(20) not null,
    age varchar(20) not null,
    belt varchar(15),
    weight numeric,
    id_iko varchar(20),
    foreign key (user_id) references user_auth (id) on delete cascade on update cascade
);
CREATE TABLE training_calendar (
    id bigserial not null unique,
    date date not null,
    description varchar(50) not null
);