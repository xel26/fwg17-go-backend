CREATE TABLE "users"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "fullName" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL,
    "address" TEXT,
    "picture" TEXT,
    "phoneNumber" VARCHAR(20) UNIQUE,
    "role" VARCHAR(15) DEFAULT 'customer',
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP
)