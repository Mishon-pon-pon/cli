package initjson

// DefaultJSON ...
var DefaultJSON string = `{
	"DataBase": {
		"connectionString": "server=IP_сервера;user id=Имя_Пользователя;password=Пароль;database=База_Данных",
		"db_manager": "mssql"
	},
	"Repository": {
		"remotePath": "полный путь до удаленного репозитория",
		"devPath": "полный путь до dev репозитория",
		"testPath": "полный путь до test репозитория"
 
	}, 
	"Modules": {
		"newsresder": {
			"pathFrom": "путь до модуля",
			"pathIn": "путь до папки куда будем устанавливать"
		}
	}
}
`
