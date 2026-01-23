# 🧪 Laboratorio CI/CD con Go, Gin y Render

Este proyecto demuestra un flujo completo de **Integración Continua (CI)** y **Despliegue Continuo (CD)** usando **Go (Golang)**, el framework **Gin**, **GitHub Actions** y **Render**.

El objetivo es garantizar que el despliegue a producción ocurra **solo si las pruebas unitarias pasan correctamente**.

---

## 🚀 Tecnologías utilizadas

- **Go 1.23+**
- **Gin Web Framework**
- **GitHub Actions** — CI pipeline (tests + build)
- **Render** — Hosting y despliegue automatizado vía *Deploy Hook*
- **Docker** — Contenedor de la aplicación para despliegue reproducible

---

## ⚙️ Flujo general

```mermaid
graph TD
A[Commit / Push a main] --> B[GitHub Actions CI]
B -->|Ejecuta tests con Go| C{Tests pasan?}
C -->|❌ No| D[Cancelado ❌ No hay despliegue]
C -->|✅ Sí| E[Render Deploy Hook 🔁]
E --> F[Render construye imagen Docker]
F --> G[Render despliega nueva versión 🚀]
