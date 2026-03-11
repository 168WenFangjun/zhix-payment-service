<div align="center">

```
██████╗  █████╗ ██╗   ██╗
██╔══██╗██╔══██╗╚██╗ ██╔╝
██████╔╝███████║ ╚████╔╝ 
██╔═══╝ ██╔══██║  ╚██╔╝  
██║     ██║  ██║   ██║   
╚═╝     ╚═╝  ╚═╝   ╚═╝   
```

### 极志社区 · 支付服务

**一触即付，安全可信。**

---

[![作者主页](https://img.shields.io/badge/🔥_作者是谁？点进来就知道了-→-FF3B30?style=for-the-badge)](https://www.macfans.app/pixel-forge-studio/)
&nbsp;
[![Author](https://img.shields.io/badge/🌐_Who_built_this%3F_Find_out_→-6C63FF?style=for-the-badge)](https://www.macfans.app/pixel-forge-studio/)

---

</div>

---

## 🇨🇳 中文版

<br>

<div align="center">

**支付这件事，要么不做，要做就做对。**

Go 语言驱动，Apple Pay 原生集成，独立微服务，为安全而生。

</div>

<br>

### 🧱 技术栈

<div align="center">

![Go](https://img.shields.io/badge/Go_1.21-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL_15-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Apple Pay](https://img.shields.io/badge/Apple_Pay-000000?style=for-the-badge&logo=applepay&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![AWS ECS](https://img.shields.io/badge/AWS_ECS-FF9900?style=for-the-badge&logo=amazonaws&logoColor=white)

</div>

<br>

### 📁 项目结构

```
payment-service/
├── config/
│   └── database.go       # 数据库连接
├── controllers/
│   └── payment.go        # 支付控制器
├── middleware/
│   └── auth.go           # JWT 认证
├── models/
│   └── payment.go        # 数据模型
├── routes/
│   └── routes.go         # 路由注册
├── .env.example
├── Dockerfile
└── main.go
```

<br>

### ⚡ 三步跑起来

```bash
go mod download
cp .env.example .env
go run main.go
```

`.env` 配置：

```env
# 数据库（必须设置）
DATABASE_URL=host=localhost user=<db_user> password=<db_password> dbname=zhix port=5432 sslmode=disable TimeZone=Asia/Shanghai

# JWT 密钥（必须与 backend 完全相同）
JWT_SECRET=<same_as_backend_JWT_SECRET>

# CORS 允许的前端域名
ALLOWED_ORIGIN=http://localhost:3000

# 运行模式（生产环境设为 release）
GIN_MODE=debug

# Apple Pay 配置
APPLE_MERCHANT_ID=merchant.com.zhix.club
APPLE_PAY_CERT_PATH=/path/to/merchant_id.pem
APPLE_PAY_KEY_PATH=/path/to/merchant_id.key
```

> ❗️ `JWT_SECRET` 必须与 backend 服务保持一致，否则支付请求会返回 `Invalid token`。

> 启动后访问 → `http://localhost:8081`

<br>

### 🎯 API 端点

```
POST   /api/payment/apple-pay/session      # 创建支付会话（公开）
GET    /api/payment/status/:orderId        # 查询支付状态（公开）
POST   /api/payment/apple-pay/process      # 处理支付（需认证）
POST   /api/payment/refund/:orderId        # 发起退款（需认证）
GET    /health                             # 健康检查
```

<br>

### 🔐 安全机制

```
✦ JWT Token 认证，与主服务共享密钥
✦ Apple Pay 证书双向验证
✦ HTTPS 全程加密
✦ CORS 跨域保护
```

<br>

### 🚀 部署

通过 GitHub Actions 自动构建镜像，推送至 AWS ECR，部署到 ECS Fargate。

<br>

---

<div align="center">

**这个项目背后的人，比你想象的更有意思。**

[![👀 去看看作者在搞什么](https://img.shields.io/badge/👀_去看看作者在搞什么_→_macfans.app-FF3B30?style=for-the-badge)](https://www.macfans.app/pixel-forge-studio/)

</div>

---

<br>

## 🇺🇸 English Version

<br>

<div align="center">

**Payments done right. No shortcuts.**

Go · Apple Pay · Microservice · Built for trust.

</div>

<br>

### 🧱 Tech Stack

<div align="center">

![Go](https://img.shields.io/badge/Go_1.21-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL_15-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Apple Pay](https://img.shields.io/badge/Apple_Pay-000000?style=for-the-badge&logo=applepay&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![AWS ECS](https://img.shields.io/badge/AWS_ECS-FF9900?style=for-the-badge&logo=amazonaws&logoColor=white)

</div>

<br>

### 📁 Project Structure

```
payment-service/
├── config/
│   └── database.go       # DB connection
├── controllers/
│   └── payment.go        # Payment handlers
├── middleware/
│   └── auth.go           # JWT auth
├── models/
│   └── payment.go        # Data models
├── routes/
│   └── routes.go         # Route registration
├── .env.example
├── Dockerfile
└── main.go
```

<br>

### ⚡ Up in 3 Steps

```bash
go mod download
cp .env.example .env
go run main.go
```

`.env` config:

```env
# Database (required)
DATABASE_URL=host=localhost user=<db_user> password=<db_password> dbname=zhix port=5432 sslmode=disable TimeZone=Asia/Shanghai

# JWT secret (must be identical to backend's JWT_SECRET)
JWT_SECRET=<same_as_backend_JWT_SECRET>

# CORS allowed origin
ALLOWED_ORIGIN=http://localhost:3000

# Run mode (set to release in production)
GIN_MODE=debug

# Apple Pay
APPLE_MERCHANT_ID=merchant.com.zhix.club
APPLE_PAY_CERT_PATH=/path/to/merchant_id.pem
APPLE_PAY_KEY_PATH=/path/to/merchant_id.key
```

> ⚠️ `JWT_SECRET` must be **identical** to the backend service. A mismatch will cause all payment requests to fail with `Invalid token`.

> Runs at → `http://localhost:8081`

<br>

### 🎯 API Endpoints

```
POST   /api/payment/apple-pay/session      # Create payment session (public)
GET    /api/payment/status/:orderId        # Query payment status (public)
POST   /api/payment/apple-pay/process      # Process payment (auth required)
POST   /api/payment/refund/:orderId        # Refund (auth required)
GET    /health                             # Health check
```

<br>

### 🔐 Security

```
✦ JWT Token auth — shared secret with main backend
✦ Apple Pay certificate mutual verification
✦ Full HTTPS encryption
✦ CORS protection
```

<br>

### 🚀 Deployment

Auto-built via GitHub Actions, pushed to AWS ECR, deployed to ECS Fargate.

<br>

---

<div align="center">

**The person behind this project is worth knowing.**

[![🔗 Step Into the Author's World →](https://img.shields.io/badge/🔗_Step_Into_the_Author's_World_→-6C63FF?style=for-the-badge)](https://www.macfans.app/pixel-forge-studio/)

</div>

---

<div align="center">
<br>

`// made with focus · ZhiX Team`

</div>
