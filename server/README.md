hcmcafe server
===

Server
---

```bash
$ make serve
```

Test
---

### Unit Test

```bash
$ go test ./...
```

### HTTP Test

* required [httpie](https://github.com/jakubroztocil/httpie)
* Serve at first
* Use `fixture/*.json`

For example,

```bash
$ http http://localhost:9000/rpc @fixture/Admin.Create.json
```
