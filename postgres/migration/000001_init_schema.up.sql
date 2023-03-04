CREATE TABLE "podcasts" (
                            "id" serial PRIMARY KEY NOT NULL,
                            "name" varchar(255) NOT NULL DEFAULT '',
                            "last_updated" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "episodes" (
                            "id" serial PRIMARY KEY NOT NULL,
                            "name" varchar(255) NOT NULL DEFAULT '',
                            "number_series" int NOT NULL DEFAULT 0,
                            "number_overall" int NOT NULL DEFAULT 0,
                            "release_date" timestamptz NOT NULL DEFAULT (now()),
                            "description" varchar(255) NOT NULL DEFAULT '',
                            "body" text NOT NULL DEFAULT '',
                            "transcript_url" varchar(255) NOT NULL DEFAULT '',
                            "podcast_id" int NOT NULL DEFAULT 0,
                            "series_id" int NOT NULL DEFAULT 0,
                            "last_updated" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "series" (
                          "id" serial PRIMARY KEY NOT NULL,
                          "name" varchar(255) NOT NULL DEFAULT '',
                          "podcast_id" int NOT NULL DEFAULT 0,
                          "last_updated" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "podcasts" ("name");

CREATE UNIQUE INDEX ON "series" ("name");

ALTER TABLE "episodes" ADD FOREIGN KEY ("podcast_id") REFERENCES "podcasts" ("id");

ALTER TABLE "episodes" ADD FOREIGN KEY ("series_id") REFERENCES "series" ("id");

ALTER TABLE "series" ADD FOREIGN KEY ("podcast_id") REFERENCES "podcasts" ("id");

-- Add a generated column that contains the search document
ALTER TABLE ONLY episodes
    ADD COLUMN fts_doc_en tsvector GENERATED ALWAYS AS (
        to_tsvector('english', body || ' ' || description)
        )
        stored;

-- Create a GIN index to make searches faster
CREATE INDEX episodes_fts_doc_en_idx ON episodes USING GIN (fts_doc_en);