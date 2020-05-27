-- TODO: add working hours
CREATE TABLE IF NOT EXISTS places (
     id             bigserial primary key,
     name           varchar(80) NOT NULL,
     minCost        integer NOT NULL,
     grade          float NOT NULL,
     image          varchar(200),
     logo           varchar(200),
     latitude       float NOT NULL,
     longitude      float NOT NULL,
     deliveryTime   integer NOT NULL
);

CREATE TABLE IF NOT EXISTS dishes (
    id          bigserial primary key,
    name        varchar(120) NOT NULL,
    cafe        bigint REFERENCES places(id) ON DELETE RESTRICT,
    ingredients varchar(200),
    calories    integer,
    weight      integer,
    cost        integer NOT NULL,
    image       varchar(200),
    inStock     boolean default true,
    tag         integer NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id         bigserial primary key,
    "user"     bigint NOT NULL,
    address    varchar(100),
    cafe       bigint REFERENCES places(id) ON DELETE RESTRICT,
    status     integer NOT NULL,
    cost       integer NOT NULL,
    date       timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS orders_details (
    order_id    bigint REFERENCES orders(id) ON DELETE RESTRICT,
    dish_id     bigint REFERENCES dishes(id) ON DELETE RESTRICT,
    count       integer NOT NULL
);

CREATE INDEX idx_orders_details ON orders_details(order_id);
