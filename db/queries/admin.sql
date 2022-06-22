-- name: SeedAdmin :one
INSERT INTO admins (full_name,password, role)
VALUES('Brian Kibet',"1234",'Registration') RETURNING full_name;