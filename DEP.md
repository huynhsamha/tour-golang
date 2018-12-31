## DEP - Dependency Pakage Management Go

This package used to manage dependency packages using in project (such as `NPM` or `yarn` in **Node.JS**), with:
+ **Gopkg.lock**
+ **Gopkg.toml**
+ **vendor**  =>  **node_modules** in **Node.JS**

### Installation
Google how to install `dep`

### Find location to create new project - fucking GOPATH
```bash
cd $GOPATH/src/[host]/[username]/[repo]

# Example
cd $GOPATH/src/github.com/huynhsamha/tour-golang
```

### Intilize new Project

```bash
dep init
```

### Add dependencies

```bash
dep ensure -add [github.com]/[username]/[repo]

# Example
dep ensure -add github.com/pkg/errors
dep ensure -add golang.org/x/lint
```

### Update dependencies

```bash
dep ensure -update [github.com]/[username]/[repo]
```

### Ensure used packages

```bash
dep ensure
```

### View used packages
```bash
dep status
```