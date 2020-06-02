1. Если Prometheus и Ingress не установлены, запустить из папки infra

        ./install_prom.sh && ./install_ingress.sh

2. В chart/values.yaml указать адрес ingress в свойстве ingressAddress
3. Запустить установку helm с опцией --wait

        helm install hw07 ./chart --wait

После старта приложения запустится stress-test-job, который будет создавать нагрузку
