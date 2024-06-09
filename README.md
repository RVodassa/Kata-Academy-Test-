<p align="center">
      <img src="https://i.ibb.co/ncyc7K0/2024-06-06-10-31-10.png" width="726">
</p>

<p align="center">
   <img alt="Static Badge" src="https://img.shields.io/badge/Last%20version-2.0-brightgreen">
   <img alt="Static Badge" src="https://img.shields.io/badge/Language-Golang-blue">
</p>

## Что это ? 

# Это тестовое задание от Kata Academy: калькулятор с двумя системами исчисления (римская и арабская)
## Условия:
- Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.
- Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…), так и с римскими (I, II, III, IV, V…) числами.
- Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.
- Калькулятор умеет работать только с целыми числами.
- Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен выдать панику и прекратить работу.
- При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими. 
- При вводе пользователем не подходящих чисел приложение выдаёт панику и завершает работу. 
- При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выдаёт панику и завершает работу. 
- Результатом операции деления является целое число, остаток отбрасывается.
- Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна выдать панику. 

### Примеры работы программы
|INPUT                |OUTPUT                    
|----------------|------------------
|2 * 2           | 4
|10 / 4          | 2            
|  5 / 0         | panic: Выдача паники, так как диапазон доступных чисел: от 1 до 10
| IV * IX				 | XXXVI
|3 - II					 | panic: Выдача паники, так как используются одновременно разные системы счисления
|		IIII + V		 | panic: Выдача паники, так как используются неверные римские числа.
| V - X  				 | panic: Выдача паники, так как в римской системе нет отрицательных чисел.
| XI + II        | panic: Выдача паники, так как диапазон доступных чисел: от I до X
| 1							 |panic: Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).
|2               | panic: Выдача паники, так как строка не является математической операцией.
|4 + 4 + 2       | panic: Выдача паники, так как формат математической операции не удовлетворяет заданию.
|x - x           | panic: Выдача паники, так как в римской системе нет нуля.

## Developers

- [RVodassa](https://github.com/RVodassa)

