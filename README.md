# personal_security
В 1 практике была добавлен ассинхронная работа с бд с помощью гоурутин
Была подключена СУБД postgres была выбрана postgres из-за своей простоты
Были доофрмлены слои:
router на фреймворке gin go для реализации api
handlers принимает для обработки входящих данных
usecase для реализации логики
repository для доступа к бд
были обновлены тесты 
был сформирован dockerfile
Примеры работы
![login.PNG](login.PNG)
![create_contact.PNG](create_contact.PNG)
![create_event.PNG](create_event.PNG)
![create_reminder.PNG](create_reminder.PNG)
![send_remind.PNG](send_remind.PNG)