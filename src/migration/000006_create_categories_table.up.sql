CREATE TABLE "categories"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL UNIQUE,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP
)