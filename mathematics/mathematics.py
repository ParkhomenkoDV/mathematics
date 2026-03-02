import numpy as np
from numpy import arctan as atan
from numpy import (
    array,
    cos,
    gcd,
    inf,
    isnan,
    matmul,
    nan,
    pi,
    sin,
    sqrt,
    tan,
)
from scipy import integrate
from scipy.optimize import curve_fit


class Constants:
    """
    Класс-контейнер для констант с возможностью:
    - Создания отдельных экземпляров-контейнеров
    - Заморозки/разморозки констант (для защиты от изменений после заморозки)
    """

    def __init__(self, *, freeze: bool = False, **kwargs) -> None:
        """
        Инициализация контейнера констант.
        Принимает:
        - freeze: флаг заморозки констант сразу после инициализации
        - начальные константы как именованные аргументы
        """
        object.__setattr__(self, "_frozen", False)
        for name, value in kwargs.items():
            self.add(name, value)

        if freeze:
            self.freeze()

    def add(self, name: str, value) -> None:
        """
        Добавляет новую константу.
        Вызывает AttributeError, если константа уже существует и контейнер заморожен.
        """
        if self._frozen:
            raise AttributeError(f'Constant "{name}" cannot be modified (container is frozen)')
        object.__setattr__(self, name, value)

    def freeze(self) -> None:
        """Замораживает контейнер, запрещая изменения констант"""
        object.__setattr__(self, "_frozen", True)

    def unfreeze(self) -> None:
        """Размораживает контейнер, разрешая изменения"""
        object.__setattr__(self, "_frozen", False)

    def is_frozen(self) -> bool:
        """Проверяет, заморожен ли контейнер"""
        return self._frozen

    def __setattr__(self, name, value) -> None:
        if name == "_frozen":
            object.__setattr__(self, name, value)
        elif self._frozen:
            raise AttributeError(f'Constant "{name}" cannot be modified (container is frozen)')
        else:
            self.add(name, value)

    def __delattr__(self, name) -> None:
        if self._frozen:
            raise AttributeError(f'Constant "{name}" cannot be deleted (container is frozen)')
        object.__delattr__(self, name)

    def __repr__(self):
        constants = {k: v for k, v in self.__dict__.items() if k != "_frozen"}
        return f"Constants(frozen={self._frozen}, values={constants})"

    def keys(self) -> tuple:
        return tuple(key for key in self.__dict__.keys() if not key.startswith("_"))

    def values(self) -> tuple:
        return tuple(value for key, value in self.__dict__.items() if not key.startswith("_"))

    def items(self) -> tuple:
        return tuple((key, value) for key, value in self.__dict__.items() if not key.startswith("_"))


prefixes = Constants(
    yotta=Constants(degree=+24, value=10**+24, freeze=True),
    zetta=Constants(degree=+21, value=10**+21, freeze=True),
    exa=Constants(degree=+18, value=10**+18, freeze=True),
    peta=Constants(degree=+15, value=10**+15, freeze=True),
    tera=Constants(degree=+12, value=10**+12, freeze=True),
    giga=Constants(degree=+9, value=10**+9, freeze=True),
    mega=Constants(degree=+6, value=10**+6, freeze=True),
    kilo=Constants(degree=+3, value=10**+3, freeze=True),
    hecto=Constants(degree=+2, value=10**+2, freeze=True),
    deca=Constants(degree=+1, value=10**+1, freeze=True),
    deci=Constants(degree=-1, value=10**-1, freeze=True),
    centi=Constants(degree=-2, value=10**-2, freeze=True),
    milli=Constants(degree=-3, value=10**-3, freeze=True),
    micro=Constants(degree=-6, value=10**-6, freeze=True),
    nano=Constants(degree=-9, value=10**-9, freeze=True),
    pico=Constants(degree=-12, value=10**-12, freeze=True),
    femto=Constants(degree=-15, value=10**-15, freeze=True),
    atto=Constants(degree=-18, value=10**-18, freeze=True),
    zepto=Constants(degree=-21, value=10**-21, freeze=True),
    yocto=Constants(degree=-24, value=10**-24, freeze=True),
)


def derivative(
    f: callable,
    x0: int | float | np.number | np.ndarray,
    method: str = "central",
    dx: float | np.floating = 1e-6,
) -> float:
    """
    Производная функции f в точке x0

    Parameters
    ----------
    f : function
        Vectorized function of one variable
    x0 : number
        Compute derivative at x = a
    method : string
        Difference formula:
        'central': (f(a+h) - f(a-h)) / 2h
        'forward': (f(a+h) - f(a)) / h
        'backward': (f(a) - f(a-h)) / h

    dx : number
        Step size in difference formula
    """
    assert isinstance(dx, float) and dx != 0

    if method == "central":
        return (f(x0 + dx / 2) - f(x0 - dx / 2)) / dx
    elif method == "forward":
        return (f(x0 + dx) - f(x0)) / dx
    elif method == "backward":
        return (f(x0) - f(x0 - dx)) / dx
    else:
        raise ValueError('method must be "central", "forward" or "backward"!')


def approximate(function, x, y) -> tuple:
    """Подбор параметров переданной ф-и по точкам"""

    """
    # Исходные данные
    x = array([1, 2, 3, 4, 5])
    y = array([1, 4, 9, 16, 25])

    def func(x, a, b, c): return a * x ** 2 + b * x + c  # функция для аппроксимации

    popt, covar = curve_fit(func, x, y)  # Аппроксимация данных
    
    '''
    Массив popt содержит оптимальные значения параметров функции, 
    которые были найдены в результате аппроксимации данных. 
    
    Массив covar содержит ковариационную матрицу, 
    которая показывает, насколько точно определены значения параметров. 
    Если диагональные элементы этой матрицы близки к нулю, 
    то это означает, что значения параметров определены с высокой точностью. 
    Если же диагональные элементы большие, 
    то это может указывать на неопределенность в значениях параметров.'''

    # График исходных данных и аппроксимационной кривой
    plt.plot(x, y, 'o')
    plt.plot(linspace(min(x), max(x), 100), func((linspace(min(x), max(x), 100), *popt), '-')
    plt.show()
    """
    assert callable(function)
    assert isinstance(x, (tuple, list, np.ndarray))
    assert isinstance(y, (tuple, list, np.ndarray))
    popt, covar = curve_fit(function, x, y)  # оптимальные парамеры ф-и и ковариация
    return popt, covar


def polynomial(x: tuple | list | np.ndarray, y: tuple | list | np.ndarray, deg: int | np.integer):
    """Коэффициенты полинома степени deg"""
    assert isinstance(x, (tuple, list, np.ndarray))
    assert isinstance(y, (tuple, list, np.ndarray))
    assert len(x) == len(y)
    assert isinstance(deg, int)
    return np.polyfit(x, y, deg)


def distance(point1: tuple, point2: tuple) -> float:
    """Декартово расстояние между 2D точками"""
    assert isinstance(point1, (tuple, list)) and isinstance(point2, (tuple, list))
    assert len(point1) == len(point2)
    return float(np.linalg.norm(array(point1) - array(point2)))


def distance2line(point: tuple, ABC: tuple) -> float:
    """Расстояние от точки до прямой"""
    assert isinstance(point, (tuple, list, np.ndarray))
    assert isinstance(ABC, (tuple, list, np.ndarray))
    assert len(point) == 2
    assert len(ABC) == 3

    A, B, C = ABC
    x, y = point

    return abs(A * x + B * y + C) / sqrt(A**2 + B**2)


def coefficients_line(func=None, x0=None, p1=None, p2=None) -> tuple[float, float, float]:
    """Коэффициенты A, B, C касательной в точке x0 кривой f или прямой, проходящей через точки p1 и p2"""
    if func is not None and x0 is not None:
        df_dx = derivative(func, x0)
        return df_dx, -1, func(x0) - df_dx * x0
    elif p1 is not None and p2 is not None:
        return (
            (p2[1] - p1[1]) / (p2[0] - p1[0]),
            -1,
            (p2[0] * p1[1] - p1[0] * p2[1]) / (p2[0] - p1[0]),
        )
    else:
        raise ValueError("func, x0, p1, p2 must not be None!")


def coordinate_intersection_lines(ABC1: tuple | list | np.ndarray, ABC2: tuple | list | np.ndarray) -> tuple[float, float]:
    """Точка пересечения прямых с коэффициентами ABC1 = (A1, B1, C1) и ABC2 = (A2, B2, C2)"""
    assert isinstance(ABC1, (tuple, list, np.ndarray)) and isinstance(ABC2, (tuple, list, np.ndarray))
    assert all(isinstance(el, (int, float, np.number)) for el in ABC1)
    assert all(isinstance(el, (int, float, np.number)) for el in ABC2)

    A1, B1, C1 = ABC1
    A2, B2, C2 = ABC2

    D = A1 * B2 - A2 * B1  # детерминант системы уравнений

    if D == 0 or isnan(D):  # параллельность или совпадение прямых
        x, y = nan, nan  # None
    else:  # общий случай пересечения прямых
        x = (C2 * B1 - C1 * B2) / D
        y = (C1 * A2 - C2 * A1) / D
    return x, y


def angle_between(k1=nan, k2=nan, points=(tuple(), tuple(), tuple())) -> float:
    """Острый угол [рад] между прямыми"""
    if all(points):
        p0, p1, p2 = points  # разархивирование точек
        k1 = (p0[1] - p1[1]) / (p0[0] - p1[0]) if (p0[0] - p1[0]) != 0 else inf
        k2 = (p1[1] - p2[1]) / (p1[0] - p2[0]) if (p1[0] - p2[0]) != 0 else inf
    return abs(atan((k2 - k1) / (1 + k1 * k2)))


def cot(x: float | int) -> float:
    """Котангенс"""
    return 1 / tan(x) if x != 0 else inf


def tan2cos(tg: float | int | np.number | np.ndarray):
    """Преобразование тангенса в косинус"""
    return sqrt(1 / (tg**2 + 1))


def cot2sin(ctg: float | int | np.number | np.ndarray):
    """Преобразование котангенса в синус"""
    return sqrt(1 / (ctg**2 + 1))


def tan2sin(tg: float | int | np.number | np.ndarray):
    """Преобразование тангенса в синус"""
    return tg * tan2cos(tg)


def cot2cos(ctg: float | int | np.number | np.ndarray):
    """Преобразование котангенса в косинус"""
    return ctg * cot2sin(ctg)


def sum_atan(a1, a2):
    """atan(a1) + atan(a2)"""
    if a1 * a2 < 1:
        return atan((a1 + a2) / (1 - a1 * a2))
    elif a1 > 0 and a1 * a2 > 1:
        return pi + atan((a1 + a2) / (1 - a1 * a2))
    elif a1 < 0 and a1 * a2 > 1:
        return -pi + atan((a1 + a2) / (1 - a1 * a2))


def discriminant(a, b, c) -> float:
    """Дискриминант"""
    assert isinstance(a, (int, float, np.number))
    assert isinstance(b, (int, float, np.number))
    assert isinstance(c, (int, float, np.number))
    return b**2 - 4 * a * c


def quadratic_equation(a, b, c) -> tuple[float, float] | float | None:
    """Решение квадратного уравнения"""
    d = discriminant(a, b, c)  # assert внутри
    if d > 0:
        return (-b - sqrt(d)) / (2 * a), (-b + sqrt(d)) / (2 * a)
    elif d == 0:
        return -b / (2 * a)
    else:
        return None


def is_coprime(a: int, b: int) -> bool:
    """Проверка на взаимно простые числа"""
    return gcd(a, b) == 1


def is_prime(n: int) -> bool:
    """Проверка на простое число"""
    if not isinstance(n, int):
        return False
    if n < 2:
        return False
    for i in range(2, int(sqrt(n)) + 1):
        if n % i == 0:
            return False
    return True


def prime_factorization(n: int, repeat: bool = True, sort: bool = False) -> list[int]:
    """Разложение на простые множители"""
    result, divisor = list(), 2
    while divisor * divisor <= n:
        if n % divisor == 0:
            result.append(divisor)
            n //= divisor
        else:
            divisor += 1
    if n > 1:
        result.append(n)
    if not repeat:
        result = list(set(result))
    if sort:
        result.sort(reverse=False)
    return result


def eps(type_eps: str, x1: float | int | np.number, x2: float | int | np.number) -> float:
    """Погрешность"""
    if type_eps == "rel":
        try:
            return (x2 - x1) / x1  # относительная
        except ZeroDivisionError:
            return inf
    elif type_eps == "abs":
        return x2 - x1  # абсолютная
    else:
        raise ValueError('type_eps must be "rel" or "abs"!')


def integral_average(function, *ranges) -> tuple[float, float]:
    """Среднее интегральное"""
    assert callable(function), TypeError(f"function {function} must be callable")
    func_args = function.__code__.co_varnames
    assert len(func_args) == len(ranges), ArithmeticError(f"count function args {len(func_args)} must be equal count borders {len(ranges)}")

    denominator: float = 1.0  # знаменатель = произведение разниц границ интегрирования
    for rang in ranges:
        assert isinstance(rang, (tuple, list, np.ndarray)), TypeError(f"type range must be tuple, but has {type(rang)}")
        assert len(rang) == 2, ArithmeticError(f"integral has only 2 borders, but has {len(rang)}")
        assert all(map(lambda x: isinstance(x, (float, int, np.number)), rang)), TypeError(f"type of ranges must be number, but has {type(rang[0]), type(rang[1])}")
        denominator *= rang[1] - rang[0]

    result, abserr = integrate.nquad(func=function, ranges=ranges, full_output=False)

    if denominator == 0:  # деление на 0
        return nan, nan

    return result / denominator, abserr


class Axis:
    """Система координат"""

    @staticmethod
    def to_cartesian(r: int | float | np.number, a: int | float | np.number) -> tuple[float, float]:
        """Преобразование в декартову СК"""
        return r * cos(a), r * sin(a)

    @staticmethod
    def to_polar(x: int | float | np.number, y: int | float | np.number) -> tuple[float, float]:
        """Преобразование в полярную СК"""
        return distance(point1=(0, 0), point2=(x, y)), atan(y / x)

    @staticmethod
    def transform(
        coordinates: tuple | list | np.ndarray,
        transfer: tuple | list | np.ndarray = (0, 0),
        angle: float | int | np.number = 0,
        scale: float | int | np.number = 1,
        dtype: str = "float64",
    ):
        """Перенос-поворот-масштабирование осей против часовой стрелки"""
        coordinates = array(coordinates, dtype=dtype)
        transfer = array(transfer, dtype=dtype)
        c, s = cos(angle), sin(angle)
        rotation_matrix = array(((c, -s), (s, c)), dtype=dtype)

        return scale * matmul(coordinates - transfer, rotation_matrix)

    @staticmethod
    def mirror(x: int | float | np.number, y: int | float | np.number):
        pass  # TODO


class Angle:
    UNITS = {
        "rad": ("", "rad", "radians"),
        "deg": ("deg", "degrees", "°"),
        "str": ("str",),
    }

    __slots__ = ("__angle",)

    def __init__(self, angle: str, unit: str, normalize: bool = True):
        assert isinstance(angle, str)

        assert isinstance(unit, str)
        unit = unit.strip().lower()

        assert isinstance(normalize, bool)

        if unit in Angle.UNITS["rad"]:
            self.__angle = float(angle) % (2 * pi) if normalize else float(angle) % (2 * pi)
        elif unit in Angle.UNITS["deg"]:
            self.__angle = (float(angle) * pi / 180) % 360 if normalize else (float(angle) * pi / 180)
        elif unit in ("str",):
            angle = angle
        else:
            raise Exception(f"unit {unit} not in {Angle.UNITS.values()}")

    @property
    def rad(self) -> float:
        return self.__angle

    @property
    def deg(self) -> float:
        return self.__angle * 180 / pi

    @property
    def minutes(self) -> float:
        return self.deg / 60

    @property
    def seconds(self) -> float:
        return self.deg / 3600

    @property
    def str(self) -> str:
        return f"{self.deg}°{self.minutes * 60}'{self.seconds * 3600}''"


if __name__ == "__main__":
    # Создаем экземпляр с начальными константами
    app_consts = Constants(PI=3.14159, MAX_USERS=100)

    print(app_consts.PI)  # 3.14159

    # Добавляем новую константу
    app_consts.add("TIMEOUT", 30)

    # Замораживаем
    app_consts.freeze()

    print(app_consts.keys())
    print(app_consts.values())
    print(app_consts.items())

    try:
        app_consts.PI = 3.14  # Ошибка!
    except AttributeError as e:
        print(e)  # Constant 'PI' cannot be modified (container is frozen)

    # Создаем другой независимый контейнер
    db_consts = Constants(DB_HOST="localhost", DB_PORT=5432)
    db_consts.add("DB_NAME", "my_db")
