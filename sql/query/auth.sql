-- name: LoginUser :one
SELECT id, firstname, lastname, email, role FROM users
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL;

-- name: Signup :one
INSERT INTO users (firstname, lastname, email, password, phone, birthday) 
VALUES ($1, $2, $3, crypt($4, gen_salt('bf')), $5, $6)
RETURNING *;

-- name: ExistUserByEmail :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = $1
    AND deleted_at IS NULL
);

-- name: UpdateUserPassword :exec
UPDATE users
SET password = crypt($2, gen_salt('bf')), updated_at = NOW()
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1
AND deleted_at IS NULL;

-- name: CheckEmailExist :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = $1
    AND deleted_at IS NULL
);

-- name: SignProvider :one
INSERT INTO users (firstname, lastname, email, password, phone, birthday, firebase_id_token, firebase_uid, firebase_provider) 
VALUES ($1, $2, $3, crypt($4, gen_salt('bf')), $5, $6, $7, $8, $9)
RETURNING *;

-- name: ExistGetUserByFireBaseUid :one
SELECT EXISTS(
	SELECT * FROM users
	WHERE deleted_at IS NULL
	AND firebase_uid = sqlc.arg('firebase_uid')
);

-- name: GetUserByFireBaseUid :one
SELECT * FROM users
WHERE deleted_at IS NULL
AND firebase_uid = $1;


-- name: UpdateUserProvider :exec
UPDATE users
SET firebase_id_token = $2, firebase_uid = $3, firebase_provider = $4, updated_at = NOW()
WHERE id = $1
RETURNING *;