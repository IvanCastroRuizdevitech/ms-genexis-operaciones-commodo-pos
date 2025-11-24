package entities

type PendingTransmission struct {
	IDTransmision int    `json:"id_transmision"`
	URL           string `json:"url"`
	Request       string `json:"request"`
	Method        string `json:"method"`
	IDTipo        int    `json:"id_tipo"`
}
