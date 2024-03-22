# Arquitetura

MVVM - Model-View-ViewModel. É uma arquitetura de design que visa separar preocupações e facilitar a manutenção, testabilidade e escalabilidade de aplicativos. O MVVM é especialmente popular em desenvolvimento de aplicativos.

# Objetivo

Desenvolver um aplicativo de Gerenciamento de Clinicas Ususando o Flutter.

# Regras iniciais, limite e Análise

- Todo o projeto respeita as regras de Lint escritas no pacote analysis_options.
- Projeto divido em Pacotes.
-- Core

# Entidades
- users.: Usuários do Sistema.
- adm_users.: Administradores do Sistema.
- attendantDeskAssignment.: Terminal que o Atendente ira realizar o atendimento.
- patients.: Pacientes.
- patientInformationForm.: Informações do Paciente.
- painelCheckin.: Informações de senha

# Casos de Uso

Sistema de Gerenciamento Clinico


# Design Patterns
- Core Project: Padrão de Desenvolvimento de Projetos, orientado a Pacotes 
- Repository Pattern: Para acesso a API externa.
- Service Pattern: Para isolar trechos de códigos com outras responsabilidades.
- Dependency Injection: Resolver dependências das classes.
- Store: Guardar e mudar estados.
- State pattern: Padrão que auxilia no gerenciamento estados.
- Adapter: Converter um objeto em outro.
- Result: Trabalhar com retorno Múltiplo.


# Package externos (Comum)
- PostgreSQL.: como banco de dados
- SQCL.: para lidar com nossas consultas ao banco de dados
- Golang.: Migrate para lidar com nossas migrations
- Go Chi.: para criar nossas rotas
- Go playground.: validator será responsável por validar nossos dados de entrada
- Swaggo.: para criar nossa documentação no padrão OpenAPI
- Viper.: para gerenciar nossa variáveis de ambiente
- Docker.: para rodar nosso banco de dados

# Estrutura do projeto
- cmd: Aqui vamos deixar nossos arquivos main.go, responsáveis por iniciar nossa aplicação.
- config: Vamos salvar algumas configs aqui, como envs, logs.
- internal: Aqui é onde vai ficar nossa regra de negócio
- internal/dto: Onde vamos determinar os tipos de dados que vamos permitir entrar na aplicação
- internal/entity: Aqui vamos salvar as entidades da nossa aplicação
- internal/handler: Essa pasta vai ficar nossos arquivos de roteamento (pode chamar de controller se preferir)
- internal/database: Essa pasta é onde vamos salvar tudo que for relacionado ao banco de dados
- internal/database/migrations: Vamos salvar nossas migrations aqui
- internal/database/queries: Onde vai ficar nossas queries sql de consulta ao banco
- internal/database/sqlc: Aqui vai ficar os arquivos gerados automaticamente pelo sqlc
- internal/repository: Aqui onde vai ficar nossa camada de repositórios, não seria preciso colocar essa camada, poderíamos usar diretamente as estruturas do sqlc, mas vamos adicionar mais essa camada, para deixar um pouco mais desacoplado do sqlc.
- internal/service: Por último, nossa camada de service, onde a regra de negócio vai ficar (pode chamar de usecase se preferir).