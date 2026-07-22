ALTER TABLE "users" ADD COLUMN "created_by" BIGINT REFERENCES "users"("id");
