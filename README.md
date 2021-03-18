# oneget
Консольная утилита для загрузки пакетов с releases.1c.ru


[![Release](https://img.shields.io/github/release/v8platform/oneget.svg?style=for-the-badge)](https://github.com/v8platform/oneget/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Build status](https://img.shields.io/github/workflow/status/v8platform/oneget/goreleaser?style=for-the-badge)](https://github.com/v8platform/oneget/actions?workflow=goreleaser)
[![Codecov branch](https://img.shields.io/codecov/c/github/v8platform/oneget/master.svg?style=for-the-badge)](https://codecov.io/gh/v8platform/oneget)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](http://godoc.org/github.com/v8platform/oneget)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=for-the-badge)](https://saythanks.io/to/khorevaa)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=for-the-badge)](https://github.com/goreleaser)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)

### Команда `get` описание использования

Команда получения релизов проектов с сайта 1С `https://releases.1c.ru/`

Быстрый запуск:
```shell
export ONEC_USERNAME=user
export ONEC_PASSWORD=password
oneget get --path ./tmp/dist/ platform@8.3.18.1334

# or
oneget --user user --pwd password get --path ./tmp/dist/ platform83@8.3.18.1334

```

#### Описание формата аргумента `RELEASE`

Шаблон формата `platform83[:filter.[filter]...]@8.3.18.1334[:filter]`
Где, 
 * `platform83` - имя проекта (обязательный)                                  
 * `[:filter.[filter]...]` - набор фильтров файлов                                   
 * `@` - разделитель между проектов и версией релиза                                 
 * `8.3.18.1334[:filter]` - описание версии релиза                                 

Минимальный использование указание только имени проекта.
Например, `platform83` - будет трактоваться как `platform83@latest`

> Имя проекта - подсмотреть можно в адресе, ссылки имеют вид например https://releases.1c.ru/project/EnterpriseERP20

> Синонимы проектов для быстрого доступа:
> * `platform` -> `platform83`
> * `edt`      -> `DevelopmentTools10`
> * `ring`     -> `EnterpriseLicenseTools`
> * `executor` -> `Executor`
> * `pg`       -> `AddCompPostgre`

##### Набор фильтров 
Список предопределенных фильтров для проектов:
 * По ОС:
   * `win`, `windows`- фильтр по MS Windows
   * `mac`           - фильтр по OS X
   * `deb`           - фильтр по DEB-based Linux-систем
   * `rpm`           - фильтр по RPM-based Linux-систем
* По разрядности OS:
    * `x32`           - фильтр по 32-bit (по умолчанию, если не указан фильтр разрядности)
    * `x64`           - фильтр по 64-bit

> Важно! Для OSX флаг разрядности игнорируется

**Пример использования:**
```shell
  export ONEC_USERNAME=user
  export ONEC_PASSWORD=password
  # Т.к. не указана разрядность OS будет скачены дистрибутивы для x32 
  # скачать файлы с фильтрацией по Windows
  oneget get platform:win  
  # скачать файлы с фильтрацией по OSX
  oneget get platform:mac 
  # скачать файлы с фильтрацией по DEB-based Linux-систем
  oneget get platform:deb 
  # скачать файлы с фильтрацией по RPM-based Linux-систем
  oneget get platform:rpm
```
**Пример для x64:**
```shell
  export ONEC_USERNAME=user
  export ONEC_PASSWORD=password
  
  # Там где не указана разрядность OS будет скачены дистрибутивы для x32 
  # скачать файлы с фильтрацией по Windows
  oneget get platform:win.x64  
  # скачать файлы с фильтрацией по OSX
  # Важно для OSX флаг разрядности игнорируется 
  oneget get platform:mac 
  # скачать файлы с фильтрацией по DEB-based Linux-систем
  oneget get platform:deb.x64 
  # скачать файлы с фильтрацией по RPM-based Linux-систем двух разрядностей сразу
  oneget get platform:rpm.x64 platform:rpm.x32
```
##### Специальные фильтры для проектов       
**Для проекта platform (platform83)**
   * `thin-client`, `thin`   - фильтр для файлов тонкого клиента 1С. Предприятие
   * `client`              - фильтр для файлов клиента 1С. Предприятие
   * `server`              - фильтр для файлов сервера 1С. Предприятие
   * `full`                - фильтр для файлов "Технологическая платформа" (только для Windows)

> Важно! Для OSX фильтр `server` игнорируется

> Важно! Фильтр `full` игнорируется для всех других фильтров кроме `win`

**Пример использования:**
```shell
  export ONEC_USERNAME=user
  export ONEC_PASSWORD=password

  # Там где не указана разрядность OS будет скачены дистрибутивы для x32 
  # скачать файлы сервера для Windows
  oneget get platform:win.server.x64  
  # скачать файлы клиента для OSX
  # Важно для OSX флаг разрядности игнорируется 
  oneget get platform:mac.client 
  # скачать файлы тонкого клиента для DEB-based Linux-систем
  oneget get platform:deb.thin.x64 
  # скачать файлы сервера для RPM-based Linux-систем
  oneget get platform:rpm.server.x64
```  

**Для проекта edt (DevelopmentTools10)**

  * `jdk`    - фильтр для файлов Bellsoft JDK
  * `online` - фильтр для файлов онлайн установщика 1С:EDT

> Важно. Для проекта `edt` игнорируются фильтры разрядности

**Пример использования:**
```shell
  export ONEC_USERNAME=user
  export ONEC_PASSWORD=password

  # скачать файлы 1C:EDT для Windows
  oneget get edt:win  
  # скачать файлы 1C:EDT для OSX
  oneget get edt:mac 
  # скачать файлы 1C:EDT для Linux и Bellsoft JDK для DEB-based Linux-систем
  oneget get edt:deb 
  # скачать файлы 1C:EDT для Linux и Bellsoft JDK для RPM-based Linux-систем
  oneget get edt:rpm
 
   # скачать файлы онлайн установщика 1C:EDT для Windows
  oneget get edt:win.online 
```
##### Описание формата версии релиза

> В версии релиза может быть указан номер версии или специальные фильтры версии.

> Если версия релиза пустая, то подставляется фильтр "latest" 
> ("edt" -> "edt@latest")

**Специальные фильтры версии релиза:**

 * `latest`         - выбирает наиболее старшую версию релиза
 * `latest:regexp`  - фильтрует список версию по `regexp`, и берет наиболее старшую
 * `from:date`      - фильтрует список версий по дате, у которых дата релиза больше `date` где, `date` - формате 02.06.21
 * `from-v:version` - фильтрует список версий, которые старше версии релиза `version` где, `version` - формате номер версии
 * `regexp`        - фильтрует список по регулярному выражению указанному в `regexp`

**Пример использования:**
```shell
  export ONEC_USERNAME=user
  export ONEC_PASSWORD=password

  # скачать файлы последней версию релиза 1C:EDT для Windows
  oneget get edt:win@latest 
  # or / или
  # oneget get edt:win
   
  # скачать файлы Платформы 1С. Предприятие для всех систем
  # всех версии релизов, выпущенные начиная с даты 2020.01.01
  oneget platform@from:01.01.21
  
  # скачать файлы Платформы 1С. Предприятие для DEB-based Linux-систем
  # всех версии релизов, у которых версия старше чем 8.3.18.1363
  oneget platform:deb.x64@from-v:8.3.18.1363
 
 
  # скачать файлы сервера Платформы 1С. Предприятие для DEB-based Linux-систем
  # последней выпущенной версии 8.3.16
  oneget platform:deb.server.x64@latest:8.3.16
 
  # скачать файлы Платформы 1С. Предприятие для OSX
  # всех версии релизов 8.3.16.x
  oneget platform:mac@8.3.16
```

## Запуск в докере

```shell
docker run -v $(pwd):/tmp/dist demoncat/oneget \
    --user $ONEC_USERNAME \
    --pwd $ONEC_PASSWORD \
    --path /tmp/dist/ \
    --nicks platform83 \
    --version-filter 8.3.16.1876 \
    --distrib-filter 'deb64.tar.gz$'
```

## Настройка логов

### Через файл настройки
Создать рядом с приложением файл `logos.yaml` с содержимым

```yaml
appenders:
  console:
    - name: CONSOLE
      target: stdout
      encoder:
        console:

  rolling_file:
    - name: FILE
      file_name: ./logs/oneget.log
      max_size: 100
      max_age: 10
      encoder:
        json:
loggers:
  root:
    level: info
    appender_refs:
      - CONSOLE
  logger:
    - name: "github.com/v8platform/oneget"
      appender_refs:
        - CONSOLE
        - FILE
      level: debug     

```

### Через переменные окружения
```shell
export LOGOS_CONFIG="appenders.rolling_file.0.name=FILE;
appenders.rolling_file.0.file_name=./logs/oneget.log;
appenders.rolling_file.0.max_size=100;
appenders.rolling_file.0.encoder.json;
loggers.logger.0.level=debug;
loggers.logger.0.name=github.com/v8platform/oneget;
loggers.logger.0.appender_refs.0=CONSOLE;
loggers.logger.0.appender_refs.1=FILE;"
```
#TODO

Идея и часть кода взята из [этого](https://github.com/korableg/Downloader1C) проекта 