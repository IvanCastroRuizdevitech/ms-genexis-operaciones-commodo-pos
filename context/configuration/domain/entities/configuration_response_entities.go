package entities

type Config struct {
	TipoAutorizacion         string `json:"tipo_autorizacion"`
	SolicitarLecturasTanques string `json:"solicitar_lecturas_tanques"`
	MontoMinimoFe            string `json:"MONTO_MINIMO_FE"`
	ObligatorioFe            string `json:"OBLIGATORIO_FE"`
	VersionPos               string `json:"version_pos"`
	PosPrincipal             string `json:"POS_PRINCIPAL"`
}
