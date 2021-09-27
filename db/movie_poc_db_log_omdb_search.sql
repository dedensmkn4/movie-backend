create table log__omdb_search
(
    id         int auto_increment
        primary key,
    path       varchar(255) not null,
    response   text         null,
    created_at datetime     not null
);
