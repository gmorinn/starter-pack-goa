CREATE TYPE "role" AS ENUM (
  'admin',
  'pro',
  'user'
);

CREATE TYPE "categories" AS ENUM (
  'men',
  'women',
  'sneaker',
  'hat',
  'jacket',
  'nothing'
);

CREATE TABLE "products" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamptz,
  "name" text NOT NULL,
  "category" categories NOT NULL DEFAULT 'nothing',
  "cover" text NOT NULL,
  "price" float8 NOT NULL DEFAULT 0
);

CREATE TABLE "files" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamptz,
  "name" text,
  "url" text,
  "mime" text,
  "size" bigint
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "lastname" text NOT NULL,
  "firstname" text NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "role" role NOT NULL DEFAULT 'user',
  "birthday" text,
  "phone" text,
  "password_confirm_code" text,
  "firebase_id_token" text DEFAULT NULL,
  "firebase_uid" text DEFAULT NULL,
  "firebase_provider" text DEFAULT NULL
);

CREATE TABLE "refresh_token" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamptz,
  "token" text NOT NULL,
  "expir_on" timestamptz NOT NULL,
  "user_id" uuid NOT NULL
);

ALTER TABLE "refresh_token" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
