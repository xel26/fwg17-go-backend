CREATE TABLE "intRandom"(
    "id" SERIAL PRIMARY KEY NOT NULL,
    "intRand" VARCHAR NOT NULL UNIQUE,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP
)