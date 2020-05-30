package initjson

// DefaultJSON ...
var DefaultJSON string = `{
	"DataBase": {
		"connectionString": "server=IP_сервера;user id=Имя_Пользователя;password=Пароль;database=База_Данных",
		"db_manager": "mssql"
	},
	"GitRepository": {
		"remotePath": "полный путь до удаленного репозитория",
		"devPath": "полный путь до dev репозитория",
		"testPath": "полный путь до test репозитория"
 
	}, 
	"NpmRepository": {
		"registry": "адрес вашего npm-репозитория"
	},
	"Services": {
		"имя_модуля": {
			"pathFrom": "путь до модуля в node_modules",
			"pathIn": "путь до папки куда будем устанавливать"
		}
	},
	"Modules": {
		"имя_модуля": {
			"pathFrom": "путь до модуля в node_modules",
			"pathIn": "путь до папки куда будем устанавливать"
		}
	}
}
`
