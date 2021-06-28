<h1 align="center">
  <img src="https://res.cloudinary.com/gochronicles/image/upload/v1624670043/autobots_vfwwcx.png" width="750px"/><br/>
  Create Microservices MonoRepo in GO/Python
</h1>
<p align="center">Create a new production-ready project with <b>backend</b> (Golang), (Python) by running one CLI command.<br/><br/>Focus on <b>writing</b> code and <b>thinking</b> of business-logic! The CLI will take care of the rest.</p>

<p align="center"><img src="https://img.shields.io/badge/version-v1.0.0-blue?style=for-the-badge&logo=none" alt="cli version" /></a>&nbsp;<img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;</a>&nbsp;</a>&nbsp;<img src="https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none" alt="license" /></p>

## ⚡️ Quick start

Installation is done by using [`brew install`] command :

```bash
brew tap gochronicles/cli https://github.com/gochronicles/cli
```
```bash
brew install autobots
```

Let's create a new project via **interactive console UI** (or **CUI** for short) in current folder:

```bash
autobots roll
```

That's all you need to know to start! 🎉

## ⚙️ Commands & Options

### `roll`

CLI command for creating a new project with the interactive console UI.

```bash
autobots roll [OPTION]
```

| Option | Description                                              | Type   | Default | Required? |
| ------ | -------------------------------------------------------- | ------ | ------- | --------- |
| `-t`   | Enables to define custom backend and frontend templates. | `bool` | `false` | No        |

## 📝 Production-ready project templates

### Backend

- Backend template with [Fiber](https://github.com/gofiber/fiber):
  - [`fiber`](https://github.com/gochronicles/monorepo-fiber-neo4j) — complex REST API with CRUD and DB.

### Database

- Autobot configures Neo4j as a database by default for Go project and postgres for python [Neo4j](https://neo4j.com/):
  - `neo4j` — configured neo4j container with apply migrations for backend.

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `Autobots CLI`:

- Add a [GitHub Star] to the project.

## ⚠️ License

`Autobots CLI` is free and open-source software licensed under the [MIT License]
