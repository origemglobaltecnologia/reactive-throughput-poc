# Reactive Throughput PoC - Termux Edition

PoC de alta performance testando os limites de processamento assíncrono em ambiente mobile.

## 🚀 Performance
- **RPS Estável:** ~1.300 req/s
- **Stack:** Java 21 (Virtual Threads), Spring WebFlux, Go, RabbitMQ.

## 🏗️ Estrutura
- `/reactive`: Servidor Spring Boot (Ingestor + Consumidor).
- `/load-generator`: Canhão de carga em Go.

## 📊 Monitoramento
1. Spring Actuator exposto em `/actuator/prometheus`.
2. Em breve: Painel unificado com Prometheus local.
