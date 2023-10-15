# Proyecto: Indexador de Correos

## Pasos para la Configuración del Proyecto

### 1. Preparación de los Datos
- Descarga el dataset de correos desde [Enron Email Dataset](http://www.cs.cmu.edu/~enron/) y agrégalo a la carpeta `parsermods/data/maildir/`.

### 2. Configuración de ZincSearch
- Descarga y ejecuta [ZincSearch](https://zincsearch-docs.zinc.dev/) en el puerto `4080`.

### 3. Transformación e Indexación de los Datos
- Navega al directorio `parsermods/parser` y ejecuta el siguiente comando para transformar la data de Enron a formato JSON e indexarla con ZincSearch. Este proceso debería demorar aproximadamente 1 minuto.

### 4. Configuración del Proxy
- Navega al directorio `proxymods` y ejecuta el siguiente comando `go run .` para iniciar el proxy, a través del cual el frontend se comunicará con el indexador. Este comando iniciará un proxy en el puerto `8000`.

### 5. Inicio de la Interfaz de Usuario
- Navega al directorio `frontend` y ejecuta el siguiente comando `npm run serve` para iniciar la interfaz de usuario. Luego, visita `localhost:8080` para realizar consultas.

### 6. Profiling del parser
- Navega al directorio `parsermods/parser` para ver el profiling de CPU y manejo de memoria del parser. Puedes comparar el tiempo de CPU con tiempo real.
- Para revisar resultados puedes correr `go tool pprof -http=:6060 cpu.pprof` y navegar a `localhost:6060`