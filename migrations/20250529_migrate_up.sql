CREATE TABLE IF NOT EXISTS customers(
    id serial PRIMARY KEY NOT NULL,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS wallets(
    wallet_id VARCHAR(255) PRIMARY KEY NOT NULL,
    customer_id serial NOT NULL,
    CONSTRAINT fk_customers FOREIGN KEY(customer_id) REFERENCES customers(id)
);