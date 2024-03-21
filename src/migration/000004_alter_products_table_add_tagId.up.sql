ALTER TABLE "products"
ADD COLUMN "tagId" INT,
ADD CONSTRAINT "fk_tagId" FOREIGN KEY ("tagId") REFERENCES "tags"("id");