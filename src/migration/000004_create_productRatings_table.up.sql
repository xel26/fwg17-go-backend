CREATE TABLE "productRatings"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "productId" INT NOT NULL,
    "rate" INT NOT NULL,
    "userId" INT NOT NULL,
    "reviewMessage" TEXT,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP,
    FOREIGN KEY ("productId") REFERENCES "products" ("id"),
    FOREIGN KEY ("userId") REFERENCES "users" ("id"),
    CHECK ("rate" > 0 AND "rate" < 6)
)