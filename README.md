# Mathematics

Библиотека `mathematics` предоставляет обширный набор математических функций и утилит для научных и инженерных расчётов на Python. Включает инструменты для работы с константами, численными методами, геометрией, теорией чисел и системами координат.

![](./assets/images/mathematics.jpg)

## Install

### Python
```bash
pip install --upgrade git+https://github.com/ParkhomenkoDV/mathematics.git@main
```

### Go
```bash
go get github.com/ParkhomenkoDV/mathematics
```

## Requirements

### Python
- `numpy` — для численных вычислений
- `scipy` — для оптимизации и интегрирования

## Usage

### Consts

```python
from mathematics import Constants

# Создание контейнера с константами
const = Constants(
    SPEED_OF_LIGHT=299792458,  # м/с
    PLANCK=6.62607015e-34,     # Дж·с
    GRAVITY=9.80665,           # м/с²
    freeze=True                # сразу заморозить
)

# Доступ к константам
print(const.SPEED_OF_LIGHT)  # 299792458

# Добавление новой константы
const.add("BOLTZMANN", 1.380649e-23)

# Попытка изменения замороженной константы вызовет ошибку
# const.GRAVITY = 10  # AttributeError!

# Управление заморозкой
const.unfreeze()           # разморозка
const.GRAVITY = 9.81       # теперь можно изменить
const.freeze()             # снова заморозить

# Получение списка констант
print(const.keys())    # ('SPEED_OF_LIGHT', 'PLANCK', 'GRAVITY', 'BOLTZMANN')
print(const.values())  # (299792458, 6.62607015e-34, 9.81, 1.380649e-23)
print(const.items())   # кортеж пар (ключ, значение)
```

### Prefixes SI

```python
from mathematics import prefixes

print(prefixes.kilo.value)   # 1000.0
print(prefixes.micro.value)  # 1e-06
print(prefixes.kilo.degree)  # 3

# Все доступные префиксы
print(prefixes.keys())
# ('yotta', 'zetta', 'exa', 'peta', 'tera', 'giga', 'mega', 'kilo', 
#  'hecto', 'deca', 'deci', 'centi', 'milli', 'micro', 'nano', 
#  'pico', 'femto', 'atto', 'zepto', 'yocto')
```