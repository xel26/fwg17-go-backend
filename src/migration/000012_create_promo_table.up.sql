CREATE TABLE "promo"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" VARCHAR UNIQUE NOT NULL,
    "code" VARCHAR UNIQUE NOT NULL,
    "description" TEXT,
    "percentage" FLOAT NOT NULL,
    "isExpired" BOOLEAN,
    "maximumPromo" INT NOT NULL,
    "minimumAmount" INT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP
)