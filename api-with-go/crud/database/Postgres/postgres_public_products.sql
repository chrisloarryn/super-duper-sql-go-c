create table products
(
    id           serial
        constraint products_id_pk
            primary key,
    name         varchar(100)            not null,
    observations varchar(100),
    price        integer                 not null,
    created_at   timestamp default now() not null,
    updated_at   timestamp
);

alter table products
    owner to postgres;

INSERT INTO public.products (id, name, observations, price, created_at, updated_at) VALUES (1, 'Curso DB con GO', 'Esta es una observación', 380, '2022-06-21 21:16:49.498756', null);
INSERT INTO public.products (id, name, observations, price, created_at, updated_at) VALUES (4, 'Curso de Comer popo', null, 500, '2022-06-21 21:19:24.612281', '2022-06-21 22:20:00.710331');
