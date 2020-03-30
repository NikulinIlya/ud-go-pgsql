-- +goose Up
ALTER TABLE product_items_collection ADD COLUMN prod_sec_name text NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE product_items_collection DROP COLUMN prod_sec_name;