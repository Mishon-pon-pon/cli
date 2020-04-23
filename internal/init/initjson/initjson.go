package initjson

// DefaultJSON ...
var DefaultJSON string = `{
	"DataBase": {
		"connectionString": "server=IP_сервера;user=Имя_Пользователя;password=Пароль;database=База_Данных",
		"db_manager": "mssql"
	},
	"Repository": {
		"remotePath": "полный путь до удаленного репозитория",
		"devPath": "полный путь до dev репозитория",
		"testPath": "полный путь до test репозитория"
 
	}
}
`
