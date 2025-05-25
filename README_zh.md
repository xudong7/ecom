# Go 电商后端系统

基于Go构建的现代化电商后端服务，具备用户管理、商品目录和订单处理功能。

灵感来源于 [Complete Backend Engineering Course in Go](https://www.youtube.com/watch?v=7VLmLOiQ3ck&list=PLYEESps429vrFV0yiN_MCaDPhnYb0qRxK)

## 系统要求

- Go 1.23.5 或更高版本
- MySQL 数据库（生产环境）或任何 SQL 数据库（开发环境）
- Go 编程语言基础知识
- Web 开发概念和 RESTful API 基础

## 项目配置

### 1. 克隆和安装依赖

```bash
# 克隆仓库
git clone https://github.com/xudong7/ecom.git
cd ecom

# 安装 Go 依赖
go mod tidy
```

### 2. 环境配置

在项目根目录创建 `.env` 文件，包含以下变量：

```env
# 数据库配置
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_ADDRESS=localhost:3306
DB_NAME=ecom_db

# 服务器配置
PUBLIC_HOST=localhost
PORT=8080
```

### 3. 数据库设置

```bash
# 安装 golang-migrate 工具（如果尚未安装）
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 创建新的迁移文件（示例）
make migration add-new-table-name

# 运行所有待执行的迁移
make migrate-up

# 回滚最后一次迁移（如果需要）
make migrate-down
```

## 运行项目

### 开发环境

```bash
# 清理、整理依赖、构建并运行
make run

# 或者运行单独的命令：
make clean    # 清理构建产物
make tidy     # 整理 Go 模块
make build    # 构建应用程序
make test     # 运行测试
```

服务器将在 `http://localhost:8080` 启动

### 可用的 Make 命令

| 命令 | 描述 |
|------|------|
| `make run` | 清理、整理、构建并运行服务器 |
| `make build` | 构建应用程序二进制文件 |
| `make test` | 运行所有测试 |
| `make tidy` | 整理 Go 模块依赖 |
| `make clean` | 删除构建产物 |
| `make migration <name>` | 创建新的数据库迁移 |
| `make migrate-up` | 应用所有待执行的迁移 |
| `make migrate-down` | 回滚最后一次迁移 |

## 项目结构

```
├── cmd/
│   ├── main.go                 # 应用程序入口点
│   ├── api/                    # API 服务器和路由
│   └── migrate/                # 数据库迁移工具
│       └── migrations/         # SQL 迁移文件
├── config/                     # 配置管理
├── db/                         # 数据库连接和设置
├── service/                    # 业务逻辑服务
│   ├── auth/                   # 认证服务
│   └── user/                   # 用户管理服务
├── types/                      # 类型定义和模型
├── utils/                      # 工具函数
├── go.mod                      # Go 模块依赖
├── go.sum                      # 依赖校验和
├── Makefile                    # 构建和开发命令
└── README.md                   # 项目文档
```

## API 端点

应用程序提供以下 RESTful API 端点：

- **用户管理**：注册、认证、个人资料管理
- **商品目录**：商品列表、详情、搜索
- **订单处理**：购物车管理、订单创建、订单历史

## 数据库结构

项目包含以下主要数据表：

- `users` - 用户账户和认证信息
- `products` - 商品目录
- `orders` - 订单信息
- `order_items` - 订单中的具体商品项

## 开发指南

### 创建新的迁移

```bash
# 创建用于添加表的新迁移
make migration add-table-name

# 创建用于修改现有数据的新迁移
make migration update-table-name

# 始终创建上行和下行迁移以确保可逆性
```

### 测试

```bash
# 运行所有测试
make test

# 带详细输出运行测试
go test -v ./...

# 运行特定包的测试
go test -v ./service/user
```

## 涉及的技术要点

本项目演示了：

1. 使用 Gorilla Mux 搭建 Go Web 服务器
2. 正确的路由处理 HTTP 请求和响应
3. 使用 MySQL 数据库和迁移
4. 实现用户认证和密码哈希
5. 构建 RESTful API 端点
6. 数据库连接管理和 SQL 操作
7. 项目结构和组织最佳实践
8. Go Web 应用程序测试
9. 环境配置管理
10. 使用 Makefile 进行构建自动化

## 故障排除

### 常见问题

**迁移工具未找到：**

```bash
# 确保已安装 golang-migrate
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 检查 $GOPATH/bin 是否在 PATH 中
echo $PATH | grep $GOPATH/bin
```

**数据库连接问题：**

- 验证 MySQL 正在运行
- 检查 `.env` 文件中的数据库凭据
- 确保在运行迁移之前数据库已存在

**构建问题：**

```bash
# 清理并重新构建
make clean
make build
```

## 参考资源

- [Go 官方文档](https://golang.org/doc/)
- [Gorilla Mux 路由器](https://github.com/gorilla/mux)
- [Golang Migrate](https://github.com/golang-migrate/migrate)
- [Go MySQL 驱动](https://github.com/go-sql-driver/mysql)
- [Go Web 示例](https://gowebexamples.com/)

## 贡献指南

1. Fork 仓库
2. 创建功能分支
3. 进行更改
4. 为新功能添加测试
5. 运行测试并确保通过
6. 提交 Pull Request

## 许可证

本项目仅用于教育目的。