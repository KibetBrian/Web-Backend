CREATE TYPE "role" AS ENUM (
  'Registration',
  'Support'
);

CREATE TABLE "Users" (
  "id" uuid PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "Voters" (
  "id" uuid PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "registered_at" timestamptz NOT NULL DEFAULT 'now()',
  "voted_at" timestamptz NOT NULL DEFAULT 'now()',
  "voters_public_address" varchar NOT NULL
);

CREATE TABLE "Admins" (
  "id" uuid PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "role" role NOT NULL
);

CREATE TABLE "Contestants" (
  "id" uuid PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "position" varchar NOT NULL,
  "registered_at" timestamptz DEFAULT 'now()',
  "description" varchar NOT NULL
);

CREATE INDEX ON "Admins" ("full_name");

CREATE INDEX ON "Contestants" ("full_name");

ALTER TABLE "Voters" ADD FOREIGN KEY ("email") REFERENCES "Users" ("email");

ALTER TABLE "Admins" ADD FOREIGN KEY ("email") REFERENCES "Users" ("email");

ALTER TABLE "Contestants" ADD FOREIGN KEY ("email") REFERENCES "Users" ("email");
