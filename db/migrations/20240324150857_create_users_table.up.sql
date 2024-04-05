create table users (
                          id varchar not null ,
                          name varchar not null,
                          email varchar not null,
                          password varchar not null,
                          deleted_at timestamp null,
                          constraint users_pk primary key (id)
)