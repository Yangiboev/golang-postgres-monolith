CREATE TABLE IF NOT EXISTS categories(
    id uuid PRIMARY KEY,
    name VARCHAR,
    parent_id uuid,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);
create index if not exists categories_i1 on categories (parent_id);


CREATE TABLE IF NOT EXISTS retailers(
    id uuid PRIMARY KEY,
    name VARCHAR(150) not null,
    website VARCHAR(150) not null,
    description TEXT,
    created_at timestamp default current_timestamp
);
CREATE TABLE IF NOT EXISTS retailer_branches(
    id uuid PRIMARY KEY,
    retailer_id uuid NOT NULL REFERENCES retailers(id),
    name VARCHAR(150) not null,
    phone_number INT not null,
    address TEXT not null,
    description TEXT,
    long float8,
    lat float8,
    created_at timestamp default current_timestamp
);
create index if not exists retailer_branches_i1 on retailer_branches (retailer_id);


CREATE TABLE IF NOT EXISTS products(
    id uuid PRIMARY KEY,
    name VARCHAR NOT NULL,
    image VARCHAR NOT NULL,
    description TEXT,
    category_id uuid NOT NULL REFERENCES categories(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp
);
create index if not exists products_i1 on products (name);

CREATE TABLE IF NOT EXISTS product_prices(
    id uuid PRIMARY KEY,
    price INT NOT NULL,
    product_id uuid NOT NULL REFERENCES products(id),
    retailer_id uuid NOT NULL REFERENCES retailers(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp
);
create index if not exists product_prices_i1 on product_prices (price);
create index if not exists product_prices_i2 on product_prices (product_id);
