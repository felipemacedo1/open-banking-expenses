# 📌 Gestão de Gastos Integrado ao Open Banking

Este projeto é um **sistema de gestão de gastos** que se integra a APIs bancárias para importar transações financeiras automaticamente. Utiliza **Golang** no backend e um painel em **React/Next.js** para visualização.

---

## ✅ **Checklist de Implementação**

### **1️⃣ Integração com Open Banking (OAuth 2.0)**
- [ ] Implementar **OAuth 2.0** para autenticação segura
- [ ] Criar um serviço em `services/bank_api.go` para autenticar via Open Banking
- [ ] Criar uma tabela `users` para armazenar tokens bancários
- [ ] Criar endpoint `/auth/connect` para iniciar a autenticação
- [ ] Criar endpoint `/auth/callback` para receber tokens

### **2️⃣ Autenticação JWT**
- [ ] Implementar **JWT (JSON Web Token)**
- [ ] Criar serviço `services/auth.go` para gerar e validar tokens JWT
- [ ] Criar endpoints:
    - [ ] `/register` - Cadastro de usuários
    - [ ] `/login` - Autenticação via JWT
- [ ] Criar middleware `middlewares/auth.go` para proteger rotas
- [ ] Proteger endpoint `/transactions` para exigir token válido

### **3️⃣ Criar Dashboard em React/Next.js**
- [ ] Criar uma nova aplicação React:
  ```sh
  npx create-next-app@latest dashboard
  cd dashboard
  npm install axios recharts
  ```
- [ ] Criar **tela de login** e autenticação via JWT
- [ ] Consumir API `/transactions` para exibir gastos
- [ ] Criar **gráficos interativos** com Recharts para análise financeira

### **4️⃣ Implementar Relatórios Financeiros**
- [ ] Criar endpoint `/reports` para gerar relatórios
- [ ] Implementar lógica para:
    - [ ] Total gasto no mês
    - [ ] Categorias mais gastas
    - [ ] Comparação com meses anteriores
- [ ] Criar filtros para relatório: **período, categoria, valor mínimo/máximo**
- [ ] Permitir exportação dos relatórios para **PDF/CSV**
- [ ] Integrar visualização no frontend

---

## 🔧 **Tecnologias Utilizadas**
### **Backend (Golang)**
- Gin (framework web)
- PostgreSQL (banco de dados)
- OAuth2 (autenticação com Open Banking)
- JWT (autenticação segura)
- Docker (ambiente de desenvolvimento)

### **Frontend (React/Next.js)**
- Next.js (framework frontend)
- Recharts (gráficos interativos)
- Axios (requisições HTTP)

---

## 🚀 **Como Rodar o Projeto**

### **1️⃣ Clonar o repositório**
```sh
  git clone https://github.com/seu-usuario/open-banking-expenses.git
  cd open-banking-expenses
```

### **2️⃣ Subir o PostgreSQL com Docker**
```sh
docker run --name pg-open-banking -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=open_banking -p 5432:5432 -d postgres
```

### **3️⃣ Rodar o Backend**
```sh
  go run main.go
```

### **4️⃣ Rodar o Frontend**
```sh
  cd dashboard
  npm install
  npm run dev
```

---

## 📈 **Próximos Passos**
- [ ] Melhorar interface do dashboard
- [ ] Criar testes automatizados
- [ ] Adicionar suporte a múltiplos bancos via Open Finance Brasil

🚀 **Contribua!** Caso tenha sugestões ou melhorias, sinta-se à vontade para abrir uma issue! 😃
