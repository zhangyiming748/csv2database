**✅ 最终整理文档**

### MySQL 数据库迁移到 SQLite 完整流程

#### 使用工具
- **工具名称**：`mysql-to-sqlite3`
- **作用**：直接连接 MySQL，把整个数据库（包含多个表）迁移到 SQLite 文件

---

#### 一、最终成功执行的命令（PowerShell）

```powershell
mysql2sqlite `
  --sqlite-file "C:\Users\zhang\Documents\mysql\erp.db" `
  --mysql-database Permissions `
  --mysql-user root `
  --mysql-password "163453" `
  --mysql-host host.docker.internal `
  --mysql-port 3306 `
  --mysql-charset utf8mb4
```

---

#### 二、整个过程中遇到的问题及解决

| 步骤 | 遇到的问题 | 解决方案 |
|------|------------|----------|
| 1 | 权限拒绝：`Access denied for user 'root'@'172.20.0.1'` | MySQL 只允许 `root` 从 `localhost` 登录，Docker 网络下 IP 被识别为 `172.20.0.1` |
| 2 | 执行 `CREATE USER 'root'@'%'` 失败 | 用户已存在，不能重复创建 |
| 3 | PowerShell 命令换行问题 | 使用反引号 `` ` `` 进行换行 |

---

#### 三、解决 MySQL 允许远程访问（核心修复步骤）

在 MySQL 中执行以下 SQL（推荐在容器内或已连接的情况下执行）：

```sql
-- 1. 授予 root 从任意 IP 访问权限
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;

-- 2. 刷新权限
FLUSH PRIVILEGES;

-- 3. （可选）确认 root 用户权限
SELECT user, host FROM mysql.user WHERE user = 'root';
```

**执行完成后**，`root` 用户就可以从 `host.docker.internal`、`127.0.0.1` 等地址正常连接了。

---

#### 四、推荐的最佳实践命令（以后直接用）

```powershell
# 最简洁版本（推荐）
mysql2sqlite `
  --sqlite-file "C:\Users\zhang\Documents\mysql\erp.db" `
  --mysql-database Permissions `
  --mysql-user root `
  --mysql-password "你的密码" `
  --mysql-host host.docker.internal `
  --mysql-port 3306
```

**额外常用参数**（根据需要添加）：
- `--chunk 5000`：大数据量时分批处理
- `--vacuum`：迁移完成后压缩数据库
- `--without-foreign-keys`：跳过外键（SQLite 兼容性更好）

---

迁移已成功，`erp.db` 文件已生成在对应目录。

需要我再给你补充：
- 如何验证迁移是否完整？
- 如何只迁移部分表？
- 如何处理大数据库的优化命令？

随时告诉我！