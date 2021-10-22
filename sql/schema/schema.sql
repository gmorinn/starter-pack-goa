CREATE TABLE "books" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "price" float8 NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "lastname" text NOT NULL,
  "firstname" text NOT NULL
);

CREATE TABLE "refresh_token" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamptz,
  "ip" text NOT NULL,
  "user_agent" text NOT NULL,
  "token" text NOT NULL,
  "expir_on" timestamptz NOT NULL,
  "user_id" uuid NOT NULL
);

-- ALTER TABLE "refresh_token" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
