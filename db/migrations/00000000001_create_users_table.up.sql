create table users (
    id         varchar(100) not null,
    name       varchar(100) not null,
    password   varchar(100) not null,
    token      varchar(100) null,
    created_at bigint not null,
    updated_at bigint not null,
    primary key(id)
) engine = InnoDB;

-- id varchar(100) not null,
--     username varchar(100) not null,
--     name varchar(100) not null, 
--     password varchar(100) not null,
--     created_at timestamp default current_timestamp not null,
--     updated_at timestamp default current_timestamp on update current_timestamp not null,
--     constraint username_unique unique (username),
    -- created_at bigint       not null default current_timestamp,
    -- updated_at bigint       not null default current_timestamp on update current_timestamp,