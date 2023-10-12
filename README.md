# Grocery List (Golang)

## Описание проекта

Этот проект представляет собой простое приложение для создания и управления списком покупок, написанное на языке программирования Golang. С его помощью вы можете легко добавлять, удалять и отмечать продукты в вашем списке покупок.

## Развертывание проекта

Для развертывания этого проекта на вашей локальной машине, выполните следующие шаги:

1. Установите Go (если у вас его еще нет) с официального сайта: [https://golang.org/dl/](https://golang.org/dl/)

2. Клонируйте репозиторий с помощью Git:

   ```
   git clone https://github.com/RobertGoodman08/grocery_list_golang.git
   ```

3. Перейдите в каталог проекта:

   ```
   cd grocery_list_golang
   ```

4. Установите зависимости проекта:

   ```
   go get -d -v ./...
   ```

5. Подключите базу данных. Вам потребуется создать базу данных.

6. В файле `package/db/db.go` укажите параметры подключения к вашей базе данных:

   ```go
   const (
       dbHost     = "localhost"
       dbPort     = your  port
       dbUser     = "your  user"
       dbPassword = "your  password"
       dbName     = "your  dbname"
   )
   ```

7. Запустите приложение:

   ```
   go run ./cmd
   ```

После выполнения этих шагов приложение будет доступно по адресу `http://localhost:8080`.

## Основной функционал

- Добавление продуктов в список покупок.
- Удаление продуктов из списка.
- Отметка продуктов как купленных.
- Просмотр всего списка покупок.


## Лицензия

Этот проект распространяется под лицензией MIT. Подробности можно найти в файле [LICENSE](LICENSE).



