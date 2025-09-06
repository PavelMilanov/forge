
Утилита **forge** позволяет динамически обновлять конфигурацию файлов `docker compose ` и `docker swarm.`
Для контроля версий используется хранилище Vault.
#### Конфигурация

Для работы утилиты необходимо создать конфигурационный файл `forge.yml` в директории `/var/forge/`. Пример:
```yml
vault:
  url: http://127.0.0.1:8080
  token: vault_token
```

#### Docker compose

Спецификация позволяет отслеживать **образ и его тег**. Пример:
```yaml
docker-compose.yaml
services:
  registry:
    image: {{.Image}}:{{.Tag}}
    container_name: registry
    restart: unless-stopped
    ports:
      - 5050:5050
    volumes:
      - ./conf:/registry/conf.d:ro
      - registry-data:/registry/var

volumes:
  registry-data:
```

При инициализации конфигурационного файла в хранилище Vault появится секрет:
![init_compose](/docs/images/init_compose.png)
#### Docker swarm

Спецификация позволяет отслеживать **названия образа и его тег** и **количество реплик**. Пример:
```yaml
docker-stack.yaml
services:
  registry:
    image: {{.Image}}:{{.Tag}}
    deploy:
      resources:
        limits:
          cpus: "0.40"
          memory: 400M
        reservations:
          cpus: "0.30"
          memory: 400M
      mode: replicated
      replicas: {{.Replicas}}
      restart_policy:
        condition: on-failure
        max_attempts: 3
        window: 15s
      update_config:
        parallelism: 1
        delay: 10s
        order: start-first
        failure_action: rollback
      rollback_config:
        parallelism: 1
        delay: 10s
        order: start-first
        failure_action: rollback
    ports:
      - 5050:5050
    volumes:
      - ./conf:/registry/conf.d:ro
      - registry-data:/registry/var

volumes:
  registry-data:
```

При инициализации конфигурационного файла в хранилище Vault появится секрет:
![init_swarm](/docs/images/init_swarm.png)
#### Использование

#####  init
```bash
forge init -f <file> -m compose -a <string>
```

Сопоставляет шаблон указанного файла и создает `vault secret` исходя из его спецификации.

флаги:
- `-f` путь до файла конфигурации;
- `-m` спецификация файла. По-умолчанию `compose`;
	Разрешенные значения: `compose` | `swarm`
- `-a` название секрета в Vault / название проекта.

пример команды:
```bash
forge init -f /path/to/template.yaml -m compose -a dev
```

##### set
```bash
forge set <project> -p tag=<string> -p replicas=<number>
```

Обновляет параметры модели указанного проекта согласно спецификации.

параметры:
- `<project>` название проекта.

флаги:
- `-p` параметр, который необходимо обновить. Можно указать несколько.

пример команды:
```bash
forge set dev -p tag=alpine
```
>Для спецификации docker swarm можно обновлять один произвольный параметр или указать сразу два.

##### get
```bash
forge get <project>
```

```bash
forge get <project> -p <param>
```

Выводит информацию о текущей конфигурации указанного проекта.

параметры:
- `<project>` название проекта.

флаги:
- `-p` параметр, который необходимо получить и вывести в консоль.

пример команды:
```bash
forge get dev -p tag
```
> При вызове команды без флагов в консоль будет выведен форматированный вывод данных согласно спецификации.

##### deploy
```bash
forge deploy <project> -f path/to/file.yaml
```

Генерирует файл конфигурации согласно сохраненным данным в `Vault` и указанному шаблону.
Итоговый файл конфигурации будет с генерирован в директории `/var/forge/` с названием в формате `спецификация-проект.yml`

параметры:
- `<project>` название проекта.
-
флаги:
- `-f` путь до файла конфигурации.

пример команды:
```bash
forge deploy dev -f path/to/template.yaml
```
> будет сгенерирован файл /var/forge/compose-dev.yml
