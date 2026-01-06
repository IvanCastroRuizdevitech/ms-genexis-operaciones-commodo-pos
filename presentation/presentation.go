package presentation

import (
	container_comanda "ms-genexis-pos-operaciones/context/comanda/presentation/container"
)

func Run() {
	container_comanda.ResolveUpdateComandaStatusContainer()
}
