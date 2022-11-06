 BEGIN;

  CREATE EXTENSION pgcrypto;

  CREATE TYPE "role" AS ENUM (
    'admin',
    'pro',
    'user'
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
    "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
    "email" text NOT NULL CONSTRAINT emailchk CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    "password" text NOT NULL CONSTRAINT passwordchk CHECK (char_length(password) >= 9),
    "firstname" text CONSTRAINT firstnamechk CHECK (char_length(firstname) >= 2 AND char_length(firstname) <= 20 AND  firstname ~ '^[^0-9]*$') DEFAULT NULL,
    "lastname" text CONSTRAINT lastnamechk CHECK (char_length(lastname) >= 2 AND char_length(lastname) <= 20 AND lastname ~ '^[^0-9]*$') DEFAULT NULL,
    "password_confirm_code" text DEFAULT NULL CONSTRAINT code_passwordchk CHECK (char_length(password_confirm_code) = 5),
    "role" role NOT NULL DEFAULT 'user',
    "avatar" text DEFAULT NULL
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

COMMIT;
