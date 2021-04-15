CREATE TABLE points(
    id varchar(128) PRIMARY KEY,
    type int,
    data text,
    g GEOGRAPHY(POINT, 4326)
);

CREATE INDEX points_g_idx ON points USING GIST (g);