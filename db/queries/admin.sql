-- name: SeedAdmin :one
INSERT INTO admins (full_name,email,password, role)
VALUES('Brian Kibet','briankibet@gmail.com','1234','Registration') RETURNING full_name;