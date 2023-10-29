CREATE TABLE "products" (
  "product_id" bigserial PRIMARY KEY,
  "shop_owner_name" varchar NOT NULL,
  "pic_path" varchar NOT NULL DEFAULT ' ',
  "describe" varchar NOT NULL DEFAULT ' ',
  "price" int NOT NULL DEFAULT 0,
  "quantity" int NOT NULL DEFAULT 0,
  "create_time" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "products" ("product_id");

CREATE INDEX ON "products" ("shop_owner_name");
