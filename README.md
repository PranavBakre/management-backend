# backend

[![Test](https://github.com/PranavBakre/management-backend/actions/workflows/test.yml/badge.svg)](https://github.com/PranavBakre/management-backend/actions/workflows/test.yml)
[![Security](https://github.com/PranavBakre/management-backend/actions/workflows/security.yml/badge.svg)](https://github.com/PranavBakre/management-backend/actions/workflows/security.yml)
[![Linter](https://github.com/PranavBakre/management-backend/actions/workflows/linter.yml/badge.svg)](https://github.com/PranavBakre/management-backend/actions/workflows/linter.yml)

## Development

### Start the application 


```bash
go run app.go
```


## Production

### Build and run the docker image

```bash
docker build -t backend .
docker run -d -p 3000:3000 backend
```
