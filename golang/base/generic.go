package main

import (
	"encoding/json"
	"fmt"
)

/*
Any 泛型
Go 1.18版本增加了对泛型的支持，泛型也是自 Go 语言开源以来所做的最大改变。
什么是泛型
泛型允许程序员在强类型程序设计语言中编写代码时使用一些以后才指定的类型，在实例化时作为参数指明这些类型。
换句话说，在编写某些代码或数据结构时先不提供值的类型，而是之后再提供。
泛型是一种独立于所使用的特定类型的编写代码的方法。使用泛型可以编写出适用于一组类型中的任何一种的函数和类型。

泛型为Go语言添加了三个新的重要特性:
1、函数和类型的类型参数。
2、将接口类型定义为类型集，包括没有方法的类型。
3、类型推断，它允许在调用函数时在许多情况下省略类型参数。
*/
func Any() {
	fmt.Println("============ 泛型 ============")
	/*
		借助泛型，我们可以声明一个适用于一组类型的min函数。
	*/
	//注意：a和b的类型一定要是一样的才行。比如：a是int，b一定是int，不能b为float64。【反推：a是float64，b一定是float64】
	minNum := minAny(1, 2)      //int
	fmt.Println(minNum)         //1
	minNum2 := minAny(1.2, 2.0) //float64
	fmt.Println(minNum2)        //1.2

	/*
		类型参数的使用
		type Slice[T int | string] []T

		type Map[K int | string, V float32 | float64] map[K]V

		type Tree[T interface{}] struct {
			left, right *Tree[T]
			value       T
		}

		在上述泛型类型中，T、K、V都属于类型形参，类型形参后面是类型约束，类型实参需要满足对应的类型约束。
		泛型类型可以有方法，例如为上面的Tree实现一个查找元素的Lookup方法。

			func (t *Tree[T]) Lookup(x T) *Tree[T] { ... }

		要使用泛型类型，必须进行实例化。Tree[string]是使用类型实参string实例化 Tree 的示例。

			var stringTree Tree[string]
	*/

	/*
		类型约束
		普通函数中的每个参数都有一个类型; 该类型定义一系列值的集合。
		例如，我们上面定义的非泛型函数minFloat64那样，声明了参数的类型为float64，那么在函数调用时允许传入的实际参数就必须是可以用float64类型表示的浮点数值。
		类似于参数列表中每个参数都有对应的参数类型，类型参数列表中每个类型参数都有一个类型约束。类型约束定义了一个类型集——只有在这个类型集中的类型才能用作类型实参。
		Go语言中的类型约束是接口类型。
		就以上面提到的min函数为例，我们来看一下类型约束常见的两种方式。
		类型约束接口可以直接在类型参数列表中使用。
		在使用类型约束时，如果省略了外层的interface{}会引起歧义，那么就不能省略。例如：

			type IntPtrSlice [T *int] []T  // T*int ?
			type IntPtrSlice[T *int,] []T  // 只有一个类型约束时可以添加`,`
			type IntPtrSlice[T interface{ *int }] []T // 使用interface{}包裹
	*/

	/*
		any接口（type any = interface{}）
		空接口在类型参数列表中很常见，在Go 1.18引入了一个新的预声明标识符，作为空接口类型的别名。
	*/
	//foo2()

	/*
		类型推断
		最后一个新的主要语言特征是类型推断。从某些方面来说，这是语言中最复杂的变化，但它很重要，因为它能让人们在编写调用泛型函数的代码时更自然。
	*/
	minAny(1, 2.0) //1，2.0由编译器推断类型。在许多情况下，编译器可以从普通参数推断 T 的类型实参。这使得代码更短，同时保持清晰。

	/*
		约束类型推断
		Go 语言支持另一种类型推断，即约束类型推断。
	*/
	//Scale2(1,2)

	/*
		总结
		总之，如果你发现自己多次编写完全相同的代码，而这些代码之间的唯一区别就是使用的类型不同，这个时候你就应该考虑是否可以使用类型参数。
		泛型和接口类型之间并不是替代关系，而是相辅相成的关系。泛型的引入是为了配合接口的使用，让我们能够编写更加类型安全的Go代码，并能有效地减少重复代码。
	*/

	fmt.Println("============ 泛型 ============")
}

// 注意：a和b的类型一定要是一样的才行。比如：a是int，b一定是int，不能b为float64。【反推：a是float64，b一定是float64】
func minAny[T int | float64](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// 类型约束字面量，通常外层interface{}可省略
func minAny2[T interface{ int | float64 }](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// Value 事先定义好的类型约束类型
type value interface {
	int | float64
}

// 作为类型约束使用的接口类型可以事先定义并支持复用。
func minAny3[T value](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// IntPtrSlice 使用interface{}包裹
type intPtrSlice[T interface{ *int }] []T

/*
Go语言扩展了接口类型的语法，让我们能够向接口中添加类型。
下面的代码就定义了一个包含 int、 string 和 bool 类型的类型集。
*/
type v interface {
	int | string | bool
}

func foo2[S ~[]E, E any]() {

}

// Scale 返回切片中每个元素都乘c的副本切片
func Scale[E int](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

type Point []int32

func (p Point) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

/*
ScaleAndPrint 不幸的是，这代码会编译失败，输出r.String undefined (type []int32 has no field or method String的错误。
问题是Scale函数返回类型为[]E的值，其中E是参数切片的元素类型。当我们使用Point类型的值调用Scale（其基础类型为[]int32）时，我们返回的是[]int32类型的值，而不是Point类型。
这源于泛型代码的编写方式，但这不是我们想要的。
*/
func ScaleAndPrint(p Point) {
	//r := Scale(p, 2)
	//fmt.Println(r.String()) // 编译失败
}

/*
Scale2 为了解决这个问题，我们必须更改 Scale 函数，以便为切片类型使用类型参数。
我们引入了一个新的类型参数S，它是切片参数的类型。我们对它进行了约束，使得基础类型是S而不是[]E，函数返回的结果类型现在是S。
由于E被约束为整数，因此效果与之前相同：第一个参数必须是某个整数类型的切片。对函数体的唯一更改是，现在我们在调用make时传递S，而不是[]E。
现在这个Scale函数，不仅支持传入普通整数切片参数，也支持传入Point类型参数。
这里需要思考的是，为什么不传递显式类型参数就可以写入 Scale 调用？也就是说，为什么我们可以写 Scale(p, 2)，没有类型参数，而不是必须写 Scale[Point, int32](p, 2) ？
新 Scale 函数有两个类型参数——S 和 E。在不传递任何类型参数的 Scale(p, 2) 调用中，如上所述，函数参数类型推断让编译器推断 S 的类型参数是 Point。
但是这个函数也有一个类型参数 E，它是乘法因子 c 的类型。
相应的函数参数是2，因为2是一个非类型化的常量，函数参数类型推断不能推断出 E 的正确类型(最好的情况是它可以推断出2的默认类型是 int，而这是错误的，因为Point 的基础类型是[]int32)。
相反，编译器推断 E 的类型参数是切片的元素类型的过程称为约束类型推断。
约束类型推断从类型参数约束推导类型参数。当一个类型参数具有根据另一个类型参数定义的约束时使用。
当其中一个类型参数的类型参数已知时，约束用于推断另一个类型参数的类型参数。
通常的情况是，当一个约束对某种类型使用 ~type 形式时，该类型是使用其他类型参数编写的。我们在 Scale 的例子中看到了这一点。
S 是 ~[]E，后面跟着一个用另一个类型参数写的类型[]E。如果我们知道了 S 的类型实参，我们就可以推断出E的类型实参。S 是一个切片类型，而 E是该切片的元素类型。
*/
func Scale2[S ~[]E, E int](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}
