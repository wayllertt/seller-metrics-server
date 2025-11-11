CREATE TABLE sellers (
    id          SERIAL PRIMARY KEY,
    name        TEXT        NOT NULL,
    is_blocked  BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMP   NOT NULL DEFAULT NOW()
);

CREATE TABLE orders (
    id           SERIAL PRIMARY KEY,
    seller_id    INT         NOT NULL REFERENCES sellers(id),
    created_at   TIMESTAMP   NOT NULL,
    promised_at  TIMESTAMP   NOT NULL,
    delivered_at TIMESTAMP
);
