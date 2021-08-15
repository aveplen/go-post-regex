# Простейший в мире\* компилятор регулярных выражений

У данного движка (пока что) только одна функция - взять из командной строки регулярное выражение, скомпилировать его в недетерминированный конечный автомат (графовое представление), и проверить, допускает ли этот автомат строку, переданную в той же командной строке в качестве второго аргумента.



# Алгоритм разбора регулярного выражения

Сие чудо допускает использование операторов конкатенации ```.``` (точка), объединения ```|``` (лог. или), звезды Клини ```*``` (знак умножения), плюса Клини ```+``` (плюс) и модификатора нуля или одного вхождения ```?``` (знак вопроса).

Для перевода заданного регулярного выражения в удобную для обработки форму используется [обратная польская запись](https://ru.wikipedia.org/wiki/%D0%9E%D0%B1%D1%80%D0%B0%D1%82%D0%BD%D0%B0%D1%8F_%D0%BF%D0%BE%D0%BB%D1%8C%D1%81%D0%BA%D0%B0%D1%8F_%D0%B7%D0%B0%D0%BF%D0%B8%D1%81%D1%8C), поэтому алгоритм выглядит следующим образом:

1. Везде, где это необходимо, вставить знак принудительной конкатенации (потому что операция перехода по строке для конечного автомата в общем случае определяется через цепочку переходов по буквам, из которых состоит строка, т.е. ```a(b|cd)*f = a.(b|c.d)*.f```).

2. Перевести полученное выражение в постфиксную запись (она же обратная польская нотация) с использованием стека (может есть и другие алгоритмы, которые подразумевают использование других структур данных, но о них не известно).

3. Построить конечный автомат из результата шага 2 с использованием стека (в этот раз в стеке будут лежать не ASCII символы строки, а готовые блоки, из которых прямо в этом стеке собирается автомат).

4. Произвести проход по автомату (по сути это [BFS](https://ru.wikipedia.org/wiki/%D0%9F%D0%BE%D0%B8%D1%81%D0%BA_%D0%B2_%D1%88%D0%B8%D1%80%D0%B8%D0%BD%D1%83) или [поиск в ширину](https://ru.wikipedia.org/wiki/%D0%9F%D0%BE%D0%B8%D1%81%D0%BA_%D0%B2_%D1%88%D0%B8%D1%80%D0%B8%D0%BD%D1%83), так как автомат в нашем случае представлен графом состояний с рёбрами, которые либо не имеют никакого веса, либо взвешены ASCII символами). Стоит также отметить, что при наличии в регулярном выражении ```+``` или ```*``` автомат гарантированно будет иметь циклы, с которыми надо как-то мириться, поэтому предлагается алгоритм обхода с детерминизацией "на лету".



# Из каких блоков строится автомат?

Самым базовым блоком является автомат, который принимает пустую строку.
```
     ┌───┐   eps   ╔═══╗
────>│   ├────────>║   ║
     └───┘         ╚═══╝


func FromEpsilon() *NFA {
	from := state.New(false)
	to := state.New(true)
	from.AddEpsilon(to)
	return &NFA{
		Start: from,
		End:   to,
	}
}
```

---

Следующим блоком выступает автомат, допускающий конкретную букву:
```
     ┌───┐    a    ╔═══╗
────>│   ├────────>║   ║
     └───┘         ╚═══╝


func FromSymbol(symbol rune) *NFA {
	from := state.New(false)
	to := state.New(true)
	from.AddTransition(to, symbol)
	return &NFA{
		Start: from,
		End:   to,
	}
}
```

---

Конкатенации в терминах автоматов выглядит как цепочка из двух блоков:
```

   ╭───────────────╮   ╭───────────────╮
   │ ┌───┐   ┌───┐ │eps│ ┌───┐   ╔═══╗ │
───┼>│   │   │   ├─┼───┼>│   │   ║   ║ │
   │ └───┘   └───┘ │   │ └───┘   ╚═══╝ │
   ╰───────────────╯   ╰───────────────╯


func Concat(first *NFA, second *NFA) *NFA {
	first.End.AddEpsilon(second.Start)
	first.End.IsEnd = false
	return &NFA{
		Start: first.Start,
		End:   second.End,
	}
}
```

---

Объединение:
```
                  ╭───────────────╮
              eps │ ┌───┐   ┌───┐ │eps
             ┌────┼>│   │   │   ├─┼───┐
             │    │ └───┘   └───┘ │   │
     ┌───┐   │    ╰───────────────╯   │    ╔═══╗
────>│   ├───┤                        ├───>║   ║
     └───┘   │    ╭───────────────╮   │    ╚═══╝
             │    │ ┌───┐   ┌───┐ │   │
             └────┼>│   │   │   ├─┼───┘
              eps │ └───┘   └───┘ │eps
                  ╰───────────────╯


func Union(first *NFA, second *NFA) *NFA {
	res := New()
	res.Start = state.New(false)
	res.End = state.New(true)

	res.Start.AddEpsilon(first.Start)
	res.Start.AddEpsilon(second.Start)

	first.End.IsEnd = false
	first.End.AddEpsilon(res.End)

	second.End.IsEnd = false
	second.End.AddEpsilon(res.End)

	return res
}
```

---

Замыкание (звезда Клини):
```
               eps
       ┌─────────────────┐
       v     ╭───────────┼───╮
     ┌───┐eps│ ┌───┐   ┌─┴─┐ │eps ╔═══╗
────>│   ├───┼>│   │   │   ├─┼───>║   ║
     └─┬─┘   │ └───┘   └───┘ │    ╚═══╝
       │     ╰───────────────╯      ^
       └────────────────────────────┘
                   eps


func Closure(nfa *NFA) *NFA {
	res := New()
	res.Start = state.New(false)
	res.End = state.New(true)

	res.Start.AddEpsilon(res.End)
	res.Start.AddEpsilon(nfa.Start)

	nfa.End.AddEpsilon(res.End)
	nfa.End.AddEpsilon(nfa.Start)
	nfa.End.IsEnd = false

	return res
}
```

---

Усечённое замыкание (плюс Клини):
```
               eps
       ┌─────────────────┐
       v     ╭───────────┼───╮
     ┌───┐eps│ ┌───┐   ┌─┴─┐ │eps ╔═══╗
────>│   ├───┼>│   │   │   ├─┼───>║   ║
     └───┘   │ └───┘   └───┘ │    ╚═══╝
             ╰───────────────╯


func Plus(nfa *NFA) *NFA {
	res := New()
	res.Start = state.New(false)
	res.End = state.New(true)

	// res.Start.AddEpsilon(res.End)
	res.Start.AddEpsilon(nfa.Start)

	nfa.End.AddEpsilon(res.End)
	nfa.End.AddEpsilon(nfa.Start)
	nfa.End.IsEnd = false

	return res
}
```

---

Опциональное вхождение (знак вопроса? - не знаю, как еще это назвать):
```
             ╭───────────────╮
     ┌───┐eps│ ┌───┐   ┌───┐ │eps ╔═══╗
────>│   ├───┼>│   │   │   ├─┼───>║   ║
     └─┬─┘   │ └───┘   └───┘ │    ╚═══╝
       │     ╰───────────────╯      ^
       └────────────────────────────┘
                   eps


func Optional(nfa *NFA) *NFA {
	res := New()
	res.Start = state.New(false)
	res.End = state.New(true)

	res.Start.AddEpsilon(res.End)
	res.Start.AddEpsilon(nfa.Start)

	nfa.End.AddEpsilon(res.End)
	nfa.End.IsEnd = false

	return res
}
```
# Как произвести обход полученного графа за относительно малое время?

Алгоритм быстрого обхода состоит в том, чтобы исполнитель находился в нескольких состояниях сразу. Для этого достаточно воспользоваться простым трюком:
**вместо учёта состояния, в котором находится исполнитель, будем следить сразу за всем множеством состояний, в которых исполнитель мог бы оказаться.** Например, если у нас возникла ситуация, в которой из состояния ```Q1``` мы можем попасть либо в ```Q2```, либо в ```Q3```, мы не должны выбирать что-то одно, а просто добавить в наше множество текущих состояний обе опции:

```{Q1} -> {Q2, Q3}```


# Основная задача

Таким образом, мы свели исходную задачу: "проверить, удовлетворяет ли данная строка заданному регулярному выражению" сводится к тому, чтобы проверить, достижима ли конечная (допускающая) вершина из начальной вершины по истечению данной строки в построенном автомате. Просто проходимся по строке, пока она не иссякнет, и проверяем, попала ли в наше множество состояний конечная вершина. (алгоритм конструирования автомата из блоков гарантирует, что конечная вершина только одна)

> \*предположение сделано на основе скудности функционала представленного ПО
