![image](https://hub.steampipe.io/images/plugins/turbot/hackernews-social-graphic.png)

# Hacker News Plugin for Steampipe

Use SQL to query stories, users and other items from [Hacker News](https://news.ycombinator.com).

- **[Get started →](https://hub.steampipe.io/plugins/turbot/hackernews)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/hackernews/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-hackernews/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install hackernews
```

Run a query:

```sql
select * from hackernews_show_hn
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-hackernews.git
cd steampipe-plugin-hackernews
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/hackernews.spc
```

Try it!

```
steampipe query
> .inspect hackernews
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-hackernews/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-hackernews/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Hacker News Plugin](https://github.com/turbot/steampipe-plugin-hackernews/labels/help%20wanted)
