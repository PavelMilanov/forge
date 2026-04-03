# Документация `forge`

`forge` — CLI для управления переменными деплоя и версионированием конфигурации в HashiCorp Vault.

Поддерживаемые сценарии:
- хранение текущего состояния деплоя (`image`, `tag`, `replicas`) в Vault;
- генерация итогового YAML из Go-шаблона;
- обновление/откат версий секрета;
- деплой и обновление стеков через Portainer API.

## Конфигурация

По умолчанию `forge` читает файл `var/forge/forge.yml` (относительно текущей рабочей директории процесса).

Минимальная конфигурация:
```yaml
vault:
  url: "http://127.0.0.1:8200"
  role_id: "approle-role-id"
  secret_id: "approle-secret-id"

agent:
  type: ssh
```

Конфигурация для работы `deploy` через Portainer:
```yaml
vault:
  url: "http://127.0.0.1:8200"
  role_id: "approle-role-id"
  secret_id: "approle-secret-id"

agent:
  type: portainer
  credentials:
    url: "http://portainer.local:9000"
    key: "portainer-api-key"
    teams: [5]
```

Пояснение к `teams`:
- список ID команд Portainer, которым назначается доступ к новому стеку после `create`;
- если `teams` не задан, шаг обновления `resource control` пропускается.

Путь к шаблонам: `var/forge/templates`.

## Спецификации

### Compose

Отслеживаемые параметры:
- `image`
- `tag`

Пример шаблона:
```yaml
services:
  registry:
    image: {{.Image}}:{{.Tag}}
```

### Swarm

Отслеживаемые параметры:
- `image`
- `tag`
- `replicas`

Пример шаблона:
```yaml
services:
  registry:
    image: {{.Image}}:{{.Tag}}
    deploy:
      replicas: {{.Replicas}}
```

### Kubernetes

Режим `kubernetes` доступен как значение флага `-m`, но в текущей версии реализован как заглушка и не генерирует рабочую конфигурацию.

## Команды

### Шаблоны

Список доступных шаблонов:
```bash
forge templates list
```

Алиасы команды:
```bash
forge tpl list
forge tmpl list
```

### Окружение (`env`)

Инициализация проекта:
```bash
forge env init <project> -t <template.yml> -m <compose|swarm|kubernetes>
```

Пример:
```bash
forge env init dev -t registry.yml -m compose
```

Обновление параметров:
```bash
forge env set <project> -p tag=<value> -p image=<value>
```

Для `swarm` также доступен параметр `replicas`:
```bash
forge env set stage -p tag=1.2.3 -p replicas=3
```

Просмотр текущих значений:
```bash
forge env get <project>
```

Рендер итогового конфигурационного файла из шаблона:
```bash
forge env get <project> -c
```

История версий секрета:
```bash
forge env versions <project>
```

Откат к версии:
```bash
forge env rollback <project> -v <version>
```

### Деплой (`deploy`)

Список endpoint-ов Portainer:
```bash
forge deploy list
```

Создание или обновление стека из локального файла:
```bash
forge deploy stack file <endpoint> -n <stack-name> -f <path-to-stack.yml>
```

Создание нового стека из custom template в Portainer:
```bash
forge deploy stack template <endpoint> -n <stack-name> -t <template-name>
```

Режим выполнения операции для deploy из файла:
```bash
forge deploy stack file <endpoint> -n <stack-name> -f <path> --mode <create|update|upsert>
```

Обновление существующего стека без изменения источника (legacy-сценарий):
```bash
forge deploy stack refresh <endpoint> -n <stack-name>
```

`refresh` перечитывает текущий `StackFileContent` стека из Portainer и повторно применяет его как update (без локального файла и без template).

## Примеры рабочего потока

1. Добавить шаблон в `var/forge/templates`.
2. Инициализировать проект:
```bash
forge env init admin-stage -t stack.yml -m swarm
```
3. Обновить тег:
```bash
forge env set admin-stage -p tag=sha-abc1234
```
4. Сгенерировать итоговый YAML:
```bash
forge env get admin-stage -c > /tmp/admin-stage.yml
```

## Примечания

- При `init` в Vault создается секрет с полями `deploy`, `mode`, `template`.
- Если проект уже существует в Vault, команда `env init` выводит `The project already initialized`.
- Примеры состояния секрета после инициализации:  
  ![init_compose](/docs/images/init_compose.png)  
  ![init_swarm](/docs/images/init_swarm.png)
