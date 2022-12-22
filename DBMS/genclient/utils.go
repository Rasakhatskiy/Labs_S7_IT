package genclient

type PostDatabasesJSONRequestBody struct {
	Name string `json:"name"`
}

type TableHeaderJSON struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// PostDatabasesDbNameJSONRequestBody table
type PostDatabasesDbNameJSONRequestBody struct {
	Name    string            `json:"name"`
	Headers []TableHeaderJSON `json:"headers"`
	Values  [][]interface{}   `json:"values"`
}

type GetDatabasesDbNameJoinedTablesParams struct {
	T1 *string
	T2 *string
	C1 *string
	C2 *string
}

type Error struct {
	Code    int
	Message string
}

type PostDatabasesDbNameTableNameJSONRequestBody struct {
}

type DatabaseInfo struct {
}

type Table struct {
}
