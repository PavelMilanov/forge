
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

  Перед инициализацией конфигурационного файла необходимо создать шаблон и расположить его в каталоге `var/forge/templates`.

  Инициализация конфигурационных файла разрешена только из списка доступных шаблонов.

##### template
```bash
forge env tmpl list
```

Выводит список доступных шаблонов.

##### env
###### init
```bash
forge env init <project> -t <template.yml> -m compose
```

Сопоставляет шаблон указанного файла и создает `vault secret` исходя из его спецификации.

параметры:
- `<project>` название проекта.

флаги:
- `-t` файл конфигурации;
- `-m` спецификация файла. По-умолчанию `compose`.
	Разрешенные значения: `compose` | `swarm`

пример команды:
```bash
forge env init dev -t template.yaml -m compose
```

###### set
```bash
forge env set <project> -p tag=<string> -p replicas=<number>
```

Обновляет параметры модели указанного проекта согласно спецификации.

параметры:
- `<project>` название проекта.

флаги:
- `-p` параметр, который необходимо обновить. Можно указать несколько.

пример команды:
```bash
forge env set dev -p tag=alpine
```
>Для спецификации docker swarm можно обновлять один произвольный параметр или указать сразу два.

###### get
```bash
forge env get <project>
```

```bash
forge env get <project> -c
```

Выводит информацию о текущей конфигурации указанного проекта.

параметры:
- `<project>` название проекта.

флаги:
- `-c` вывод конфигурационного файла.

пример команды:
```bash
forge env get dev
```
> При вызове команды без флагов в консоль будет выведен форматированный вывод данных согласно спецификации.

###### versions
```bash
forge env versions <project>
```

Выводит информацию о версиях конфигурации указанного проекта.

параметры:
- `<project>` название проекта.

пример команды:
```bash
forge env versions dev
```

###### rollback
```bash
forge env rollback <project> <version> -v <version>
```

Откатывает конфигурацию проекта до указанной версии.

параметры:
- `<project>` название проекта.

флаги:
- `-v` версия конфигурации в Vault.

пример команды:
```bash
forge env rollback dev -v 2
```

##### deploy

##### enable completion

###### bash
```bash
forge completion bash > completion.sh
chmod +x completion.sh
sudo mv completion.sh /etc/bash_completion.d/
source ./bashrc
```
