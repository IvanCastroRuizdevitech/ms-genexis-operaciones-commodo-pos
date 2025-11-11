package constants

const (
	API_PATH string = "/api/v1"

	API_TIPOS_OPERACIONES             = "/tipos/operaciones"
	API_TIPOS_MOVIMIENTOS             = "/tipos/movimientos"
	API_BODEGAS_POR_EDS               = "/bodegas/eds"
	API_PRODUCTOS_TIENDA_EDS          = "/productos/tienda/eds"
	API_MOVIMIENTO_INVENTARIO         = "/movimiento/inventario"
	API_MOVIMIENTO_INVENTARIO_AFECTAR = "/movimiento/inventario/afectar"
	API_ORDEN_DE_COMPRA               = "/orden-compra"
	API_ORDEN_DE_COMPRA_CONFIRMAR     = "/orden-compra/confirmar"
	API_FORMAS_PAGOS                  = "/formas-pagos"
	API_PROVEEDORES                   = "/proveedores"

	ACTION_CREATE = "CREATE"
	ACTION_UPDATE = "UPDATE"
)
