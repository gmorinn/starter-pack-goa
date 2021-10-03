CREATE TABLE books (
  id uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  created_at timestamptz NOT NULL DEFAULT (NOW()),
  price float8 NOT NULL,
  name varchar(255) NOT NULL
);
