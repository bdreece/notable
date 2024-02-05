-- Modify "devices" table
ALTER TABLE "public"."devices" ALTER COLUMN "id" SET DEFAULT gen_random_uuid(), ALTER COLUMN "created_at" SET DEFAULT now(), ALTER COLUMN "updated_at" SET DEFAULT now();
-- Modify "directories" table
ALTER TABLE "public"."directories" ALTER COLUMN "id" SET DEFAULT gen_random_uuid(), ALTER COLUMN "created_at" SET DEFAULT now(), ALTER COLUMN "updated_at" SET DEFAULT now();
-- Modify "notes" table
ALTER TABLE "public"."notes" ALTER COLUMN "id" SET DEFAULT gen_random_uuid(), ALTER COLUMN "created_at" SET DEFAULT now(), ALTER COLUMN "updated_at" SET DEFAULT now();
-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "id" SET DEFAULT gen_random_uuid(), ALTER COLUMN "created_at" SET DEFAULT now(), ALTER COLUMN "updated_at" SET DEFAULT now(), ALTER COLUMN "email_verified" SET DEFAULT false;
