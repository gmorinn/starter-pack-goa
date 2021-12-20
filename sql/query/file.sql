-- name: CreateFile :one
INSERT INTO files (name, url, mime, size) 
VALUES ($1, $2, $3, $4)
RETURNING id, name, url, mime, size;

-- name: GetFileByURL :one
SELECT * FROM files
WHERE deleted_at IS NULL
AND url = $1;

-- name: DeleteFile :exec
UPDATE files
SET deleted_at = NOW()
WHERE id = $1;