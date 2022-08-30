create table if not exists delivery
(
    id      bigserial constraint delivery_pkey primary key,
    name    text,
    phone   text,
    zip     text,
    city    text,
    address text,
    region  text,
    email   text
);

create table if not exists payment
(
    id            bigserial constraint payment_pkey primary key,
    transaction   text,
    request_id    text,
    currency      text,
    provider      text,
    amount        bigint,
    payment_dt    bigint,
    bank          text,
    delivery_cost bigint,
    goods_total   bigint,
    custom_fee    bigint
);

create table if not exists items
(
    id           bigserial constraint items_pkey primary key,
    chrt_id      bigint,
    track_number text,
    price        bigint,
    rid          text,
    name         text,
    sale         bigint,
    size         text,
    total_price  bigint,
    nm_id        bigint,
    brand        text,
    status       bigint
);

create table if not exists model
(
    id                 bigserial constraint model_pkey primary key,
    order_uid          text,
    track_number       text,
    entry              text,
    locale             text,
    internal_signature text,
    customer_id        text,
    delivery_service   text,
    shardkey           text,
    sm_id              bigint,
    date_created       timestamp with time zone,
    oof_shard          text,
    delivery_id        bigint constraint fk_delivery references delivery,
    payment_id         bigint constraint fk_payment references payment,
    items_id           bigint constraint fk_items references items
);