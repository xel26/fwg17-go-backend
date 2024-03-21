CREATE TABLE "sizes"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "size" VARCHAR NOT NULL UNIQUE,
    "additionalPrice" INT DEFAULT 0,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP
)