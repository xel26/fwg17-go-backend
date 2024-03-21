CREATE TABLE "forgotPassword"(
    "id" SERIAL PRIMARY KEY NOT NULL,
    "otp" VARCHAR,
    "email" VARCHAR,
    "userId" INT,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP,
    Foreign Key ("userId") REFERENCES "users" ("id")
)