CREATE TABLE "clients" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "surname" varchar NOT NULL,
    "patronymic" varchar,
    "age" int NOT NULL,
    "gender" varchar NOT NULL,
    "country_id" varchar NOT NULL
)