package integrate

import (
	"errors"
	"math"
)

var (
	ErrNegativeLimit = errors.New("limit must be positive")
	ErrMaxIntervals  = errors.New("maximum number of intervals exceeded")
)

// interval представляет один подынтервал с уже вычисленными значениями функции на концах и в середине.
type interval struct {
	a, b       float64
	fa, fm, fb float64
	S          float64 // приближение Симпсона на этом интервале
	absErr     float64 // оценка абсолютной ошибки
}

// adaptiveSimpson вычисляет определённый интеграл от f на [a,b] с заданной точностью.
// Возвращает Result с оценкой интеграла, погрешностью и числом вычислений функции.
func adaptiveSimpson(
	f func(float64) float64,
	a, b,
	epsabs, epsrel float64,
	limit int) (Result, error) {
	if limit <= 0 {
		return Result{}, ErrNegativeLimit
	}
	if a == b {
		return Result{Value: 0, AbsError: 0, NEval: 0}, nil
	}
	// Перестановка пределов, если нужно
	sign := 1.0
	if a > b {
		a, b = b, a
		sign = -1.0
	}

	// Начальные вычисления
	fa, fm, fb := f(a), f((a+b)/2), f(b)
	S := (b - a) / 6 * (fa + 4*fm + fb)

	// Стек интервалов (LIFO) — обрабатываем в глубину, но это не критично для точности
	stack := make([]interval, 0, 5)
	stack = append(stack, interval{
		a: a, b: b,
		fa: fa, fm: fm, fb: fb,
		S:      S,
		absErr: math.Inf(1), // начальная ошибка завышена, чтобы точно разбить
	})

	total := 0.0    // накопленное значение интеграла
	totalErr := 0.0 // накопленная ошибка
	neval := 3      // уже вызвали fa, fm, fb
	intervalsUsed := 0

	// Пока есть интервалы и не превышен лимит
	for len(stack) > 0 && intervalsUsed < limit {
		// Извлекаем интервал с наибольшей ошибкой (можно использовать кучу, но для простоты берём последний)
		idx := len(stack) - 1
		it := stack[idx]
		stack = stack[:idx]

		// Проверяем, можно ли остановиться на этом интервале
		tol := epsabs + epsrel*math.Abs(it.S)
		if it.absErr <= tol || intervalsUsed+1 >= limit {
			// Принимаем текущее приближение
			total += it.S
			totalErr += it.absErr
			intervalsUsed++
			continue
		}

		// Разбиваем интервал пополам
		mid := (it.a + it.b) / 2
		leftMid, rightMid := (it.a+mid)/2, (mid+it.b)/2

		fl, fr := f(leftMid), f(rightMid)
		neval += 2

		// Симпсон на левой половине
		leftS := (mid - it.a) / 6 * (it.fa + 4*fl + it.fm)
		// Симпсон на правой половине
		rightS := (it.b - mid) / 6 * (it.fm + 4*fr + it.fb)
		S2 := leftS + rightS

		// Оценка ошибки для текущего интервала (по правилу 1/15)
		delta := math.Abs(S2-it.S) / 15.0

		// Если ошибка мала, принимаем сумму половинок
		if delta <= tol {
			total += S2
			totalErr += delta
			intervalsUsed++
			continue
		}

		// Иначе помещаем оба подынтервала в стек
		stack = append(stack,
			interval{ // Для левого
				a: it.a, b: mid,
				fa: it.fa, fm: fl, fb: it.fm,
				S:      leftS,
				absErr: delta, // начальная ошибка для левого (будет уточнена при разбиении)
			},
			interval{ // Для правого
				a: mid, b: it.b,
				fa: it.fm, fm: fr, fb: it.fb,
				S:      rightS,
				absErr: delta,
			},
		)
		// Увеличиваем счётчик использованных интервалов (мы разбили один на два, но ещё не приняли)
		// Вместо этого будем считать принятые интервалы, поэтому здесь не увеличиваем.
	}

	// Если превысили лимит, возвращаем ошибку
	if intervalsUsed >= limit && len(stack) > 0 {
		// Но мы уже накопили частичный результат, можно вернуть его с ошибкой
		return Result{
			Value:    sign * total,
			AbsError: totalErr,
			NEval:    neval,
		}, ErrMaxIntervals
	}

	// Если стек опустел, возвращаем итог
	return Result{
		Value:    sign * total,
		AbsError: totalErr,
		NEval:    neval,
	}, nil
}
