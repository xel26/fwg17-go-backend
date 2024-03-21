CREATE TABLE "message"(
    "id" SERIAL PRIMARY KEY NOT NULL,
    "recipientId" INT NOT NULL,
    "senderId" INT NOT NULL,
    "text" TEXT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP,
    Foreign Key ("recipientId") REFERENCES "users" ("id"),
    Foreign Key ("senderId") REFERENCES "users" ("id")
)