package main

import (
	"fmt"
)

type node struct {
	// value            float64
	value                [][]float64
	next                 *node // link to the next Post
	prev                 *node
	operation            string
	parameter_exists     string
	i_derivation         [][]float64
	incoming_parameter   [][]float64
	parameter_derivation [][]float64
	sep                  string
	parameter_label      string
	source_value         [][]float64
	loss                 float64
	// source_value    float64
	// i_derivation     float64
	// parameter_value  float64
}

type Feed struct {
	length int // we'll use it later, or not
	start  *node
	end    *node
}

func (b *Feed) Append(newPost *node) {
	if b.length == 0 {
		b.start = newPost
		b.end = newPost
	} else {
		lastPost := b.end
		lastPost.next = newPost
		newPost.prev = lastPost
		b.end = newPost
	}
	b.length++
}

func transpose(x [][]float64) [][]float64 {
	out := make([][]float64, len(x[0]))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(x[0]); j += 1 {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

// func multiply_float_and_matrix(x float64, y [][]float64) ([][]float64, error) {

// 	out := make([][]float64, len(y))
// 	for i := 0; i < len(y); i++ {
// 		out[i] = make([]float64, len(y[0]))
// 		for k := 0; k < len(y[0]); k++ {

// 			fmt.Print("i: ", i)
// 			// fmt.Print(" j: ", j)
// 			fmt.Print(" k: ", k)
// 			fmt.Print(" Adding: ", x)
// 			fmt.Print(" with ", y[i][k])
// 			out[i][k] = x * y[i][k]
// 			// fmt.Print(" Product: ", x[i][k]*y[k][j])
// 			fmt.Println(" output: ", out[i][k])
// 		}

// 	}

// 	return out, nil
// }

// func add_float_with_matrix(x float64, y [][]float64) ([][]float64, error) {

// 	out := make([][]float64, len(y))
// 	for i := 0; i < len(y); i++ {
// 		out[i] = make([]float64, len(y[0]))
// 		for k := 0; k < len(y[0]); k++ {

// 			fmt.Print("i: ", i)
// 			// fmt.Print(" j: ", j)
// 			fmt.Print(" k: ", k)
// 			fmt.Print(" Adding: ", x)
// 			fmt.Print(" with ", y[i][k])
// 			out[i][k] = x + y[i][k]
// 			// fmt.Print(" Product: ", x[i][k]*y[k][j])
// 			fmt.Println(" output: ", out[i][k])
// 		}

// 	}

// 	return out, nil
// }

func multiply_for_matrices(x, y [][]float64) [][]float64 {
	// fmt.Println("len of x[0]: ", len(x[0]))
	// fmt.Println("len of y: ", len(y))
	// fmt.Println("and len of y[0]: ", len(y[0]))
	// fmt.Printf("Previous matrix being multiplied by a %v x %v matrix\n", len(y), len(y[0]))
	if len(x[0]) != len(y) {
		fmt.Println("Can't do matrix multiplication.")
	}

	out := make([][]float64, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float64, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(y); k++ {
				// fmt.Print("i: ", i)
				// fmt.Print(" j: ", j)
				// fmt.Print(" k: ", k)
				// fmt.Print(" Multiplying: ", x[i][k])
				// fmt.Print(" with ", y[k][j])
				out[i][j] += x[i][k] * y[k][j]
				// fmt.Print(" Product: ", x[i][k]*y[k][j])
				// fmt.Println(" output: ", out[i][j])
			}
		}
	}
	return out
}

func add_for_matrices(x, y [][]float64) [][]float64 {

	out := make([][]float64, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			// fmt.Print("i: ", i)
			// // fmt.Print(" j: ", j)
			// fmt.Print(" k: ", k)
			// fmt.Print(" Adding: ", x[i][k])
			// fmt.Print(" with ", y[i][k])
			out[i][k] = x[i][k] + y[i][k]
			// fmt.Print(" Product: ", x[i][k]*y[k][j])
			// fmt.Println(" output: ", out[i][k])
		}

	}

	return out
}

func subtract_for_matrices(x, y [][]float64) [][]float64 {

	out := make([][]float64, len(x))
	// a := 0.0
	for i := 0; i < len(x); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			// fmt.Print("i: ", i)
			// // fmt.Print(" j: ", j)
			// fmt.Print(" k: ", k)
			// fmt.Print(" Adding: ", x[i][k])
			// fmt.Print(" with ", y[i][k])
			out[i][k] = x[i][k] - y[i][k]
			// a += x[i][k] - y[i][k]
			// fmt.Print(" Product: ", x[i][k]*y[k][j])
			// fmt.Println(" output: ", out[i][k])
		}

	}

	return out
}

func relu_for_matrices(x [][]float64) [][]float64 {

	out := make([][]float64, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float64, len(x[0]))
		for k := 0; k < len(x[0]); k++ {

			// fmt.Print("i: ", i)
			// fmt.Print(" k: ", k)
			// fmt.Println(" Current Element: ", x[i][k])
			if x[i][k] < 0 {
				out[i][k] = 0
			} else {
				out[i][k] = x[i][k]
			}

		}

	}

	return out
}

func square_for_matrices(y [][]float64) [][]float64 {

	out := make([][]float64, len(y))
	for i := 0; i < len(y); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			out[i][k] = y[i][k] * y[i][k]

		}

	}

	return out
}

func loss_calculator(x, y [][]float64) float64 {

	// out := make([][]float64, len(x))
	a := 0.0
	for i := 0; i < len(x); i++ {
		// out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			// fmt.Print("i: ", i)
			// // fmt.Print(" j: ", j)
			// fmt.Print(" k: ", k)
			// fmt.Print(" Adding: ", x[i][k])
			// fmt.Print(" with ", y[i][k])
			// out[i][k] = x[i][k] - y[i][k]
			a += (x[i][k] - y[i][k]) * (x[i][k] - y[i][k])
			// fmt.Print(" Product: ", x[i][k]*y[k][j])
			// fmt.Println(" output: ", out[i][k])
		}

	}

	// a = a * a

	return a
}

func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a float64, b float64) float64 {
	return a - b
}

func product(a float64, b float64) float64 {
	return a * b
}

func square(a float64) float64 {
	return a * a
}

func relu(a float64) float64 {
	if a < 0 {
		return 0
	} else {
		return a
	}
}

func returning_one(row int, column int) [][]float64 {
	one_matrix := make([][]float64, row)
	for i := 0; i < row; i++ {
		one_matrix[i] = make([]float64, column)
		for k := 0; k < column; k++ {

			one_matrix[i][k] = 1.0

		}

	}

	// fmt.Printf("one matrix dimensions: %v x %v\n", len(one_matrix), len(one_matrix[0]))

	return one_matrix
}

func derivative_of_relu(current_node *node) [][]float64 {

	a := current_node.value

	relu_derivative := make([][]float64, len(a))
	for i := 0; i < len(a); i++ {
		relu_derivative[i] = make([]float64, len(a[0]))
		for k := 0; k < len(a[0]); k++ {

			if a[i][k] < 0 {
				relu_derivative[i][k] = 0.0
			} else {
				relu_derivative[i][k] = 1.0
			}

		}

	}

	// fmt.Printf("relu matrix dimensions: %v x %v\n", len(relu_derivative), len(relu_derivative[0]))

	return relu_derivative
}

func addition_derivative(y [][]float64) [][]float64 {

	out := make([][]float64, len(y))
	for i := 0; i < len(y); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			out[i][k] = 1.0

		}

	}

	return out
}

func subtraction_derivative(y [][]float64) [][]float64 {

	out := make([][]float64, len(y))
	for i := 0; i < len(y); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			out[i][k] = -1

		}

	}

	return out
}

func derivative_of_subtraction(current_node *node) [][]float64 {

	a := current_node.value

	subtraction_der := subtraction_derivative(a)

	// fmt.Printf("subtraction matrix dimensions: %v x %v\n", len(subtraction_der), len(subtraction_der[0]))

	return subtraction_der
}

func derivative_of_addition(current_node *node) [][]float64 {

	a := current_node.value

	addition_der := addition_derivative(a)

	// fmt.Printf("subtraction matrix dimensions: %v x %v\n", len(addition_der), len(addition_der[0]))

	return addition_der
}

func loss_derivative_calculator(current_node *node) [][]float64 {

	y := current_node.value

	out := make([][]float64, len(y))
	// a := 0.0
	for i := 0; i < len(y); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {
			out[i][k] = 2.0 * y[i][k]
		}

	}

	return out
}

func derivative_conditions(current_node *node) [][]float64 {
	// fmt.Printf("i7 %v\n", current_node.prev.operation)

	if current_node.operation == "product" {

		current_node.i_derivation = returning_one(len(current_node.value), len(current_node.value[0]))

	} else if current_node.operation == "relu" {
		current_node.i_derivation = derivative_of_relu(current_node)
	} else if current_node.operation == "subtract" {
		current_node.i_derivation = derivative_of_subtraction(current_node)
	} else if current_node.sep == "last" {
		current_node.i_derivation = loss_derivative_calculator(current_node)

	} else if current_node.operation == "add" {
		current_node.i_derivation = returning_one(len(current_node.value), len(current_node.value[0]))
	}
	// else if current_node.operation == "add" {
	// 	current_node.i_derivation = 1.0

	// } else if current_node.operation == "product" {
	// 	current_node.i_derivation = current_node.source_value

	// } else if current_node.operation == "subtract" {
	// 	current_node.i_derivation = -1.0

	// }

	return current_node.prev.i_derivation

}

// func derivative_conditions(current_node *node) float64 {
// 	// fmt.Printf("i7 %v\n", current_node.prev.operation)

// 	if current_node.operation == "square" {
// 		current_node.i_derivation = 2.0 * current_node.prev.value

// 	} else if current_node.operation == "add" {
// 		current_node.i_derivation = 1.0

// 	} else if current_node.operation == "product" {
// 		current_node.i_derivation = 1.0

// 	} else if current_node.operation == "subtract" {
// 		current_node.i_derivation = -1.0

// 	}

// 	return current_node.prev.i_derivation

// }

func calculate_parameter(current_node *node) {
	if current_node.operation == "add" {

		current_node.parameter_derivation = returning_one(len(current_node.incoming_parameter), len(current_node.incoming_parameter[0]))

	} else if current_node.operation == "product" {

		current_node.parameter_derivation = current_node.source_value
	}
}

// func calculate_parameter(current_node *node) {
// 	if current_node.operation == "add" {
// 		current_node.parameter_value = 1.0 * current_node.next.i_derivation
// 	} else if current_node.operation == "product" {
// 		current_node.parameter_value = current_node.source_value
// 	}
// }

func main() {
	f := &Feed{}

	x := [][]float64{
		[]float64{1.0, -2.0, 3.0},
	}

	a := [][]float64{
		[]float64{0.2, 0.2, 0.2, 0.2, 0.2},
		[]float64{0.2, 0.2, 0.2, 0.2, 0.2},
		[]float64{0.2, 0.2, 0.2, 0.2, 0.2},
	}

	// m1 := [3][3]int{}
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		m1[i][j] = rand.Intn(9)
	// 	}
	// }

	// fmt.Println("a: ", m1)

	b := [][]float64{
		[]float64{0.5},
		[]float64{0.5},
		[]float64{0.5},
		[]float64{0.5},
		[]float64{0.5},
	}

	// c := [][]float64{
	// 	[]float64{1.0, 1.0},
	// 	[]float64{1.0, 1.0},
	// 	[]float64{1.0, 1.0},
	// 	[]float64{1.0, 1.0},
	// 	[]float64{1.0, 1.0},
	// }

	y := [][]float64{
		[]float64{5.0, 5.0},
		// []float64{5.0, 5.0},
		// []float64{5.0, 5.0},
		// []float64{5.0, 5.0},
		// []float64{5.0, 5.0},
	}

	// temp := [][]float64{
	// 	[]float64{0.0, 0.0, 0.0},
	// 	[]float64{0.0, 0.0, 0.0},
	// }

	// x := 2.0
	// a := 0.2
	// b := 0.5
	// c := 1.0
	// y := 5.0

	i1 := node{
		value:            x,
		operation:        "init",
		parameter_exists: "no",
		sep:              "no",
	}
	f.Append(&i1)

	current_node := f.start

	// fmt.Print("Current output dimensions: ", len(i1.value))
	// fmt.Println(" x", len(i1.value[0]))

	// i2 := node{
	// 	value:            relu_for_matrices(i1.value),
	// 	operation:        "relu",
	// 	parameter_exists: "no",
	// 	sep:              "no",
	// }
	// f.Append(&i2)

	// current_node = current_node.next

	// The row of the first trainable parameter/ weight must match the the column of the value that
	// comes out of the previous node

	i2 := node{
		value:              multiply_for_matrices(i1.value, a),
		operation:          "product",
		parameter_exists:   "yes",
		parameter_label:    "a",
		sep:                "no",
		source_value:       i1.value,
		incoming_parameter: a,
	}
	f.Append(&i2)

	current_node = current_node.next
	derivative_conditions(current_node)

	i3 := node{
		value:            relu_for_matrices(i2.value),
		operation:        "relu",
		parameter_exists: "no",
		sep:              "no",
	}
	f.Append(&i3)

	current_node = current_node.next

	derivative_conditions(current_node)

	i4 := node{
		value:              add_for_matrices(i3.value, b),
		operation:          "add",
		parameter_exists:   "yes",
		parameter_label:    "b",
		sep:                "no",
		source_value:       i3.value,
		incoming_parameter: b,
	}
	f.Append(&i4)

	current_node = current_node.next
	derivative_conditions(current_node)

	i5 := node{
		value:            i4.value,
		parameter_exists: "no",
		sep:              "last",
	}
	f.Append(&i5)

	current_node = current_node.next

	loss_value := loss_calculator(y, current_node.value)

	current_node.loss = loss_value

	fmt.Println("Loss: ", current_node.loss)

	derivative_conditions(current_node)

	// i6 := node{
	// 	value:            square_for_matrices(i5.value),
	// 	operation:        "square",
	// 	parameter_exists: "no",
	// 	sep:              "last",
	// }
	// f.Append(&i6)

	// current_node = current_node.next

	// derivative_conditions(current_node)

	// i3_1 := node{
	// 	value:            multiply_for_matrices(i1.value, transpose(b)),
	// 	operation:        "product",
	// 	parameter_exists: "yes",
	// 	parameter_label:  "b",
	// 	sep:              "yes",
	// 	source_value:     i1.value,
	// }
	// f.Append(&i3_1)

	// current_node = current_node.next
	// derivative_conditions(current_node)

	// i4 := node{
	// 	value:            add_for_matrices(i3_0.value, i3_1.value),
	// 	operation:        "add",
	// 	parameter_exists: "no",
	// 	sep:              "no",
	// }
	// f.Append(&i4)

	// current_node = current_node.next
	// derivative_conditions(current_node)

	// i5 := node{
	// 	value:            add_for_matrices(i4.value, c),
	// 	operation:        "add",
	// 	parameter_exists: "yes",
	// 	parameter_label:  "c",
	// 	sep:              "no",
	// 	source_value:     i4.value,
	// }
	// f.Append(&i5)

	// current_node = current_node.next
	// derivative_conditions(current_node)

	// i6 := node{
	// 	value:            subtract_for_matrices(y, i5.value),
	// 	operation:        "subtract",
	// 	parameter_exists: "no",
	// 	sep:              "no",
	// }
	// f.Append(&i6)

	// current_node = current_node.next
	// derivative_conditions(current_node)

	// // this is the linked list node that accepts the float64 value

	// i7 := node{
	// 	value:            square_for_matrices(i6.value),
	// 	operation:        "square",
	// 	parameter_exists: "no",
	// 	sep:              "last",
	// }
	// f.Append(&i7)

	// current_node = current_node.next

	// derivative_conditions(current_node)

	node_for_back := f.end

	for i := 0; i < f.length-1; i++ {
		if node_for_back.sep == "no" {

			column_of_current_derivation := len(node_for_back.i_derivation[0])
			row_of_current_derivation := len(node_for_back.i_derivation)
			column_of_next_derivation := len(node_for_back.next.i_derivation[0])
			row_of_next_derivation := len(node_for_back.next.i_derivation)
			if column_of_current_derivation == row_of_next_derivation {
				fmt.Println("THIS COULD WORK")
				fmt.Printf("Current dimensions: %v x %v\n", row_of_current_derivation, column_of_current_derivation)
				fmt.Printf("Next dimensions: %v x %v\n", row_of_next_derivation, column_of_next_derivation)
				node_for_back.i_derivation = multiply_for_matrices(node_for_back.i_derivation, node_for_back.next.i_derivation)
			} else if row_of_current_derivation == row_of_next_derivation {
				fmt.Println("THIS COULD WORK")
				transpose_of_current_derivation := transpose(node_for_back.i_derivation)
				node_for_back.i_derivation = multiply_for_matrices(transpose_of_current_derivation, node_for_back.next.i_derivation)

			} else if column_of_current_derivation == column_of_next_derivation {
				fmt.Println("THIS COULD WORK")
				transpose_of_next_derivation := transpose(node_for_back.next.i_derivation)
				node_for_back.i_derivation = multiply_for_matrices(node_for_back.i_derivation, transpose_of_next_derivation)

			} else if row_of_current_derivation == column_of_next_derivation {
				fmt.Println("THIS COULD WORK")
				transpose_of_current_derivation := transpose(node_for_back.i_derivation)
				node_for_back.i_derivation = multiply_for_matrices(transpose_of_current_derivation, node_for_back.next.i_derivation)

			} else {
				fmt.Println("THIS WON'T HAPPEN")
				fmt.Printf("Current dimensions: %v x %v\n", row_of_current_derivation, column_of_current_derivation)
				fmt.Printf("Next dimensions: %v x %v\n", row_of_next_derivation, column_of_next_derivation)
			}
			// node_for_back.i_derivation = (node_for_back.i_derivation) * (node_for_back.next.i_derivation)
		} else if node_for_back.sep == "yes" {

			// node_for_back.i_derivation = node_for_back.i_derivation * node_for_back.next.i_derivation
		}
		node_for_back = node_for_back.prev
	}

	tracker_for_parameter_update := f.end
	for i := 0; i < f.length-1; i++ {
		if tracker_for_parameter_update.parameter_exists == "yes" {
			calculate_parameter(tracker_for_parameter_update)
			fmt.Printf("Intermediate derivate for %v parameter: %v\n", tracker_for_parameter_update.parameter_label, tracker_for_parameter_update.parameter_derivation)
		}
		tracker_for_parameter_update = tracker_for_parameter_update.prev
	}

	fmt.Printf("i1: %v\n", f.start.value)
	fmt.Printf("Current output dimensions: %v x %v \n", len(f.start.value), len(f.start.value[0]))

	for i := 0; i < f.length-1; i++ {

		f.start = f.start.next
		fmt.Printf("i%v: %v\n", i+2, f.start.value)
		fmt.Printf("Current output dimensions: %v x %v \n", len(f.start.value), len(f.start.value[0]))
		fmt.Printf("i%v operation: %v\n", i+2, f.start.operation)
		fmt.Printf("i%v current derived %v \n", i+2, f.start.i_derivation)
		fmt.Printf("i%v current parameter derivation %v \n", i+2, f.start.parameter_derivation)

	}

	fmt.Printf("B %v \n", a)

}
