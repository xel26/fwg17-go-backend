CREATE TABLE "orderDetails"(
    "id" SERIAL PRIMARY KEY NOT NULL,
    "productId" INT NOT NULL,
    "sizeId" INT NOT NULL,
    "variantId" INT NOT NULL,
    "quantity" INT NOT NULL,
    "orderId" INT NOT NULL,
    "subtotal" INT,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP,
    Foreign Key ("productId") REFERENCES "products" ("id"),
    Foreign Key ("sizeId") REFERENCES "sizes" ("id"),
    Foreign Key ("variantId") REFERENCES "variant" ("id"),
    Foreign Key ("orderId") REFERENCES "orders" ("id")
)