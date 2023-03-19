-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email,
  phone_number
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  email = COALESCE(sqlc.narg(email), email),
  phone_number= COALESCE(sqlc.narg(phone_number), phone_number)
WHERE
  username = sqlc.arg(username)
RETURNING *;
