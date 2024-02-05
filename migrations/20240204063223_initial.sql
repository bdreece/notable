-- Create "users" table
CREATE TABLE "public"."users" (
 "id" uuid NOT NULL,
 "created_at" timestamp NOT NULL,
 "updated_at" timestamp NOT NULL,
 "first_name" character varying(63) NOT NULL,
 "last_name" character varying(63) NOT NULL,
 "bio" text NULL,
 "email_address" character varying(127) NOT NULL,
 "email_verified" boolean NOT NULL,
 "hash" bytea NOT NULL,
 PRIMARY KEY ("id")
);
-- Create index "ix_users_email_address" to table: "users"
CREATE UNIQUE INDEX "ix_users_email_address" ON "public"."users" ("email_address");
-- Create index "ix_users_name" to table: "users"
CREATE INDEX "ix_users_name" ON "public"."users" ("last_name", "first_name");
-- Create "devices" table
CREATE TABLE "public"."devices" (
 "id" uuid NOT NULL,
 "created_at" timestamp NOT NULL,
 "updated_at" timestamp NOT NULL,
 "mac_address" macaddr NOT NULL,
 "ip_address" inet NULL,
 "connected" boolean NOT NULL,
 "user_id" uuid NOT NULL,
 PRIMARY KEY ("id"),
 CONSTRAINT "devices_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "ix_devices_connected" to table: "devices"
CREATE INDEX "ix_devices_connected" ON "public"."devices" ("user_id", "connected");
-- Create index "ix_devices_ip_address" to table: "devices"
CREATE INDEX "ix_devices_ip_address" ON "public"."devices" ("ip_address");
-- Create index "ix_devices_mac_address" to table: "devices"
CREATE UNIQUE INDEX "ix_devices_mac_address" ON "public"."devices" ("mac_address");
-- Create "directories" table
CREATE TABLE "public"."directories" (
 "id" uuid NOT NULL,
 "created_at" timestamp NOT NULL,
 "updated_at" timestamp NOT NULL,
 "name" character varying(127) NOT NULL,
 "description" text NULL,
 "parent_id" uuid NULL,
 "user_id" uuid NOT NULL,
 PRIMARY KEY ("id"),
 CONSTRAINT "directories_parent_id_fkey" FOREIGN KEY ("parent_id") REFERENCES "public"."directories" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
 CONSTRAINT "directories_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "ix_directories_parent" to table: "directories"
CREATE INDEX "ix_directories_parent" ON "public"."directories" ("parent_id");
-- Create index "ix_directories_users_name" to table: "directories"
CREATE INDEX "ix_directories_users_name" ON "public"."directories" ("user_id", "name");
-- Create "notes" table
CREATE TABLE "public"."notes" (
 "id" uuid NOT NULL,
 "created_at" timestamp NOT NULL,
 "updated_at" timestamp NOT NULL,
 "title" character varying(127) NOT NULL,
 "content" jsonb NOT NULL,
 "directory_id" uuid NOT NULL,
 "user_id" uuid NOT NULL,
 PRIMARY KEY ("id"),
 CONSTRAINT "notes_directory_id_fkey" FOREIGN KEY ("directory_id") REFERENCES "public"."directories" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
 CONSTRAINT "notes_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "ix_notes_users_directories_title" to table: "notes"
CREATE INDEX "ix_notes_users_directories_title" ON "public"."notes" ("user_id", "directory_id", "title");
