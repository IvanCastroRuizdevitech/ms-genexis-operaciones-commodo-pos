package entities

// InitialConfiguration represents the initial POS configuration payload.
// Only commonly-used fields are typed; others are omitted intentionally.
type InitialConfiguration struct {
	Data *InitialConfigurationDataResponse `json:"data,omitempty"`
}

type InitialConfigurationDataResponse struct {
	Equipo             *Equipment        `json:"equipo,omitempty"`
	Jornada            *Journal          `json:"jornada,omitempty"`
	Empresas           []Company         `json:"empresas,omitempty"`
	Promotor           *Promoter         `json:"promotor,omitempty"`
	Parametros         map[string]string `json:"parametros,omitempty"`
	Surtidores         []Dispenser       `json:"surtidores,omitempty"`
	MediosPagos        []PaymentMethod   `json:"medios_pagos,omitempty"`
	TurnoActivo        bool              `json:"turno_activo"`
	SurtidoresDetalles []DispenserDetail `json:"surtidores_detalles,omitempty"`
}

type Equipment struct {
	ID                  int     `json:"id"`
	IP                  *string `json:"ip"`
	Mac                 string  `json:"mac"`
	Port                *string `json:"port"`
	Token               string  `json:"token"`
	Estado              string  `json:"estado"`
	Password            string  `json:"password"`
	UrlFoto             *string `json:"url_foto"`
	LectorIP            *string `json:"lector_ip"`
	Autorizado          string  `json:"autorizado"`
	CreateDate          string  `json:"create_date"`
	CreateUser          int     `json:"create_user"`
	EmpresasID          int     `json:"empresas_id"`
	LectorPort          *int    `json:"lector_port"`
	UpdateDate          *string `json:"update_date"`
	UpdateUser          *int    `json:"update_user"`
	ImpresoraIP         *string `json:"impresora_ip"`
	FactorPrecio        *int    `json:"factor_precio"`
	SerialEquipo        string  `json:"serial_equipo"`
	FactorImporte       *int    `json:"factor_importe"`
	ImpresoraPort       *int    `json:"impresora_port"`
	EquiposTiposID      *int    `json:"equipos_tipos_id"`
	FactorInventario    *int    `json:"factor_inventario"`
	AlmacenamientosID   *int    `json:"almacenamientos_id"`
	EquiposProtocolosID *int    `json:"equipos_protocolos_id"`
}

type Journal struct {
	ID           int         `json:"id"`
	Saldo        float64     `json:"saldo"`
	Atributos    interface{} `json:"atributos"`
	FechaFin     *string     `json:"fecha_fin"`
	EquiposID    int         `json:"equipos_id"`
	PersonasID   int         `json:"personas_id"`
	FechaInicio  string      `json:"fecha_inicio"`
	Sincronizado int         `json:"sincronizado"`
	GrupoJornada int64       `json:"grupo_jornada"`
	SurtidoresID int         `json:"surtidores_id"`
}

type Company struct {
	ID                     int                    `json:"id"`
	Nit                    string                 `json:"nit"`
	Alias                  string                 `json:"alias"`
	Correo                 *string                `json:"correo"`
	Estado                 string                 `json:"estado"`
	Dominio                *string                `json:"dominio"`
	Telefono               *string                `json:"telefono"`
	UrlFoto                *string                `json:"url_foto"`
	Atributos              map[string]interface{} `json:"atributos"`
	Direccion              *string                `json:"direccion"`
	DominioID              *int                   `json:"dominio_id"`
	NegocioID              *int                   `json:"negocio_id"`
	CiudadesID             *int                   `json:"ciudades_id"`
	CreateDate             *string                `json:"create_date"`
	CreateUser             *int                   `json:"create_user"`
	EmpresasID             *int                   `json:"empresas_id"`
	UpdateDate             *string                `json:"update_date"`
	UpdateUser             *int                   `json:"update_user"`
	Localizacion           *string                `json:"localizacion"`
	RazonSocial            *string                `json:"razon_social"`
	CodigoEmpresa          *string                `json:"codigo_empresa"`
	FechaCreacion          *string                `json:"fecha_creacion"`
	ContactoCorreo         *string                `json:"contacto_correo"`
	ContactoNombre         *string                `json:"contacto_nombre"`
	IdTipoEmpresa          *int                   `json:"id_tipo_empresa"`
	ContactoTelefono       *string                `json:"contacto_telefono"`
	EmpresasTiposID        *int                   `json:"empresas_tipos_id"`
	CantidadSucursales     *int                   `json:"cantidad_sucursales"`
	CiudadesDescripcion    *string                `json:"ciudades_descripcion"`
	ProveedorTecnologicoID *int                   `json:"proveedor_tecnologico_id"`
}

type Promoter struct {
	ID                    int     `json:"id"`
	Tag                   *string `json:"tag"`
	Correo                *string `json:"correo"`
	Estado                string  `json:"estado"`
	Genero                *string `json:"genero"`
	Nombre                string  `json:"nombre"`
	Sangre                *string `json:"sangre"`
	Celular               *string `json:"celular"`
	Usuario               *string `json:"usuario"`
	Telefono              *string `json:"telefono"`
	Direccion             *string `json:"direccion"`
	CiudadesID            int     `json:"ciudades_id"`
	CreateDate            *string `json:"create_date"`
	CreateUser            *int    `json:"create_user"`
	EmpresasID            int     `json:"empresas_id"`
	PerfilesID            int     `json:"perfiles_id"`
	UpdateDate            *string `json:"update_date"`
	UpdateUser            *int    `json:"update_user"`
	Sincronizado          int     `json:"sincronizado"`
	SucursalesID          *int    `json:"sucursales_id"`
	Identificacion        string  `json:"identificacion"`
	FechaNacimiento       *string `json:"fecha_nacimiento"`
	TiposIdentificacionID int     `json:"tipos_identificacion_id"`
}

type Dispenser struct {
	ID                            int     `json:"id"`
	IP                            *string `json:"ip"`
	Mac                           string  `json:"mac"`
	Port                          *string `json:"port"`
	Token                         string  `json:"token"`
	Estado                        string  `json:"estado"`
	IslasID                       *int    `json:"islas_id"`
	Surtidor                      *int    `json:"surtidor"`
	LectorIP                      *string `json:"lector_ip"`
	TieneEcho                     *string `json:"tiene_echo"`
	Controlador                   *int    `json:"controlador"`
	CreateDate                    *string `json:"create_date"`
	CreateUser                    *int    `json:"create_user"`
	EmpresasID                    *int    `json:"empresas_id"`
	LectorPort                    *int    `json:"lector_port"`
	UpdateDate                    *string `json:"update_date"`
	UpdateUser                    *int    `json:"update_user"`
	DebugEstado                   *string `json:"debug_estado"`
	DebugTramas                   *string `json:"debug_tramas"`
	ImpresoraIP                   *string `json:"impresora_ip"`
	FactorPrecio                  *int    `json:"factor_precio"`
	ImpresoraPort                 *int    `json:"impresora_port"`
	BytesTotalizador              *int    `json:"bytes_totalizador"`
	FactorInventario              *int    `json:"factor_inventario"`
	SurtidoresTiposID             *int    `json:"surtidores_tipos_id"`
	FactorImporteParcial          *int    `json:"factor_importe_parcial"`
	FactorVolumenParcial          *int    `json:"factor_volumen_parcial"`
	SurtidoresProtocolosID        *int    `json:"surtidores_protocolos_id"`
	FactorPredeterminacionVolumen *int    `json:"factor_predeterminacion_volumen"`
}

type PaymentMethod struct {
	ID                             int         `json:"id"`
	Base64                         *string     `json:"base64"`
	Estado                         string      `json:"estado"`
	Descripcion                    string      `json:"descripcion"`
	MPAtributos                    interface{} `json:"mp_atributos"`
	IntegracionID                  *int        `json:"integracion_id"`
	IDSINcronizado                 *int        `json:"id_sincronizado"`
	CodigoAdquiriente              *string     `json:"codigo_adquiriente"`
	Base64Seleccionado             *string     `json:"base64_seleleccionado"`
	IDMedioPagoRecurso             *string     `json:"id_medio_pago_recurso"`
	IDMedioPagoRecursoSeleccionado *string     `json:"id_medio_pago_recurso_seleccionado"`
}

type DispenserDetail struct {
	ID                            int         `json:"id"`
	Cara                          *int        `json:"cara"`
	Grado                         *int        `json:"grado"`
	Estado                        *int        `json:"estado"`
	Puerto                        *string     `json:"puerto"`
	Bloqueo                       *string     `json:"bloqueo"`
	Conexion                      *int        `json:"conexion"`
	Manguera                      *int        `json:"manguera"`
	Surtidor                      *int        `json:"surtidor"`
	LectorIP                      *string     `json:"lector_ip"`
	BodegasID                     *int        `json:"bodegas_id"`
	FamiliaID                     *int        `json:"familia_id"`
	LectorRFID                    *string     `json:"lector_rfid"`
	ProductosID                   *int        `json:"productos_id"`
	LectorPuerto                  *int        `json:"lector_puerto"`
	SaltoLectura                  *string     `json:"salto_lectura"`
	SurtidoresID                  *int        `json:"surtidores_id"`
	EstadoPublico                 *int        `json:"estado_publico"`
	FamiliaCodigo                 *string     `json:"familia_codigo"`
	MotivoBloqueo                 *string     `json:"motivo_bloqueo"`
	PrecioFamilia                 *float64    `json:"precio_familia"`
	AcumuladoVenta                *int64      `json:"acumulado_venta"`
	ProductoPrecio                *float64    `json:"producto_precio"`
	UltimaConexion                *string     `json:"ultima_conexion"`
	BytesTotalizador              *int        `json:"bytes_totalizador"`
	FamiliaAtributos              interface{} `json:"familia_atributos"`
	AcumuladoCantidad             *int64      `json:"acumulado_cantidad"`
	AcumuladoVentaSurt            *int64      `json:"acumulado_venta_surt"`
	ProductoDescripcion           *string     `json:"producto_descripcion"`
	AcumuladoCantidadSurt         *int64      `json:"acumulado_cantidad_surt"`
	FactorPredeterminacionImporte *int        `json:"factor_predeterminacion_importe"`
}
