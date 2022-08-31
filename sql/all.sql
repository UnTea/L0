create table if not exists deliveries
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

create table if not exists payments
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

create table if not exists models
(
    id                 bigserial constraint model_pkey primary key,
    order_uid          text,
    track_number       text,
    entry              text,
    locale             text,
    internal_signature text,
    customer_id        text,
    delivery_service   text,
    shard_key          text,
    sm_id              bigint,
    date_created       timestamp with time zone,
    oof_shard          text,
    delivery_id        bigint constraint fk_delivery references deliveries,
    payment_id         bigint constraint fk_payment references payments
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
    status       bigint,
    model_id     int,
    constraint fk_items foreign key(model_id) references models(id)
);