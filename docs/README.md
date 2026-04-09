# Документация `forge`

`forge` — CLI для управления данными окружений и конфигурацией деплоя в HashiCorp Vault.

Поддерживаемые сценарии:
- хранение текущего состояния окружений (`data`) в Vault;
- генерация итогового YAML из Go-шаблона;
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

## Модель данных `env`

`forge env` использует проектный секрет со структурой:
```json
{
  "template": "api-stack.yml",
  "environments": {
    "stage": {
      "placement": {
        "enabled": false
      },
      "data": {}
    },
    "prod": {
      "placement": {
        "enabled": false
      },
      "data": {}
    }
  }
}
```

Ключевые правила:
- поле `mode` не используется;
- для каждого окружения хранится своя независимая `data`;
- `placement` хранится отдельно от `data`;
- рендер (`env get -c`) выполняется по `template` и объединенному контексту `environments.<env>.{placement,data}`.

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

Извлечение изменяемых переменных Go template из YAML-шаблона:
```bash
forge templates vars <template.yml>
```

### Окружение (`env`)

Инициализация проекта:
```bash
forge env init <project> -e <environment> -t <template.yml>
```

Пример:
```bash
forge env init api -e stage -t api-stack.yml
```
`init` автоматически создает `environments.<env>.placement` + `environments.<env>.data` и заполняет `data` ключами из `{{...}}` в шаблоне.

Добавление нового окружения в существующий проект:
```bash
forge env add <project> -e <environment>
```
`add` также автоматически заполняет `data` ключами из текущего шаблона проекта и выставляет `placement.enabled=false`.

Удаление окружения из проекта:
```bash
forge env delete <project> -e <environment>
```

Обновление параметров:
```bash
forge env set <project> -e <environment> -p key=value
```

Пример вложенных ключей:
```bash
forge env set api -e stage -p api.image=registry.service/app:1.2.3 -p api.replicas=2
```

Настройка `placement` через переменные окружения:
```bash
forge env set api -e stage -p placement.enabled=true -p placement.constraint=node.role==worker
```

Отключение `placement`:
```bash
forge env set api -e stage -p placement.enabled=false
```

Просмотр текущих значений:
```bash
forge env get <project> -e <environment>
```

Рендер итогового конфигурационного файла из шаблона:
```bash
forge env get <project> -e <environment> -c
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
forge env init api -e stage -t api-stack.yml
```
3. Обновить тег:
```bash
forge env set api -e stage -p api.image=registry.service.uhvahta.ru/api/backend:sha-abc1234 -p api.replicas=1
```
4. Сгенерировать итоговый YAML:
```bash
forge env get api -e stage -c > /tmp/api-stage.yml
```

## Примечания

- При `init` в Vault создается/обновляется структура `template + environments.<env>.{placement,data}`.
- Для команд `env get/set` флаг `-e/--env` обязателен.
- Для условного `placement` используйте блок `if` в шаблоне:
```yaml
deploy:
  mode: replicated
  {{- if .placement.enabled }}
  placement:
    constraints:
      - "{{ .placement.constraint }}"
  {{- end }}
  replicas: {{ .api.replicas }}
```
- Примеры состояния секрета после инициализации:  
  ![init_compose](/docs/images/init_compose.png)  
  ![init_swarm](/docs/images/init_swarm.png)
