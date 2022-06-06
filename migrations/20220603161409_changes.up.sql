-- create "links" table
CREATE TABLE `links` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `short_path` text NULL, `long_uri` text NOT NULL, `accessed_times` integer NOT NULL DEFAULT 0, `created_at` datetime NOT NULL);
-- create index "links_short_path_key" to table: "links"
CREATE UNIQUE INDEX `links_short_path_key` ON `links` (`short_path`);
-- create index "links_long_uri_key" to table: "links"
CREATE UNIQUE INDEX `links_long_uri_key` ON `links` (`long_uri`);
