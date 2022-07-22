## v0.4.0 [2022-07-22]

_Bug fixes_

- Fixed inconsistent table name in the `hackernews_top` table which would cause intermittent caching issues. ([#28](https://github.com/turbot/steampipe-plugin-hackernews/pull/28))
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
