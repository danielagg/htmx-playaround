CREATE SCHEMA htmx_playaround

CREATE TABLE htmx_playaround.Quotes (
    Id SERIAL PRIMARY KEY,
    Quote TEXT,
    Author TEXT,
    Ratings INTEGER[],
    CreatedOn TIMESTAMPTZ DEFAULT current_timestamp
);

INSERT INTO htmx_playaround.Quotes (Quote, Author, Ratings)
VALUES (
    'Life is what happens when you''re busy making other plans.',
    'John Lennon',
    ARRAY[5, 4, 5]
);