CREATE TABLE "testimonial"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "fullName" VARCHAR NOT NULL,
    "role" VARCHAR DEFAULT 'customer coffee shop',
    "feedback" TEXT,
    "rate" INT NOT NULL,
    "image" VARCHAR,
    "createdAt" TIMESTAMP DEFAULT now(),
    "updatedAt" TIMESTAMP,
    CHECK ("rate" > 0 AND "rate" < 6)
)