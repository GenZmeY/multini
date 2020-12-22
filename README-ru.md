# Multini

***
[English](README.md), [Русский](README-ru.md)
***

[![build](https://github.com/GenZmeY/multini/workflows/build/badge.svg)](https://github.com/GenZmeY/multini/actions?query=workflow%3Abuild)
[![tests](https://github.com/GenZmeY/multini/workflows/tests/badge.svg)](https://github.com/GenZmeY/multini/actions?query=workflow%3Atests)
[![CodeQL](https://github.com/GenZmeY/multini/workflows/CodeQL/badge.svg)](https://github.com/GenZmeY/multini/security/code-scanning)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/GenZmeY/multini)](https://golang.org)
[![GitHub](https://img.shields.io/github/license/genzmey/multini)](LICENSE)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/GenZmeY/multini)](https://github.com/GenZmeY/multini/releases)

*Утилита командной строки для манипулирования ini файлами с дублирующимися именами ключей.*

Скомпилированная версия multini доступна на [странице релизов](https://github.com/GenZmeY/multini/releases).

***

# Описание
Некоторые программы используют формат ini файлов допускающий повторяющиеся имена ключей.  
Например игры основаные на [unreal engine](https://en.wikipedia.org/wiki/Unreal_Engine).  
Это может выглядеть так (часть конфигурации Killing Floor 2):
```
[OnlineSubsystemSteamworks.KFWorkshopSteamworks]
ServerSubscribedWorkshopItems=2267561023
ServerSubscribedWorkshopItems=2085786712
ServerSubscribedWorkshopItems=2222630586
ServerSubscribedWorkshopItems=2146677560
```
Большинство реализаций поддерживают только одно свойство с заданным именем в разделе. Если их несколько, будет обрабатываться только первый (или последний) ключ, чего в данном случае недостаточно. multini решает эту проблему.

**примечание:**  
- multini чувствителен к регистру;  
- кавычки вокруг значения не обрабатываются (multini считает их частью значения);  
- многострочные значения не поддерживаются.  
(но все это может измениться в будущем)  

# Сборка и установка (вручную)
1. Установите [golang](https://golang.org), [git](https://git-scm.com/), [make](https://www.gnu.org/software/make/);
2. Клонируйте этот репозиторий: `git clone https://github.com/GenZmeY/multini`
3. Перейдите в каталог с исходниками: `cd multini`
4. Выполните сборку: `make`
5. Выполните установку: `make install`

# Использование
```
Использование: multini [ПАРАМЕТРЫ]... ДЕЙСТВИЕ ini_file [секция] [ключ] [значение]
Действия:
  -g, --get          Получить значения для заданной комбинации параметров.
  -s, --set          Установить значения для заданной комбинации параметров.
  -a, --add          Добавить значения для заданной комбинации параметров.
  -d, --del          Удалить указанную комбинацию параметров.
  -c, --chk          Показать ошибки синтаксического анализа для указанного файла.

Параметры:
  -e, --existing     Для --set и --del завершить программу с ошибкой, если элемент остутствует.
  -r, --reverse      Для --add добавлять элемент в начало секции
  -i, --inplace      Перезаписывать исходный файл.
                     Это не атомарно, но требует меньше разрешений
                     чем способ по умолчанию с заменой файла.
  -o, --output ФАЙЛ  Записать результат в ФАЙЛ. '-' означает стандартный вывод
  -u, --unix         Использовать LF в конце строки
  -w, --windows      Использовать CRLF в конце строки
  -q, --quiet        Подавить весь вывод
  -h, --help         Отобразить страницу помощи
      --version      Отобразить версию
```

# Примеры
**вывести глобальное значение вне секции:**  
`multini --get ini_file '' key`

**вывести секцию:**  
`multini --get ini_file section`

**вывести список секций:**  
`multini --get ini_file`

**вывести значение:**  
`multini --get ini_file section key`  
- если ключей несколько, отобразится список всех значений этих ключей

**создать/обновить ключ (в единственном экземпляре):**  
`multini --set ini_file section key value`  
- если ключа нет, он будет добавлен  
- если ключ существует, значение будет обновлено  
- если ключ существует и имеет несколько значений, будет установлен ключ с указанным значением, остальные значения будут удалены

**добавить ключ с указанным значением:**  
`multini --add ini_file section key value`  
- если ключа нет, он будет добавлен
- если ключ существует и не имеет указанного значения, будет добавлено новое значение
- если указанное значение повторяет существующее, никаких изменений не будет

**удалить все ключи с указанным именем:**  
`multini --del ini_file section key`

**удалить ключ с указанным именем и значением:**  
`multini --del ini_file section key value`

**удалить секцию:**  
`multini --del ini_file section`

**короткие версии параметров можно комбинировать:**  
`multini -gq ini_file section key value`  
- проверить наличие ключа с заданным значением, используя код возврата

# Лицензия
Copyright © 2020 GenZmeY

[MIT License](LICENSE).

