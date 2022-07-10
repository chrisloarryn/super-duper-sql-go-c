create table invoice_headers
(
    id         serial
        constraint invoice_headers_id_pk
            primary key,
    client     varchar(100)            not null,
    created_at timestamp default now() not null,
    updated_at timestamp
);

alter table invoice_headers
    owner to postgres;

INSERT INTO public.invoice_headers (id, client, created_at, updated_at) VALUES (1, 'Alvaro Felipe', '2022-06-22 04:37:33.043836', null);
INSERT INTO public.invoice_headers (id, client, created_at, updated_at) VALUES (2, 'Alvaro Felipe', '2022-06-22 04:41:11.555146', null);
INSERT INTO public.invoice_headers (id, client, created_at, updated_at) VALUES (3, 'Alvaro Felipe', '2022-06-22 04:42:43.515154', null);
INSERT INTO public.invoice_headers (id, client, created_at, updated_at) VALUES (8, 'Alvaro Felipe', '2022-06-22 04:44:52.554059', null);
