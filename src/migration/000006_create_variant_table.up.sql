CREATE TABLE "variant"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL UNIQUE,
    "additionalPrice" INT DEFAULT 0,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP
)