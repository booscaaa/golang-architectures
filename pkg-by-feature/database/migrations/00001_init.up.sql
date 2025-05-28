CREATE TABLE client (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL
);

CREATE TABLE product (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL
);

CREATE TABLE budget (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    client_id UUID NOT NULL REFERENCES client(id),
    product_id UUID NOT NULL REFERENCES product(id),
    amount INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);