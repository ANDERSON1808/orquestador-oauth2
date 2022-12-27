# Orquestador basico Ouath

> Implementación de un ejemplo de consumo de varios servicios con juramento con respuesta resto

- Carpeta interno: contiene el proyecto del servicio orquestador que conecta clientes con servicio oauth manejando la
  seguridad.
- Carpeta externo: contiene un servidor oauth con datos de un cliente que responde la lista de productos para simular el
  cliente uno o muchos conectado.
  ![](header.png)

## Instalacion

OS X & Linux:

```sh
git clone https://github.com/ANDERSON1808/orquestador-oauth2.git
```

```sh
cd Interno
go mod tidy
```

```sh
cd externo
go mod tidy
```

## Correrlo local

OS X & Linux:

```sh
cd Interno
go run main.go
```

```sh
cd externo
go run main.go
```

## Compilar y correr con docker-compose

```sh
docker-compose up -d  --build
```

## Extension de configuracion y manejo de proyecto

- En la carpeta interno en la ruta Interno/configClients/integraciones.json, se encuentra el archivo de configuracion de
  los clientes que se pueden conectar al servicio principal de productos.
  Este modelo espera un lista de objetos de clientes que se puede estender de manera muy facil, como mejora se puede
  pasar este modelo a una base de datos para mejor manejo y seguridad.
- Modelo integrado.

```sh
[
  {
    "name_company": "company1",
    "auth_url": "http://localhost:9096/authorize",
    "token_url": "http://localhost:9096/token",
    "client_id": "000000",
    "client_secret": "999999",
    "url_product": "http://localhost:9096/products"
  }
]
```
- La carpeta externa contiene un servidor aouth junto con un servicio de productos el cual es consultado por la aplicacion interna manejando la seguridad.
- Enpoint de consulta de productos:
```sh
curl --location --request GET 'http://localhost:8000/products/v0/products' \
--header 'Cookie: go_session_id=NDJiNzU1M2YtYThiNy00MjVmLWJiZTAtMDgxZjUxYTIzMmVk.44608bdaf08585e550c3d5d7581a84906b9f5a73'
```
```sh
http://localhost:8000/products/v0/products
```
- Esta enpoint responde:
```sh
{
    "products": [
        {
            "proveedor": "company1",
            "lista": [
                {
                    "prod_id": "128855",
                    "prod_descrip": "producto en descuento uno"
                },
                {
                    "prod_id": "8484744",
                    "prod_descrip": "producto en descuento dos"
                },
                {
                    "prod_id": "8834342",
                    "prod_descrip": "producto en descuento tres"
                }
            ]
        }
    ]
}
```
