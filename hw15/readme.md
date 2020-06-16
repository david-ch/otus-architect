# Как запустить
1. Если prometheus и ingress не установлены, то установить выполнив из папки infra

        ./install_prom.sh && ./install_ingress.sh

2. Установить приложение командой (устанавливается в default namespace)

        helm install hw15 ./chart
        
    или
    
        skaffold run

3. Запустить тесты

        newman run otus-hw15.postman_collection.json

    по-умолчанию запросы идут на arch.homework/otus.
    для изменения можно использовать опцию: 

        --env-var "baseUrl=another.url/app"

4. Посмотреть лог product-service

        kubectl logs pod/<podname>

Product-service пишет в лог информацию о работе с кэшем

# Настройка

Для включения/выключения кэша используется helm value "product.cache.enabled"

# Описание решения
Product-service хранит в БД(postgres) информацию о продуктах и дает возможность работать с ними через api

Для кэширования продуктов используется redis. Данные сохраняются с TTL в одну минуту. Наличие TTL делает данные волатильными с точки зрения redis.
Для redis включена настройка "maxmemory-policy volatile-lru", что означает вытеснение в первую очередь редкоиспользуемых волатильных данных при нехватке памяти.

