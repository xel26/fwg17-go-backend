CREATE TABLE "productImages"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "productId" INT NOT NULL,
    "imageUrl" VARCHAR NOT NULL,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP,
    FOREIGN KEY ("productId") REFERENCES "products" ("id")
)