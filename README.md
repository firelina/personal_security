# personal_security
В 1 практике была добавлен ассинхронная работа с бд с помощью гоурутин
Была подключена СУБД postgres была выбрана postgres из-за своей простоты
Были доофрмлены слои:
router на фреймворке gin go для реализации api
handlers принимает для обработки входящих данных
usecase для реализации логики
repository для доступа к бд
Примеры работы
![login.PNG](..%2F..%2F..%2F..%2FHSE%2F%EF%E8%F2%EE%ED%2Flogin.PNG)
![create_contact.PNG](..%2F..%2F..%2F..%2FHSE%2F%EF%E8%F2%EE%ED%2Fcreate_contact.PNG)
![create_event.PNG](..%2F..%2F..%2F..%2FHSE%2F%EF%E8%F2%EE%ED%2Fcreate_event.PNG)
![create_reminder.PNG](..%2F..%2F..%2F..%2FHSE%2F%EF%E8%F2%EE%ED%2Fcreate_reminder.PNG)
![send_remind.PNG](..%2F..%2F..%2F..%2FHSE%2F%EF%E8%F2%EE%ED%2Fsend_remind.PNG)