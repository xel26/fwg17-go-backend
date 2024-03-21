CREATE TABLE "productVariant"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "productId" INT NOT NULL,
    "variantId" INT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP,
    FOREIGN KEY ("productId") REFERENCES "products" ("id"),
    FOREIGN KEY ("variantId") REFERENCES "variant" ("id")
)