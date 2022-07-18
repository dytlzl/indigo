# indigo

indigo is a [WebARENA Indigo API](https://indigo.arena.ne.jp/userapi/) client written in Go.

## Installation

To install `indigo`, you can run the command below.

```sh
go install github.com/dytlzl/indigo@latest
```

And you also have to create a credential file contains a content like below as `$HOME/.indigo.yaml`.

```yaml
credential:
  key: <API Key>
  secret: <API Secret>
```

## Usage

### Listing instances

```console
$ indigo get instances
NAME     STATUS    AGE   IP                OS            PLAN
node00   Running   17d   xxx.xxx.xxx.xxx   Ubuntu20.04   2CR2GB
node01   Running   17d   xxx.xxx.xxx.xxx   Ubuntu20.04   2CR2GB
node02   Running   17d   xxx.xxx.xxx.xxx   Ubuntu20.04   2CR2GB
node03   Running   17d   xxx.xxx.xxx.xxx   Ubuntu20.04   2CR2GB
```

### Listing Firewalls

```console
$ indigo get firewalls
NAME      AGE
default   30d
outbound  15d
```
