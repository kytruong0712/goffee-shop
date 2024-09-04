CREATE TABLE categories
(
    id              BIGINT                   NOT NULL PRIMARY KEY,
    "name"          TEXT                     NOT NULL CONSTRAINT categories_name_check CHECK ("name" <> '' :: TEXT),
    description     TEXT                     NOT NULL CONSTRAINT categories_description_check CHECK (description <> '' :: TEXT),
    status          TEXT                     NOT NULL CONSTRAINT categories_status_check CHECK (status <> '' :: TEXT),
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT categories_name_uindex UNIQUE (name)
);
CREATE INDEX IF NOT EXISTS categories_name_index ON categories("name");

CREATE TABLE menu_items (
    id              BIGINT                   NOT NULL PRIMARY KEY,
    category_id     BIGINT                   NOT NULL CONSTRAINT category_id_fkey REFERENCES categories (id),
    "name"          TEXT                     NOT NULL CONSTRAINT menu_items_name_check CHECK ("name" <> '' :: TEXT),
    description     TEXT                     NOT NULL CONSTRAINT menu_items_description_check CHECK (description <> '' :: TEXT),
    status          TEXT                     NOT NULL CONSTRAINT menu_items_status_check CHECK (status <> '' :: TEXT),
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT menu_items_name_uindex UNIQUE (name)
);
CREATE INDEX IF NOT EXISTS menu_items_name_index ON menu_items("name");

CREATE TABLE menu_item_prices (
    id              BIGINT                   NOT NULL PRIMARY KEY,
    menu_item_id    BIGINT                   NOT NULL CONSTRAINT menu_item_id_fkey REFERENCES menu_items (id),
    "size"          TEXT                     NOT NULL CONSTRAINT menu_item_prices_size_check CHECK ("size"  <> '' :: TEXT),
    price           DECIMAL(10, 2)           NOT NULL CONSTRAINT menu_item_prices_price_check CHECK (price > 0.00 :: DECIMAL),
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT menu_item_prices_menu_item_id_size_uindex UNIQUE (menu_item_id, size)
);

CREATE TABLE menu_item_options (
    id                  BIGINT                   NOT NULL PRIMARY KEY,
    menu_item_id        BIGINT                   NOT NULL CONSTRAINT menu_item_id_fkey REFERENCES menu_items (id),
    "name"              TEXT                     NOT NULL CONSTRAINT menu_item_options_name_check CHECK ("name" <> '' :: TEXT),
    additional_price    DECIMAL(10, 2)                    DEFAULT 0.00,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT menu_item_options_name_uindex UNIQUE (name)
);
CREATE INDEX IF NOT EXISTS menu_item_options_name_index ON menu_item_options("name");
