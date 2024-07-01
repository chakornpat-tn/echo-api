CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "update_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_address" (
  "id" int PRIMARY KEY,
  "user_id" int,
  "recipient_name" varchar NOT NULL,
  "tel_number" varchar NOT NULL,
  "zip_code" varchar NOT NULL,
  "province" varchar NOT NULL,
  "address_detail" varchar NOT NULL
);

CREATE TABLE "shops" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "address" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "update_at" timestamp DEFAULT (now())
);

CREATE TABLE "products" (
  "id" int PRIMARY KEY,
  "shop_id" int,
  "title" varchar NOT NULL,
  "price" float NOT NULL,
  "detail" varchar,
  "created_at" timestamp DEFAULT (now()),
  "update_at" timestamp DEFAULT (now())
);

CREATE TABLE "products_categories" (
  "product_id" int,
  "category_id" int
);

CREATE TABLE "categories" (
  "id" int PRIMARY KEY,
  "title" varchar
);

CREATE TABLE "product_orders" (
  "product_id" int,
  "order_id" int,
  "qty" int DEFAULT 1
);

CREATE TABLE "orders" (
  "id" int PRIMARY KEY,
  "user_id" int,
  "address_id" int,
  "shop_id" int,
  "total_price" float NOT NULL,
  "courier_id" int,
  "order_statues" int,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "order_statuses" (
  "id" int PRIMARY KEY,
  "title" varchar
);

CREATE TABLE "couriers" (
  "id" int PRIMARY KEY,
  "title" varchar
);

ALTER TABLE "user_address" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "product_orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "product_orders" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("address_id") REFERENCES "user_address" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("courier_id") REFERENCES "couriers" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("order_statues") REFERENCES "order_statuses" ("id");
