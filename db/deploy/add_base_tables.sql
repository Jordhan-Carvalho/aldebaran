-- Deploy aldebaran:add_base_tables to pg

BEGIN;

CREATE TYPE "gender_enum" AS ENUM (
  'male',
  'female',
  'none'
);

CREATE TYPE "owner_type_enum" AS ENUM (
  'worker',
  'client'
);

CREATE TYPE "job_status_enum" AS ENUM (
  'waiting_for_worker',
  'waiting_for_payment',
  'waiting',
  'done',
  'canceled'
);

CREATE TYPE "day_of_week_enum" AS ENUM (
  'Monday',
  'Tuesday',
  'Wednesday',
  'Thursday',
  'Friday',
  'Saturday',
  'Sunday'
);

CREATE TABLE "user_client" (
  "id" int PRIMARY KEY,
  "email" varchar,
  "phone" varchar,
  "full_name" varchar,
  "gender" gender_enum,
  "last_access" timestamptz,
  "updated_at" timestamptz DEFAULT (now()),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "user_worker" (
  "id" int PRIMARY KEY,
  "email" varchar,
  "phone" varchar,
  "full_name" varchar,
  "gender" gender_enum,
  "description" varchar,
  "last_access" timestamptz,
  "updated_at" timestamptz DEFAULT (now()),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "rate" (
  "id" int PRIMARY KEY,
  "created_by" owner_type_enum,
  "client_id" int,
  "worker_id" int,
  "rate" int,
  "updated_at" timestamptz DEFAULT (now()),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "worker_service" (
  "id" int PRIMARY KEY,
  "service_id" int,
  "worker_id" int,
  "price" int,
  "note" varchar,
  "updated_at" timestamptz DEFAULT (now()),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "service" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "description" varchar,
  "updated_at" timestamptz DEFAULT (now()),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "selected_service" (
  "id" int PRIMARY KEY,
  "job_id" int,
  "service_id" int,
  "price_history" int,
  "note_history" varchar,
  "updated_at" timestamptz DEFAULT (now()),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "job" (
  "id" int PRIMARY KEY,
  "worker" int,
  "client" int,
  "total_price" int,
  "scheduling_date" timestamptz,
  "confirmation_date" timestamptz,
  "payment_date" timestamptz,
  "work_date" timestamptz,
  "status" job_status_enum,
  "updated_at" timestamptz DEFAULT (now()),
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "date" (
  "id" int PRIMARY KEY,
  "day" varchar,
  "month" varchar,
  "year" varchar,
  "full_date" varchar,
  "holiday" boolean,
  "is_weekend" boolean,
  "day_of_week" day_of_week_enum
);

CREATE TABLE "schedule" (
  "day_id" int,
  "worker_id" int,
  "is_available" boolean
);

ALTER TABLE "rate" ADD FOREIGN KEY ("client_id") REFERENCES "user_client" ("id");

ALTER TABLE "rate" ADD FOREIGN KEY ("worker_id") REFERENCES "user_worker" ("id");

ALTER TABLE "worker_service" ADD FOREIGN KEY ("service_id") REFERENCES "service" ("id");

ALTER TABLE "worker_service" ADD FOREIGN KEY ("worker_id") REFERENCES "user_worker" ("id");

ALTER TABLE "selected_service" ADD FOREIGN KEY ("job_id") REFERENCES "job" ("id");

ALTER TABLE "selected_service" ADD FOREIGN KEY ("service_id") REFERENCES "service" ("id");

ALTER TABLE "job" ADD FOREIGN KEY ("worker") REFERENCES "user_worker" ("id");

ALTER TABLE "job" ADD FOREIGN KEY ("client") REFERENCES "user_client" ("id");

ALTER TABLE "schedule" ADD FOREIGN KEY ("day_id") REFERENCES "date" ("id");

ALTER TABLE "schedule" ADD FOREIGN KEY ("worker_id") REFERENCES "user_worker" ("id");


COMMIT;
