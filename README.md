# Gator 🐊

Gator is a command-line RSS feed aggregator written in Go. It allows users to follow RSS feeds and browse new posts fetched from those feeds — all stored in PostgreSQL.

---

## 📦 Requirements

- Go 1.25+
- PostgreSQL (running locally or remotely)
- [Goose](https://github.com/pressly/goose) for managing database migrations

---

## 🧩 Installation

To install the Gator CLI:

```bash
go install github.com/vitlobo/gator@latest
```

---

## ⚙️ Configuration

Gator uses a config file located at:

```
~/.config/gator/.gatorconfig.json
```

Example:

```json
{
  "version": 1,
  "current_user_name": "username",
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

---

## 🗃️ Database Setup

Run migrations with Goose:

```bash
goose -dir ./sql/schema postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" up
```

```bash
goose -dir ./sql/schema postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" down
```

---

## 🚀 Usage

### Common Commands

| Command                | Description                           | Example                                                 |
| ---------------------- | ------------------------------------- | ------------------------------------------------------- |
| `addfeed <name> <url>` | Add a new RSS feed                    | `gator addfeed TechCrunch https://techcrunch.com/feed/` |
| `agg <duration>`       | Start feed aggregation loop           | `gator agg 45s`                                         |
| `browse [limit]`       | Browse recent posts (default limit 2) | `gator browse 10`                                       |
| `feeds`                | List all feeds being tracked          | `gator feeds`                                           |
| `follow <url>`         | Follow an existing feed               | `gator follow https://techcrunch.com/feed/`             |
| `following`            | List all feeds you follow             | `gator following`                                       |
| `login <name>`         | Log in as an existing user            | `gator login alice`                                     |
| `register <name>`      | Register and log in as a new user     | `gator register alice`                                  |
| `reset`                | Wipe the database                     | `gator reset`                                           |
| `unfollow <url>`       | Unfollow a feed                       | `gator unfollow https://techcrunch.com/feed/`           |
| `users`                | Show all users                        | `gator users`                                           |

---

## 🧠 Example Workflow

```bash
gator register username
gator addfeed TechCrunch https://techcrunch.com/feed/
gator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
gator follow https://techcrunch.com/feed/
gator agg 45s
# (In a separate terminal)
gator browse 10
```

---

## 🧩 Notes

- Use `./gator reset` to reset DB.
- Use `./gator agg 30s` in one terminal and `./gator browse` in another to simulate live aggregation.
- PostgreSQL must be running and accessible to the configured user.

---

## 📁 Project Structure

```
.
├── cmd/                # CLI commands
├── internal/           # Internal packages (core, database, config, etc.)
├── sql/                # SQL schema and queries
├── main.go             # Entry point
├── go.mod / go.sum     # Dependencies
└── README.md
```

---

## 🧰 Tech Stack

- Go 1.25
- PostgreSQL
- Goose (DB migrations)
- `github.com/fatih/color` for colored CLI output

---
