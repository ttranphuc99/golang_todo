package repository

const insertTodoSql = `
INSERT INTO 
todo(title, content, status, owner_id, created_time, updated_time) 
VALUES (?, ?, ?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())
`
const updateTodoSql = `
UPDATE todo
SET
title = ?,
content = ?,
status = ?,
updated_date = UTC_TIMESTAMP(),
WHERE id = ?
`

const getAllTodoByOwnerSql = `
SELECT id, title, content, status, owner_id, created_time, updated_time 
FROM todo WHERE owner_id = ?
`

const getAllTodoSql = `
SELECT id, title, content, status, owner_id, created_time, updated_time 
FROM todo
`

const getAllTodoByOwnerAndStatusSql = `
SELECT * FROM todo WHERE owner_id = ? AND status = ?
`
