include .env

# Cria nossos arquivos de migration.
create_migration:
	migrate create -ext=sql -dir=internal/database/migrations -seq init
# ext: determina a extensão, vamos usar o sql.
# dir: Aqui fica o diretório onde vai ser criado as nossas migrations.
# seq: Determina a sequência do nome do arquivo da migrations, vamos usar numérico, pode ser usado timestamp.

#Executa nossas migrations up
migrate_up:
	migrate -path=internal/database/migrations -database ${DATABASE_URL} -verbose up
# path: Aqui fica o diretório onde vai ser criado as nossas migrations.
# database: Url do Banco de dados
# verbose: Mostra o que esta sendo feito na criação

# Executa nossas migrations down
migrate_down:
	migrate -path=internal/database/migrations -database ${DATABASE_URL} -verbose down

# 
.PHONY: create_migration migrate_up migrate_down
