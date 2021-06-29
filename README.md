<h1 align="center">
  <img src="https://res.cloudinary.com/gochronicles/image/upload/v1624670043/autobots_vfwwcx.png" width="750px"/><br/>
  Create Microservices MonoRepo in GO/Python
</h1>
<p align="center">Create a new production-ready project with <b>backend</b> (Golang), (Python) by running one CLI command.<br/><br/>Focus on <b>writing</b> code and <b>thinking</b> of business-logic! The CLI will take care of the rest.</p>

<p align="center"><img src="https://img.shields.io/badge/version-v1.0.0-blue?style=for-the-badge&logo=none" alt="cli version" /></a>&nbsp;<img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;</a>&nbsp;</a>&nbsp;<img src="https://img.shields.io/badge/license-MIT-red?style=for-the-badge&logo=none" alt="license" /></p>

## ⚡️ Quick start

### Installation

1. **Mac OS**

```bash
brew tap gochronicles/cli https://github.com/gochronicles/cli
```
```bash
brew install autobots
```

2. **Linux**

```bash
curl https://github.com/gochronicles/cli/releases/download/v1.0.6/autobots_1.0.6_Linux_x86_64.tar.gz -o autobots.tar.gz
tar xvzf autobots.tar.gz
sudo cp autobots /usr/bin/autobots
```

3. **Build from Source**

Make sure you have [golang 1.16+](https://golang.org/doc/install) installed

```bash
git clone https://github.com/gochronicles/cli # clone latest code base
cd cli
go install gochronicles/cmd/autobots # installs required dependencies
go build -o autobots cmd/autobots/main.go # builds autobots binary file
./autobots roll # to run the cli
sudo cp autobots /usr/bin/autobots # optional: if you wish to have autobots cli to be available globally 
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
