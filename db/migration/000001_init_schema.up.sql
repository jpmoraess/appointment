CREATE TABLE "tenant" (
                          "id" bigserial PRIMARY KEY,
                          "name" varchar NOT NULL,
                          "slug" varchar UNIQUE NOT NULL,
                          "created_at" timestamptz NOT NULL DEFAULT (now()),
                          "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user" (
                        "id" bigserial PRIMARY KEY,
                        "tenant_id" bigint NOT NULL,
                        "username" varchar UNIQUE NOT NULL,
                        "password" varchar NOT NULL,
                        "role" varchar NOT NULL,
                        "created_at" timestamptz NOT NULL DEFAULT (now()),
                        "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "session" (
                           "id" bigserial PRIMARY KEY,
                           "user_id" bigint NOT NULL,
                           "refresh_token" varchar UNIQUE NOT NULL,
                           "user_agent" varchar,
                           "ip_address" varchar,
                           "is_active" boolean NOT NULL DEFAULT true,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "expires_at" timestamptz NOT NULL
);

CREATE TABLE "professional" (
                                "id" bigserial PRIMARY KEY,
                                "tenant_id" bigint NOT NULL,
                                "user_id" bigint UNIQUE NOT NULL,
                                "name" varchar NOT NULL,
                                "email" varchar,
                                "phone" varchar,
                                "created_at" timestamptz NOT NULL DEFAULT (now()),
                                "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "service" (
                           "id" bigserial PRIMARY KEY,
                           "tenant_id" bigint NOT NULL,
                           "name" varchar NOT NULL,
                           "description" varchar,
                           "duration_minutes" integer NOT NULL,
                           "price" decimal(10,2) NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "professional_service" (
                                        "professional_id" bigint NOT NULL,
                                        "service_id" bigint NOT NULL,
                                        "price" decimal(10,2),
                                        "duration_minutes" integer,
                                        "created_at" timestamptz NOT NULL DEFAULT (now()),
                                        "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "schedule" (
                            "id" bigserial PRIMARY KEY,
                            "professional_id" bigint NOT NULL,
                            "day_of_week" integer NOT NULL,
                            "start_time" time NOT NULL,
                            "end_time" time NOT NULL,
                            "interval_minutes" integer NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "schedule_override" (
                                     "id" bigserial PRIMARY KEY,
                                     "professional_id" bigint NOT NULL,
                                     "override_date" date NOT NULL,
                                     "start_time" time NOT NULL,
                                     "end_time" time NOT NULL,
                                     "is_available" boolean NOT NULL,
                                     "reason" varchar,
                                     "created_at" timestamptz NOT NULL DEFAULT (now()),
                                     "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "appointment" (
                               "id" bigserial PRIMARY KEY,
                               "tenant_id" bigint NOT NULL,
                               "professional_id" bigint NOT NULL,
                               "service_id" bigint NOT NULL,
                               "customer_name" varchar NOT NULL,
                               "appointment_date" timestamptz NOT NULL,
                               "status" varchar NOT NULL DEFAULT 'SCHEDULED',
                               "created_at" timestamptz NOT NULL DEFAULT (now()),
                               "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "tenant" ("slug");

CREATE INDEX ON "user" ("tenant_id");

CREATE INDEX ON "session" ("user_id");

CREATE INDEX ON "professional" ("tenant_id");

CREATE INDEX ON "professional" ("user_id");

CREATE INDEX ON "service" ("tenant_id");

CREATE INDEX ON "professional_service" ("professional_id");

CREATE INDEX ON "professional_service" ("service_id");

CREATE INDEX ON "schedule" ("professional_id");

CREATE INDEX ON "schedule_override" ("professional_id");

CREATE INDEX ON "appointment" ("tenant_id");

CREATE INDEX ON "appointment" ("professional_id");

COMMENT ON COLUMN "user"."role" IS 'admin, professional, customer';

COMMENT ON COLUMN "professional_service"."price" IS 'if the price varies by professional';

COMMENT ON COLUMN "professional_service"."duration_minutes" IS 'if the duration varies by professional';

COMMENT ON COLUMN "schedule"."day_of_week" IS '0 (sunday) 6 (saturday)';

COMMENT ON COLUMN "schedule_override"."reason" IS 'reason for override';

ALTER TABLE "user" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id");

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "professional" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id");

ALTER TABLE "professional" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "service" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id");

ALTER TABLE "professional_service" ADD FOREIGN KEY ("professional_id") REFERENCES "professional" ("id");

ALTER TABLE "professional_service" ADD FOREIGN KEY ("service_id") REFERENCES "service" ("id");

ALTER TABLE "schedule" ADD FOREIGN KEY ("professional_id") REFERENCES "professional" ("id");

ALTER TABLE "schedule_override" ADD FOREIGN KEY ("professional_id") REFERENCES "professional" ("id");

ALTER TABLE "appointment" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id");

ALTER TABLE "appointment" ADD FOREIGN KEY ("professional_id") REFERENCES "professional" ("id");

ALTER TABLE "appointment" ADD FOREIGN KEY ("service_id") REFERENCES "service" ("id");
