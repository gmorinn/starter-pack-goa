CREATE TABLE books (
  id uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  created_at timestamptz NOT NULL DEFAULT (NOW()),
  price float8 NOT NULL,
  name varchar(255) NOT NULL
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