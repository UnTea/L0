-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION add_customer(customer_id_in varchar, name_in varchar, phone_in varchar, zip_in varchar,
        city_in varchar, address_in varchar, region_in varchar, email_in varchar)
    RETURNS void AS
    $$
        BEGIN
            INSERT INTO customers (customer_id, name, phone, zip, city, address, region, email)
            VALUES (customer_id_in, name_in, phone_in, zip_in, city_in, address_in, region_in, email_in);
        END
    $$
    LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION add_order(order_uid_in varchar, customer_id_in varchar, track_number_in varchar,
        entry_in varchar, locale_in varchar, delivery_service_in varchar, shardkey_in varchar, sm_id_in int,
        date_created_in date, oof_shard_in varchar, internal_signature_in varchar)
    RETURNS void AS
    $$
        BEGIN
            INSERT INTO orders (order_uid, customer_id, track_number, entry, locale, delivery_service,
                                shardkey, sm_id, date_created, oof_shard, internal_signature)
                                VALUES (order_uid_in, customer_id_in, track_number_in, entry_in, locale_in,
                                        delivery_service_in,  shardkey_in, sm_id_in, date_created_in, oof_shard_in,
                                        internal_signature_in);
        END
    $$
    LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION add_order_items(order_id_in varchar, item_id_in int, status_in int)
    RETURNS void AS
    $$
        BEGIN
            INSERT INTO order_items (order_id, item_id, status)
            VALUES (order_id_in, item_id_in, status_in);
        END
    $$
    LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION add_items(chrt_id_in integer, price_in float, rid_in varchar, name_in varchar, sale_in int,
        size_in varchar, nm_id_in int, brand_in varchar)
    RETURNS void AS
    $$
        BEGIN
            INSERT INTO items (chrt_id, price, rid, name, sale, size, nm_id, brand)
            VALUES (chrt_id_in, price_in, rid_in, name_in, sale_in, size_in, nm_id_in, brand_in);
        END
    $$
    LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION add_payment(transaction_in varchar, request_id_in varchar, currency_in varchar,
        provider_in varchar, payment_dt_in int, bank_in varchar, delivery_cost_in float, goods_total_in float,
        custom_fee_in float)
    RETURNS void AS
    $$
        BEGIN
            INSERT INTO payments (transaction, request_id, currency, provider, payment_dt, bank, delivery_cost,
                                  goods_total, custom_fee)
                                  VALUES (transaction_in, request_id_in, currency_in, provider_in, payment_dt_in,
                                          bank_in, delivery_cost_in, goods_total_in, custom_fee_in);
        END
    $$
    LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_order()
    RETURNS TABLE
        (
            order_uid varchar,
            track_number varchar,
            entry varchar,
            locale varchar,
            internal_signature varchar,
            customer_id varchar,
            delivery_service varchar,
            shardkey varchar,
            sm_id int,
            date_created date,
            oof_shard varchar,
            name varchar,
            phone varchar,
            zip varchar,
            city varchar,
            address varchar,
            region varchar,
            email varchar,
            transaction varchar,
            request_id varchar,
            currency varchar,
            provider varchar,
            payment_dt int,
            bank varchar,
            delivery_cost float,
            goods_total float,
            custom_fee float
        )
        AS
        $$
            BEGIN
                RETURN QUERY
                    SELECT o.order_uid, o.track_number, o.entry, o.locale, o.internal_signature, o.customer_id,
                           o.delivery_service, o.shardkey, o.sm_id, o.date_created, o.oof_shard, c.name, c.phone, c.zip,
                           c.city, c.address, c.region, c.email, p.transaction, p.request_id, p.currency, p.provider,
                           p.payment_dt, p.bank, p.delivery_cost, p.goods_total, p.custom_fee
                    FROM orders o
                        JOIN customers c ON o.customer_id = c.customer_id
                        JOIN payments p ON o.order_uid = p.transaction;
            END
        $$
        LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_items()
    RETURNS TABLE
            (
                chrt_id int,
                price float,
                rid varchar,
                name varchar,
                sale int,
                size varchar,
                nm_id int,
                brand varchar,
                status int,
                order_id varchar
            )
            AS
            $$
            BEGIN
                RETURN QUERY
                    SELECT i.chrt_id, i.price, i.rid, i.name, i.sale, i.size, i.nm_id, i.brand, o.status, o.order_id
                    From items i
                        JOIN order_items o ON i.chrt_id = o.item_id;
            END
            $$
            LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS add_customer(varchar, varchar, varchar, varchar, varchar, varchar, varchar, varchar);
DROP FUNCTION IF EXISTS add_order( varchar, varchar, varchar, varchar, varchar, varchar, varchar, int, date, varchar, varchar);
DROP FUNCTION IF EXISTS add_items(integer, float, varchar, varchar, int, varchar, int, varchar);
DROP FUNCTION IF EXISTS add_order_items(varchar, int, int);
DROP FUNCTION IF EXISTS add_payment(varchar, varchar, varchar, varchar, int, varchar, float, float, float);
DROP FUNCTION IF EXISTS get_order();
DROP FUNCTION IF EXISTS get_items();
-- +goose StatementEnd
