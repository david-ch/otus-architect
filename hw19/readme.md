# Как запустить
1. Если prometheus и ingress не установлены, то установить выполнив из папки infra

        ./install_prom.sh && ./install_ingress.sh

1. Установить приложение командой (устанавливается в default namespace)

        helm install hw15 ./chart
        
    или
    
        skaffold run

1. Запустить тесты

        newman run otus-hw19.postman_collection.json

    по-умолчанию запросы идут на arch.homework/otus.
    для изменения можно использовать опцию: 

        --env-var "baseUrl=another.url/app"

# Описание
POST /order создает новый заказ

При этом в заголовке X-Request-Id ожидается сгенерированный клиентом идентификатор запроса

Повторный запрос с одним X-Request-Id приводит к ошибки

Проверка делается на уровне БД - идентификатор сохраняется вместе с заказом. На поле настроен unique constraint