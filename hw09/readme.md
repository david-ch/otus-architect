# Как запустить
1. Если prometheus и ingress не установлены, то установить выполнив из папки infra

    ./install_prom.sh && ./install_ingress.sh

2. установить сервисы командами

    helm install hw09-auth ./auth-service/chart
    helm install hw09-user ./user-service/chart

Если требуется изменить имя релиза или использовать не неймспейс по умолчанию, то в обоих чартах нужно указать верный serviceAddress

3. запустить тесты

    newman run otus-hw09.postman_collection.json

по-умолчанию запросы идут на arch.homework/otus
для изменения можно использовать опцию --env-var "baseUrl=another.url/app"

# Описание решения
auth-service занимается аутентификацией и менеджментом сессий (хранятся в памяти сервиса)
user-service хранит информации о профилях пользователей. В том числе, хранит хэши паролей и занимается проверкой введенного пароля

между клиентом и nginx передается cookie sessionId с идентификатором текущей сессии
между сервисами и nginx используется заголовок X-Username. В случае межсервисного взаимодейсвтия используется X-Username=tech

![arch](/readme.assets/arch.png)