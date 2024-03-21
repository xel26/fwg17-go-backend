create table "products"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" VARCHAR UNIQUE NOT NULL,
    "description" TEXT,
    "basePrice" INT NOT NULL,
    "discount" INT DEFAULT 0,
    "image" TEXT,
    "isRecommended" BOOLEAN,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP
)