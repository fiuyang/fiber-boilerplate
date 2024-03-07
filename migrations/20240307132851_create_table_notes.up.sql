CREATE TABLE "notes" (
  "id" bigserial PRIMARY KEY,
  "content" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
