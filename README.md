# Concurrente Parcial #

## Dependencias ##

- <https://github.com/wcharczuk/go-chart>

## Uso ##

```bash
$ go run main.go --k=10 --n=800 --mode=both
# Para crear el GIF (Pero tiene mal performance)
$ go run main.go --k=10 --n=800  --mode=chart --gif
```

## Modos ##

- Síncrono (por defecto): `sync`
- Asíncrono: `async`
- Ambos: `both`
- Gráficos: `chart`

## Recomendaciones ##

- No dar un valor al *k* mayor a 50


## Integrantes ##

- Diego Jara
- Mauricio Rivas
- Daniel Aragon