CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT 'uuid_generate_v4()',
  "fullname" varchar NOT NULL,
  "username" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "is_active" bool NOT NULL DEFAULT true
);
