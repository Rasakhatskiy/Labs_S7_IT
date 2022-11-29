package database

type TableHeaderJSON struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TableJSONValues struct {
	Name    string            `json:"name"`
	Headers []TableHeaderJSON `json:"headers"`
	Values  [][]interface{}   `json:"values"`
}

type TableJSON struct {
	Name    string            `json:"name"`
	Headers []TableHeaderJSON `json:"headers"`
}

type DatabaseInfoJSON struct {
	Tables []TableJSON `json:"tables"`
}
