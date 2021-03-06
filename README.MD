## CSV datasource for Grafana 6.x.x

#### Install
- Copy files from the `dist` folder to your [Grafana plugin directory](https://grafana.com/docs/grafana/latest/plugins/installation/#grafana-plugin-directory)
- Restart Grafana
- Check datasource list as http://you-grafana/datasources/new

> The plugin is unsigned, hence you may face the error:
>
> `lvl=eror msg=“Failed to load plugin” logger=plugins error=“plugin “grafana-csv-plugin” is unsigned”`
>
> To make it work you should make configuration as described in [documentation](https://grafana.com/docs/grafana/latest/installation/configuration/#allow-loading-unsigned-plugins)

#### Features:
- Read local CSV file
- Read remote CSV file by SFTP
- Filter data (thanks to a great [expr library](https://github.com/antonmedv/expr/blob/master/docs/Language-Definition.md))

###### Simple table

![](./doc/image/grid.png)

###### With filtering expression

![](./doc/image/expr.png)

![](./doc/image/expr1.png)

#### Config
- Read local file

![](./doc/image/config_local.png)

- Read remote file

![](./doc/image/config_sftp.png)

#### Build
- npm run build

#### Test
- npm test


###### Example data set: /data
###### Icon: https://freeicons.io/vector-file-types-icons/csv-icon-2272
