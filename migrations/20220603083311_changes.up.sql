-- create "links" table
CREATE TABLE `links` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `long_uri` text NOT NULL, `accessed_times` integer NOT NULL DEFAULT 0, `created_at` datetime NOT NULL);
-- create index "links_long_uri_key" to table: "links"
CREATE UNIQUE INDEX `links_long_uri_key` ON `links` (`long_uri`);
