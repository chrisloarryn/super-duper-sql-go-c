create table invoice_items
(
    id                serial
        constraint invoice_items_id_pk
            primary key,
    invoice_header_id integer                 not null
        constraint invoice_items_invoice_header_id_fk
            references invoice_headers
            on update restrict on delete restrict,
    product_id        integer                 not null
        constraint invoice_items_product_id_fk
            references products
            on update restrict on delete restrict,
    created_at        timestamp default now() not null,
    updated_at        timestamp
);

alter table invoice_items
    owner to postgres;

INSERT INTO public.invoice_items (id, invoice_header_id, product_id, created_at, updated_at) VALUES (1, 1, 1, '2022-06-22 04:37:33.043836', null);
INSERT INTO public.invoice_items (id, invoice_header_id, product_id, created_at, updated_at) VALUES (2, 1, 4, '2022-06-22 04:37:33.043836', null);
INSERT INTO public.invoice_items (id, invoice_header_id, product_id, created_at, updated_at) VALUES (3, 2, 1, '2022-06-22 04:41:11.555146', null);
INSERT INTO public.invoice_items (id, invoice_header_id, product_id, created_at, updated_at) VALUES (4, 2, 4, '2022-06-22 04:41:11.555146', null);
INSERT INTO public.invoice_items (id, invoice_header_id, product_id, created_at, updated_at) VALUES (5, 3, 1, '2022-06-22 04:42:43.515154', null);
INSERT INTO public.invoice_items (id, invoice_header_id, product_id, created_at, updated_at) VALUES (6, 3, 4, '2022-06-22 04:42:43.515154', null);
INSERT INTO public.invoice_items (id, invoice_header_id, product_id, created_at, updated_at) VALUES (15, 8, 1, '2022-06-22 04:44:52.554059', null);
INSERT INTO public.invoice_items (id, invoice_header_id, product_id, created_at, updated_at) VALUES (16, 8, 4, '2022-06-22 04:44:52.554059', null);
