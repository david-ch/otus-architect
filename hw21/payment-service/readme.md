Сервис отвечает за обработку платежей

Использует канал payment для рассылки событий

# Оплата заказа
По событию created из канала order запускается оплата.

На счету пользователя резервируется нужная сумма и в канал payment отправляется событие orderPaid

Если по каким-то причинам оплата не проходит в канал payment отправляется событие orderPaymentFailed

# Отмена оплаты
По событию orderItemsReservationFailed из канала warehouse оплата отменяется.

Отсылается событие orderPaymentFailed в канала payment