package main

import (
	"fmt"
)

type node struct {
	value                [][]float64
	next                 *node
	prev                 *node       // i don't believe i'm using this
	operation            string      // what operation is going to happen in this node, it'll just be a string , so we can use this to determine what sort of derivative operation we'll need
	parameter_exists     string      // if this is yes, meaning there exists a new trainable parameter in this step, then the code will know to come upto this node for the backpropagation
	i_derivation         [][]float64 // this calculates the i derivative for the current operation, required for the chain rule used for back propagation
	incoming_parameter   [][]float64
	parameter_derivation [][]float64 // when you backpropagate to a node that was the starting point of a trainable parameter
	sep                  string      // if the current node has no connection with the previous node, say in ax+bx these nodes will be from seperate sources
	parameter_label      string      //this is just for printing purpose
	source_value         [][]float64 // when a node comes from a seperate node, it is essential to track the source for backpropagation purpose
	loss                 float64     // loss value must be a scalar
}

type Feed struct {
	length int // we'll use it later, or not
	start  *node
	end    *node
}

// Code for appending a node

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

// Transposing a matrix

func transpose(x [][]float64) [][]float64 {
	out := make([][]float64, len(x[0]))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(x[0]); j += 1 {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

// Multiplication of two matrices

func multiply_for_matrices(x, y [][]float64) [][]float64 {

	if len(x[0]) != len(y) {
		fmt.Println("Can't do matrix multiplication.")
	}

	out := make([][]float64, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float64, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(y); k++ {

				out[i][j] += x[i][k] * y[k][j]

			}
		}
	}
	return out
}

// Addition of two matrices

func add_for_matrices(x, y [][]float64) [][]float64 {

	out := make([][]float64, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			out[i][k] = x[i][k] + y[i][k]

		}

	}

	return out
}

// Subtraction of two matrices

func subtract_for_matrices(x, y [][]float64) [][]float64 {

	out := make([][]float64, len(x))
	// a := 0.0
	for i := 0; i < len(x); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			out[i][k] = x[i][k] - y[i][k]

		}

	}

	return out
}

// Relu during forward pass

func relu_for_matrices(x [][]float64) [][]float64 {

	out := make([][]float64, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float64, len(x[0]))
		for k := 0; k < len(x[0]); k++ {

			if x[i][k] < 0 {
				out[i][k] = 0
			} else {
				out[i][k] = x[i][k]
			}

		}

	}

	return out
}

// Square

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

			a += (x[i][k] - y[i][k]) * (x[i][k] - y[i][k])

		}

	}

	return a
}

// not sure if the following functions are required anymore

// ------------------------------------------------------------------- //

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

// ------------------------------------------------------------------- //

// given the dimensions, the following function can return a matrix full of ones

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

// Relu for backpropagation

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

// func addition_derivative(y [][]float64) [][]float64 {

// 	out := make([][]float64, len(y))
// 	for i := 0; i < len(y); i++ {
// 		out[i] = make([]float64, len(y[0]))
// 		for k := 0; k < len(y[0]); k++ {

// 			out[i][k] = 1.0

// 		}

// 	}

// 	return out
// }

// func subtraction_derivative(y [][]float64) [][]float64 {

// 	out := make([][]float64, len(y))
// 	for i := 0; i < len(y); i++ {
// 		out[i] = make([]float64, len(y[0]))
// 		for k := 0; k < len(y[0]); k++ {

// 			out[i][k] = -1

// 		}

// 	}

// 	return out
// }

func derivative_of_subtraction(current_node *node) [][]float64 {

	y := current_node.value

	out := make([][]float64, len(y))
	for i := 0; i < len(y); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			out[i][k] = -1.0

		}

	}

	return out
}

func derivative_of_addition(current_node *node) [][]float64 {

	y := current_node.value

	out := make([][]float64, len(y))
	for i := 0; i < len(y); i++ {
		out[i] = make([]float64, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			out[i][k] = 1.0

		}

	}

	return out

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

	return current_node.prev.i_derivation

}

func calculate_parameter(current_node *node) {
	if current_node.operation == "add" {

		parameter_derivative := returning_one(len(current_node.incoming_parameter), len(current_node.incoming_parameter[0]))
		current_node.parameter_derivation = multiply_for_matrices(parameter_derivative, current_node.next.i_derivation)

	} else if current_node.operation == "product" {

		column_of_current_derivation := len(current_node.source_value[0])
		row_of_current_derivation := len(current_node.source_value)
		column_of_next_derivation := len(current_node.next.i_derivation[0])
		row_of_next_derivation := len(current_node.next.i_derivation)

		if column_of_current_derivation == row_of_next_derivation {

			fmt.Printf("Current dimensions: %v x %v\n", row_of_current_derivation, column_of_current_derivation)
			fmt.Printf("Next dimensions: %v x %v\n", row_of_next_derivation, column_of_next_derivation)
			current_node.parameter_derivation = multiply_for_matrices(current_node.i_derivation, current_node.next.i_derivation)
		} else if row_of_current_derivation == row_of_next_derivation {

			transpose_of_current_derivation := transpose(current_node.i_derivation)
			current_node.parameter_derivation = multiply_for_matrices(transpose_of_current_derivation, current_node.next.i_derivation)

		} else if column_of_current_derivation == column_of_next_derivation {

			transpose_of_next_derivation := transpose(current_node.next.i_derivation)
			current_node.parameter_derivation = multiply_for_matrices(current_node.source_value, transpose_of_next_derivation)

		} else if row_of_current_derivation == column_of_next_derivation {

			transpose_of_current_derivation := transpose(current_node.next.i_derivation)
			current_node.parameter_derivation = multiply_for_matrices(transpose_of_current_derivation, current_node.next.i_derivation)

		} else {
			fmt.Println("THIS WON'T HAPPEN")
			fmt.Printf("Current dimensions: %v x %v\n", row_of_current_derivation, column_of_current_derivation)
			fmt.Printf("Next dimensions: %v x %v\n", row_of_next_derivation, column_of_next_derivation)
		}

	}
}

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

	node_for_back := f.end

	for i := 0; i < f.length-1; i++ {
		if node_for_back.sep == "no" {

			column_of_current_derivation := len(node_for_back.i_derivation[0])
			row_of_current_derivation := len(node_for_back.i_derivation)
			column_of_next_derivation := len(node_for_back.next.i_derivation[0])
			row_of_next_derivation := len(node_for_back.next.i_derivation)
			if column_of_current_derivation == row_of_next_derivation {

				// fmt.Printf("Current dimensions: %v x %v\n", row_of_current_derivation, column_of_current_derivation)
				// fmt.Printf("Next dimensions: %v x %v\n", row_of_next_derivation, column_of_next_derivation)
				node_for_back.i_derivation = multiply_for_matrices(node_for_back.i_derivation, node_for_back.next.i_derivation)

			} else if row_of_current_derivation == row_of_next_derivation {

				transpose_of_current_derivation := transpose(node_for_back.i_derivation)
				node_for_back.i_derivation = multiply_for_matrices(transpose_of_current_derivation, node_for_back.next.i_derivation)

			} else if column_of_current_derivation == column_of_next_derivation {

				transpose_of_next_derivation := transpose(node_for_back.next.i_derivation)
				node_for_back.i_derivation = multiply_for_matrices(node_for_back.i_derivation, transpose_of_next_derivation)

			} else if row_of_current_derivation == column_of_next_derivation {

				transpose_of_current_derivation := transpose(node_for_back.i_derivation)
				node_for_back.i_derivation = multiply_for_matrices(transpose_of_current_derivation, node_for_back.next.i_derivation)

			} else {
				fmt.Println("THIS WON'T HAPPEN")
				fmt.Printf("Current dimensions: %v x %v\n", row_of_current_derivation, column_of_current_derivation)
				fmt.Printf("Next dimensions: %v x %v\n", row_of_next_derivation, column_of_next_derivation)
			}

		} else if node_for_back.sep == "yes" {

		}
		node_for_back = node_for_back.prev
	}

	tracker_for_parameter_update := f.end
	for i := 0; i < f.length-1; i++ {
		if tracker_for_parameter_update.parameter_exists == "yes" {
			calculate_parameter(tracker_for_parameter_update)
			fmt.Printf("Intermediate derivate for %v parameter: %v\n", tracker_for_parameter_update.parameter_label, tracker_for_parameter_update.parameter_derivation)
			fmt.Printf("Dimensions of %v: %v x %v\n", tracker_for_parameter_update.parameter_label, len(tracker_for_parameter_update.parameter_derivation), len(tracker_for_parameter_update.parameter_derivation[0]))
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

	// fmt.Printf("A %v \n", a)
	fmt.Printf("Dimensions of a: %v x %v\n", len(a), len(a[0]))

	// For Parameter Update

}
