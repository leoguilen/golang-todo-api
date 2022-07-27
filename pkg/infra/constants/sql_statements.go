package constants

const (
	SelectSingleTodoSql = "SELECT ID AS Id, TITLE AS Title, DESCRIPTION AS Description, LIMIT_DATE AS LimitDate, ASSIGNED_USER AS AssignedUser, STATUS AS Status FROM TODO WHERE ID = ? LIMIT 1;"
	SelectAllTodosSql   = "SELECT ID AS Id, TITLE AS Title, DESCRIPTION AS Description, LIMIT_DATE AS LimitDate, ASSIGNED_USER AS AssignedUser, STATUS AS Status FROM TODO;"
	InsertTodoSql       = "INSERT INTO TODO(ID, TITLE, DESCRIPTION, LIMIT_DATE, ASSIGNED_USER, STATUS) VALUES(?,?,?,?,?,?);"
	UpdateTodoSql       = "UPDATE TODO SET TITLE = ?, DESCRIPTION = ?, LIMIT_DATE = ?, ASSIGNED_USER = ?, STATUS = ? WHERE ID = ?;"
	UpdateStatusSql     = "UPDATE TODO SET STATUS = ? WHERE ID = ?;"
	DeleteTodoSql       = "DELETE FROM TODO WHERE ID = ?;"
)
