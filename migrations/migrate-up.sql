CREATE TABLE news (
    id bigserial not null unique,
    title varchar(50) not null,
    description varchar(1500) not null,
    photo varchar(80),
    time_date timestamp(0) default now()
);

CREATE TABLE user_auth (
    id bigserial not null unique primary key,
    email varchar(50) not null,
    user_type varchar(20) default 'user',
    encrypted_password varchar(100) not null
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

CREATE TYPE training_week AS ENUM ('Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс');
CREATE TABLE training_calendar (
    day training_week not null unique,
    description varchar(50) not null
);

