# Grafana Ops

Uma ferramenta escrita em Go para realizar backup dos componentes do Grafana
(dashboards, pastas, fontes de dados, etc.).

> ⚠️ Projeto em desenvolvimento – use com cuidado em produção.

## Instalação

A maneira mais simples de instalar é baixar o *release* mais recente da
página de Releases do GitHub e extrair o binário `gops`:

```sh
# Linux/amd64 como exemplo
curl -LO https://github.com/aristidesneto/grafana-ops/releases/latest/download/gops_linux_amd64.tar.gz
tar -xzf gops_linux_amd64.tar.gz
chmod +x gops
mv gops /usr/local/bin/        # ou outro diretório no PATH
```

Para outras plataformas, substitua o nome do arquivo conforme apropriado
(`gops_darwin_amd64`, `gops_windows_amd64.exe`, etc.).

> 🧰 **Alternativa de desenvolvimento**: Se você preferir compilar localmente, o
> repositório contém um `Makefile` que usa `goreleaser`. Execute `make deps &&
> make build` e o binário será colocado em `dist/`.

## Configuração e uso

O utilitário aceita opções de várias fontes, na seguinte ordem de precedência:
1. Flags de linha de comando
2. Variáveis de ambiente
3. Arquivo de configuração YAML

### Exemplo rápido – flags

```sh
gops \
  --grafana-url https://grafana.example.com \
  --grafana-token "mytoken" \
  --output ./backups \
  --loglevel debug
```

### Usando variáveis de ambiente

```sh
export GRAFANA_URL=https://grafana.example.com
export GRAFANA_TOKEN=mytoken
export OUTPUT=./backups
export LOGLEVEL=info
# executar sem flags
gops
```

### Arquivo de configuração

O arquivo é YAML e pode ser passado com `--config` ou `-c`. Exemplo de
`config.yaml`:

```yaml
# config.yaml
grafana-url: https://grafana.example.com
grafana-token: "mytoken"
output: ./backups
loglevel: info
```

```sh
gops --config /path/to/config.yaml
```

Se preferir, o diretório padrão buscado é `./` e `~/.gops` com o
nome `config.yaml`.

## Exemplos de uso

- Backup completo usando flags:
```sh
gops save --grafana-url https://grafana.local --grafana-token abc123 \  
      --output /var/backups/grafana
```

## Contribuindo

Sinta-se à vontade para enviar pull requests, reportar issues ou sugerir
melhorias.
