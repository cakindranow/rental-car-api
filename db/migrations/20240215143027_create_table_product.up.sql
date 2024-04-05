create table products (
    id varchar not null ,
    name varchar not null,
    price bigint not null,
    image_url varchar not null,
    deleted_at timestamp null,
    constraint products_pk primary key (id)
)