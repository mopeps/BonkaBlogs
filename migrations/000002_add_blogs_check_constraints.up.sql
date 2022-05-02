ALTER TABLE blogs ADD CONSTRAINT title_length_check CHECK ( char_length(text) BETWEEN 1 AND 500);

ALTER TABLE blogs ADD CONSTRAINT tags_length_check CHECK ( array_length(tags, 1) BETWEEN 1 AND 20);
