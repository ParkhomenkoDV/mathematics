import numpy as np
import pytest
from numpy import inf, isinf, isnan, nan

from mathematics import discriminant, eps, integral_average


class TestDiscriminant:
    """Тесты для функции вычисления дискриминанта"""

    # Параметризованные тесты для проверки корректных вычислений
    @pytest.mark.parametrize(
        "a, b, c, expected",
        [
            # Стандартные случаи
            (1, 5, 6, 1),  # D = 25 - 24 = 1
            (2, 4, 2, 0),  # D = 16 - 16 = 0
            (1, 0, -4, 16),  # D = 0 - (-16) = 16
            (0.5, 3, 2, 5),  # D = 9 - 4 = 5 (дробные коэффициенты)
            # Крайние значения
            (1e-10, 2e5, 3e15, 4e10 - 12e5),  # Очень большие/малые числа
            (-1, -2, -3, -8),  # Отрицательные коэффициенты (D = 4 - 12 = -8)
        ],
    )
    def test_discriminant(self, a, b, c, expected):
        """Проверка корректности вычислений"""
        assert discriminant(a, b, c) == pytest.approx(expected)

    @pytest.mark.parametrize(
        "a, b, c",
        [
            # Стандартные случаи
            (1, 5, 6),  # D = 25 - 24 = 1
            (2, 4, 2),  # D = 16 - 16 = 0
            (1, 0, -4),  # D = 0 - (-16) = 16
            (0.5, 3, 2),  # D = 9 - 4 = 5 (дробные коэффициенты)
            # Крайние значения
            (1e-10, 2e5, 3e15),  # Очень большие/малые числа
            (-1, -2, -3),  # Отрицательные коэффициенты (D = 4 - 12 = -8)
        ],
    )
    @pytest.mark.benchmark
    def test_discriminant_calculation(self, benchmark, a, b, c):
        """Бенчмарк вычислений"""

        def benchfunc():
            return discriminant(a, b, c)

        benchmark(benchfunc)

    # Тесты для numpy чисел
    def test_numpy_input(self):
        """Проверка работы с numpy-числами"""
        a = np.float64(1.5)
        b = np.int32(3)
        c = np.array([2])[0]  # numpy scalar

        result = discriminant(a, b, c)
        assert result == pytest.approx(-3.0)
        assert isinstance(result, float)

    # Тесты на обработку ошибок
    @pytest.mark.parametrize(
        "a, b, c",
        [
            ("1", 2, 3),  # Строка вместо числа
            (1, [2], 3),  # Список вместо числа
            (1, 2, None),  # None вместо числа
            (1 + 2j, 2, 3),  # Комплексное число
        ],
    )
    def test_invalid_input_types(self, a, b, c):
        """Проверка обработки нечисловых аргументов"""
        with pytest.raises(AssertionError):
            discriminant(a, b, c)

    # Тест на возвращаемый тип
    def test_return_type(self):
        """Проверка что возвращается float"""
        result = discriminant(1, 5, 6)
        assert isinstance(result, (float, int))


class TestEps:
    """Тесты для функции eps()"""

    # Тесты для абсолютной погрешности (abs)
    def test_abs_positive(self):
        assert eps("abs", 10, 12) == 2.0

    def test_abs_negative(self):
        assert eps("abs", 10, 8) == -2.0

    def test_abs_zero_diff(self):
        assert eps("abs", 10, 10) == 0.0

    def test_abs_with_numpy(self):
        assert eps("abs", np.float64(10), np.int32(12)) == 2.0

    # Тесты для относительной погрешности (rel)
    def test_rel_positive(self):
        assert eps("rel", 10, 12) == 0.2  # (12-10)/10

    def test_rel_negative(self):
        assert eps("rel", 10, 8) == -0.2  # (8-10)/10

    def test_rel_zero_diff(self):
        assert eps("rel", 10, 10) == 0.0

    def test_rel_zero_base(self):
        assert eps("rel", 0, 10) == inf

    def test_rel_small_values(self):
        assert pytest.approx(eps("rel", 1e-10, 1.1e-10)) == 0.1

    def test_rel_with_numpy(self):
        assert eps("rel", np.float64(10), np.int32(12)) == 0.2

    # Тесты на обработку ошибок
    def test_invalid_type_eps(self):
        with pytest.raises(ValueError) as excinfo:
            eps("invalid", 10, 12)
        assert '"rel" or "abs"' in str(excinfo.value)

    def test_non_numeric_x1(self):
        with pytest.raises(TypeError):
            eps("abs", "not_number", 12)

    def test_non_numeric_x2(self):
        with pytest.raises(TypeError):
            eps("abs", 10, "not_number")

    # Тесты на граничные значения
    @pytest.mark.skip
    def test_abs_large_numbers(self):
        assert eps("abs", 1e100, 1e100 + 1) == pytest.approx(1.0)

    def test_rel_large_numbers(self):
        assert eps("rel", 1e100, 1.1e100) == pytest.approx(0.1)

    def test_abs_small_numbers(self):
        assert eps("abs", 1e-100, 1.1e-100) == pytest.approx(1e-101)

    @pytest.mark.parametrize("type_eps", ["abs", "rel"])
    def test_inf_values(self, type_eps):
        assert isnan(eps(type_eps, inf, inf))
        assert isnan(eps(type_eps, -inf, -inf))
        if type_eps == "abs":
            assert isinf(eps(type_eps, 0, inf))
            assert isinf(eps(type_eps, inf, 0))
        else:
            assert isinf(eps(type_eps, 0, inf))
            assert isnan(eps(type_eps, inf, 0))


class TestIntegralAverage:
    """Тесты для функции integral_average"""

    maxError = 1e-9
    rel = 0.000_1

    @staticmethod
    def f0_arg1(x):
        return 5.0  # 5 * x

    @staticmethod
    def f1_arg1(x):
        return 2 * x + 1  # x**2 + x

    @staticmethod
    def f2_arg2(x, y):
        return x**2 + y**2

    @staticmethod
    def f1_arg3(x, y, z):
        return x + y + z

    @pytest.mark.parametrize(
        "ranges, expected",
        [
            ([(0, 2)], 5.0),
            ([(0, 0)], nan),
        ],
    )
    def test_integral_average_function_0(self, ranges, expected):
        """Тест для константной функции"""
        result, error = integral_average(self.f0_arg1, *ranges)
        if isnan(expected):
            assert isnan(result)
        else:
            assert result == pytest.approx(expected, rel=self.rel)
            assert error <= self.maxError
        assert isinstance(result, float)
        assert isinstance(error, float)

    @pytest.mark.parametrize(
        "ranges, expected",
        [
            ([(0, 1)], (1**2 + 1) / 1),
            ([(2, 4)], ((4**2 + 4) - (2**2 + 2)) / (4 - 2)),
            ([(0, 0)], nan),
        ],
    )
    def test_integral_average_basic_functionality(self, ranges, expected):
        """Тест базовой функциональности для простой функции"""
        result, error = integral_average(self.f1_arg1, *ranges)
        if isnan(expected):
            assert isnan(result)
        else:
            assert result == pytest.approx(expected, rel=self.rel)
            assert error <= self.maxError
        assert isinstance(result, float)
        assert isinstance(error, float)

    @pytest.mark.parametrize(
        "ranges, expected",
        [
            ([(0, 1), (0, 1)], (1 / 3 + 1 / 3) / 1),
            ([(0, 0), (0, 1)], nan),
            ([(0, 1), (0, 0)], nan),
            ([(0, 0), (0, 0)], nan),
        ],
    )
    def test_integral_average_2d_function(self, ranges, expected):
        """Тест для двумерной функции"""
        result, error = integral_average(self.f2_arg2, *ranges)
        if isnan(expected):
            assert isnan(result)
        else:
            assert result == pytest.approx(expected, rel=self.rel)
            assert error <= self.maxError
        assert isinstance(result, float)
        assert isinstance(error, float)

    @pytest.mark.parametrize(
        "ranges, expected",
        [
            ([(0, 1), (0, 1), (0, 1)], (0.5 + 0.5 + 0.5) / 1),
            ([(0, 0), (0, 1), (0, 1)], nan),
            ([(0, 1), (0, 0), (0, 1)], nan),
            ([(0, 1), (0, 1), (0, 0)], nan),
            ([(0, 0), (0, 0), (0, 0)], nan),
        ],
    )
    def test_integral_average_3d_function(self, ranges, expected):
        """Тест для трехмерной функции"""
        result, error = integral_average(self.f1_arg3, *ranges)
        if isnan(expected):
            assert isnan(result)
        else:
            assert result == pytest.approx(expected, rel=self.rel)
            assert error <= self.maxError
        assert isinstance(result, float)
        assert isinstance(error, float)

    def test_integral_average_error_cases(self):
        """Тест обработки ошибок"""

        # Не callable объект
        with pytest.raises((AssertionError, TypeError)):
            integral_average(5, (0, 1))

        # Неверное количество аргументов для интегрирования
        with pytest.raises((AssertionError, ArithmeticError)):
            integral_average(lambda x, y: x * y, (0, 1))

        # Неправильный тип диапазона
        with pytest.raises((AssertionError, TypeError)):
            integral_average(lambda x: x, "not_a_tuple")

        # Диапазон не из 2 элементов
        with pytest.raises(AssertionError):
            integral_average(lambda x: x, (0, 1, 2))

        # Нечисловые границы
        with pytest.raises(AssertionError):
            integral_average(lambda x: x, ("a", "b"))

    def test_integral_average_zero_range(self):
        """Тест с нулевым диапазоном (деление на ноль должно обрабатываться scipy)"""
        result, _ = integral_average(lambda x: 5 * x, (0, 0))
        assert isnan(result)

    def test_integral_average_complex_function(self):
        """Тест с более сложной функцией"""

        def complex_func(x):
            return np.sin(x) + np.cos(x)

        result, error = integral_average(complex_func, (0, np.pi / 2))
        expected = (-np.cos(np.pi / 2) + np.sin(np.pi / 2) + np.cos(0) - np.sin(0)) / (np.pi / 2)
        assert abs(result - expected) < 1e-10
        assert error <= self.maxError

    def test_integral_average_return_types(self):
        """Тест типов возвращаемых значений"""

        def simple_func(x):
            return x

        result, error = integral_average(simple_func, (0, 1))
        assert isinstance(result, float)
        assert isinstance(error, float)
