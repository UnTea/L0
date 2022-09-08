-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS customers (
  customer_id VARCHAR(40) PRIMARY KEY,
  name VARCHAR(50),
  phone VARCHAR(20) NOT NULL,
  zip VARCHAR(10),
  city VARCHAR(40),
  address VARCHAR(100),
  region VARCHAR(50),
  email VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
    chrt_id INT PRIMARY KEY,
    price float NOT NULL,
    rid VARCHAR(30),
    name VARCHAR(100) NOT NULL,
    sale INT CHECK ( sale >= 0 AND sale <= 100 ),
    size VARCHAR(30),
    nm_id INT,
    brand VARCHAR(70)
);

CREATE TABLE IF NOT EXISTS orders (
    order_uid VARCHAR PRIMARY KEY,
    customer_id VARCHAR(40) NOT NULL,
    track_number VARCHAR(30),
    entry VARCHAR(20),
    locale VARCHAR(5),
    delivery_service VARCHAR(40),
    shardkey VARCHAR(10),
    sm_id INT,
    date_created DATE NOT NULL DEFAULT now(),
    oof_shard VARCHAR(10),
    internal_signature VARCHAR,
    FOREIGN KEY (customer_id) REFERENCES customers (customer_id)
);

CREATE TABLE IF NOT EXISTS order_items (
    order_id VARCHAR NOT NULL,
    item_id INT NOT NULL,
    status INT NOT NULL,
    PRIMARY KEY (order_id, item_id),
    FOREIGN KEY (order_id) REFERENCES orders (order_uid),
    FOREIGN KEY (item_id) REFERENCES  items (chrt_id)
);

CREATE TABLE IF NOT EXISTS payments (
    transaction VARCHAR PRIMARY KEY,
    request_id VARCHAR(20),
    currency VARCHAR(10) NOT NULL ,
    provider VARCHAR(10),
    payment_dt INT,
    bank VARCHAR(20),
    delivery_cost float NOT NULL,
    goods_total float NOT NULL,
    custom_fee float NOT NULL,
    FOREIGN KEY (transaction) REFERENCES orders (order_uid)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS payments, order_items, orders, items, customers

-- +goose StatementEnd
