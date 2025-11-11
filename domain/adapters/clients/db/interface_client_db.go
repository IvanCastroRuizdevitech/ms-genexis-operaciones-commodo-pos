package domain_adapters_clients_db

type IClientDB interface {
	Select(query string, arguments []any) ([][]interface{}, error)
	Exec(query string, arguments []any) ([][]interface{}, error)
}
