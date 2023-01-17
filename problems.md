# Arquiteitura Hexagonal

## Pontos importantes sobre arquitetura
    - crescimento sustentável
    - Software precisa se pagar ao passar do tempo
    - Software deve ser desenhado por você e não pelo seu framework
    - Peças precisam se encaixar e eventualmente substituidas

## Ciclo de vida de muitos projetos
    Fase 1
        Banco de dados
        Cadastros
        Validações
        Servidor web
        Controllers
        Views
        Autenticação
    
    Fase 2
        Regras de negócio
        Criação de APIs
        Consumo de APIs
        Autorização
        Relatórios
        Logs
    
    Fase 3 
        Mais acessos
        Upgrades de Hardware
        Cache
        API parceiros
        Regras parceiros
        Relatórios

    Fase 4
        Mais acessos
        Upgrade Hardware
        BD relatórios
        Comandos
        V2 da API
    
    Fase 5
        Escala horizontal
        Sessões
        Uploads
        Refatoração
        Autoscaling
        CI/CD
    
    Fase 6
        GraphQL
        Bugs Constantes
        Logs? Ops
        Integração CRM
        Migração para React
    
    Fase 7
        Inconsistência CRM
        Containers
        CI/CD
        Memoria
        Logs
        Se livrar do legado

    Fase 8
        Microsserviços
        DB Compartilhado
        Problemas com tracing
        Lentidão
        Custo Elevado

    Fase 9
        Kubernetes
        CI/CD
        Mensageria
        Perda de Mensagens
        Consultorias
    
    Fase 10
        Fim do projeto

## Principais problemas
    visão do futuro
    limites bem definidos
    troca e adição de componentes
    escala
    otimizações frequentes

