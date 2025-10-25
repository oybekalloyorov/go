
CREATE TABLE IF NOT EXISTS "instagram_users" (
    "id" SERIAL PRIMARY KEY,
    "full_name" VARCHAR(64),
    "username" VARCHAR(64) UNIQUE,
    "birth_of_year" INT,
    "bio" TEXT,
    "created_at" TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "post" (
    "id" SERIAL PRIMARY KEY,
    "description" TEXT,
    "title" TEXT,
    "likes_count" INT,
    "created_by" INT REFERENCES "instagram_users"("id"),
    "created_at" TIMESTAMP default NOW()
);
