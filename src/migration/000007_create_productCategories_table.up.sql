CREATE TABLE "productCategories"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "productId" INT NOT NULL,
    "categoryId" INT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP,
    FOREIGN KEY ("productId") REFERENCES "products" ("id"),
    FOREIGN KEY ("categoryId") REFERENCES "categories" ("id")
)