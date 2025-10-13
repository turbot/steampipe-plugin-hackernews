## v1.2.0 [2025-10-13]

_Dependencies_

- Recompiled plugin with Go version `1.24`. ([#72](https://github.com/turbot/steampipe-plugin-hackernews/pull/72))
- Recompiled plugin with [steampipe-plugin-sdk v5.13.1](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5131-2025-09-25) that addresses critical and high vulnerabilities in dependent packages. ([#73](https://github.com/turbot/steampipe-plugin-hackernews/pull/73))

## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#68](https://github.com/turbot/steampipe-plugin-hackernews/pull/68))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#68](https://github.com/turbot/steampipe-plugin-hackernews/pull/68))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#65](https://github.com/turbot/steampipe-plugin-hackernews/pull/65))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#65](https://github.com/turbot/steampipe-plugin-hackernews/pull/65))

## v0.8.1 [2023-12-14]

_Bug fixes_

- Fixed the plugin to correctly return data without any stray characters when using SQLite.

## v0.8.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#52](https://github.com/turbot/steampipe-plugin-hackernews/pull/52))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#52](https://github.com/turbot/steampipe-plugin-hackernews/pull/52))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-hackernews/blob/main/docs/LICENSE). ([#52](https://github.com/turbot/steampipe-plugin-hackernews/pull/52))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#51](https://github.com/turbot/steampipe-plugin-hackernews/pull/51))

## v0.7.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#43](https://github.com/turbot/steampipe-plugin-hackernews/pull/43))

## v0.7.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#41](https://github.com/turbot/steampipe-plugin-hackernews/pull/41))
- Recompiled plugin with Go version `1.21`. ([#41](https://github.com/turbot/steampipe-plugin-hackernews/pull/41))

## v0.6.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#34](https://github.com/turbot/steampipe-plugin-hackernews/pull/34))

## v0.5.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#30](https://github.com/turbot/steampipe-plugin-hackernews/pull/30))
- Recompiled plugin with Go version `1.19`. ([#30](https://github.com/turbot/steampipe-plugin-hackernews/pull/30))

## v0.4.0 [2022-07-22]

_Bug fixes_

- Fixed inconsistent table name in the `hackernews_top` table which would cause intermittent caching issues. ([#28](https://github.com/turbot/steampipe-plugin-hackernews/pull/28))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v332--2022-07-11) which includes several caching fixes. ([#28](https://github.com/turbot/steampipe-plugin-hackernews/pull/28))

## v0.3.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#24](https://github.com/turbot/steampipe-plugin-hackernews/pull/24))

## v0.3.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#22](https://github.com/turbot/steampipe-plugin-hackernews/pull/22))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#21](https://github.com/turbot/steampipe-plugin-hackernews/pull/21))

## v0.2.0 [2022-04-05]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v2.2.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v220--2022-03-30)

## v0.1.0 [2021-11-23]

_Enhancements_

- Recompiled plugin with Go version 1.17 ([#15](https://github.com/turbot/steampipe-plugin-hackernews/pull/15))
- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#14](https://github.com/turbot/steampipe-plugin-hackernews/pull/14))

## v0.0.3 [2021-10-06]

_Enhancements_

- Updated the README file to include GitHub clone link and license information ([#11](https://github.com/turbot/steampipe-plugin-hackernews/pull/11))
- Updated plugin license to Apache 2.0 per [turbot/steampipe#488](https://github.com/turbot/steampipe/issues/488) ([#8](https://github.com/turbot/steampipe-plugin-hackernews/pull/8))

## v0.0.2 [2021-03-07]

_Enhancements_

  - Update plugin README file
 

## v0.0.1 [2021-02-25]

_What's new?_

- Initial release with tables

  - hackernews_ask_hn
  - hackernews_best
  - hackernews_item
  - hackernews_job
  - hackernews_new
  - hackernews_show_hn
  - hackernews_top
  - hackernews_user
