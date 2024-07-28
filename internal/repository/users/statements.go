package users

const (
	tableName  = " users "
	allColumns = ` id, email, name, surname, password, username, ` +
		`phone, role, city, country, created_at, updated_at, rating `
)

// insert
const (
	insertNewUserStmt = `
		INSERT INTO ` + tableName + `(` + allColumns + `) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW(), $11);`
)

// select
const (
	allPlusTableNameStmt = "SELECT " + allColumns + " FROM " + tableName
)
const (
	selectByIDStmt = allPlusTableNameStmt +
		" WHERE id = $1;"

	selectByEmailStmt = allPlusTableNameStmt +
		" WHERE email = $1;"

	selectByPhoneStmt = allPlusTableNameStmt +
		` WHERE phone = $1;`

	selectByUsernameStmt = allPlusTableNameStmt +
		" WHERE username = $1;"
)

// update
const (
	updatePhoneStmt = `
		UPDATE ` + tableName + `
		SET phone=$1, update_at = NOW() WHERE id = $2;`
	updateByEmailStmt = `
		UPDATE ` + tableName + `
		SET email=$1, update_at = NOW() WHERE id = $2;`
	updateByNicknameStmt = `
		UPDATE ` + tableName + `
		SET nickname=$1, update_at = NOW() WHERE id = $2;`
)

// delete
const (
	deleteByIDStmt = `DELETE FROM ` + tableName + ` WHERE id = $1;`
)
