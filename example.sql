-- create Database schooldatabase;
CREATE or REPLACE TABLE lesson (
   student VARCHAR(255),
   room INT
);
INSERT INTO lesson (student, room) VALUES ('John Doe', 404);