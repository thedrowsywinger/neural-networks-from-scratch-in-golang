package main

import (
	"errors"
	"fmt"
)

type node struct {
	value            float64
	next             *node // link to the next Post
	prev             *node
	operation        string
	parameter_exists string
	i_derivation     float64
	parameter_value  float64
	sep              string
	parameter_label  string
	source_value     float64
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

func transpose(x [][]float32) [][]float32 {
	out := make([][]float32, len(x[0]))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(x[0]); j += 1 {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func multiply_for_matrices(x, y [][]float32) ([][]float32, error) {
	if len(x[0]) != len(y) {
		return nil, errors.New("Can't do matrix multiplication.")
	}

	out := make([][]float32, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(y); k++ {
				fmt.Print("i: ", i)
				fmt.Print(" j: ", j)
				fmt.Print(" k: ", k)
				fmt.Print(" Multiplying: ", x[i][k])
				fmt.Print(" with ", y[k][j])
				out[i][j] += x[i][k] * y[k][j]
				fmt.Print(" Product: ", x[i][k]*y[k][j])
				fmt.Println(" output: ", out[i][j])
			}
		}
	}
	return out, nil
}

func add_for_matrices(x, y [][]float32) ([][]float32, error) {

	out := make([][]float32, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(y[0]))
		for k := 0; k < len(y[0]); k++ {

			fmt.Print("i: ", i)
			// fmt.Print(" j: ", j)
			fmt.Print(" k: ", k)
			fmt.Print(" Adding: ", x[i][k])
			fmt.Print(" with ", y[i][k])
			out[i][k] = x[i][k] + y[i][k]
			// fmt.Print(" Product: ", x[i][k]*y[k][j])
			fmt.Println(" output: ", out[i][k])
		}

	}

	return out, nil
}

func relu_for_matrices(x [][]float32) [][]float32 {

	out := make([][]float32, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(x[0]))
		for k := 0; k < len(x[0]); k++ {

			fmt.Print("i: ", i)
			fmt.Print(" k: ", k)
			fmt.Println(" Current Element: ", x[i][k])
			if x[i][k] < 0 {
				out[i][k] = 0
			} else {
				out[i][k] = x[i][k]
			}

		}

	}

	return out
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

func derivative_conditions(current_node *node) float64 {
	// fmt.Printf("i7 %v\n", current_node.prev.operation)

	if current_node.operation == "square" {
		current_node.i_derivation = 2.0 * current_node.prev.value

	} else if current_node.operation == "add" {
		current_node.i_derivation = 1.0

	} else if current_node.operation == "product" {
		current_node.i_derivation = 1.0

	} else if current_node.operation == "subtract" {
		current_node.i_derivation = -1.0

	}

	return current_node.prev.i_derivation

}

func calculate_parameter(current_node *node) {
	if current_node.operation == "add" {
		current_node.parameter_value = 1.0 * current_node.next.i_derivation
	} else if current_node.operation == "product" {
		current_node.parameter_value = current_node.source_value
	}
}

func main() {
	f := &Feed{}

	x := 2.0
	a := 0.2
	b := 0.5
	c := 1.0
	y := 5.0

	i1 := node{
		value:            x,
		operation:        "init",
		parameter_exists: "no",
		i_derivation:     0.0,
		parameter_value:  0.0,
		sep:              "no",
	}
	f.Append(&i1)

	current_node := f.start

	i2 := node{
		value:            square(i1.value),
		operation:        "square",
		parameter_exists: "no",
		sep:              "no",
	}
	f.Append(&i2)

	current_node = current_node.next

	i3_0 := node{
		value:            product(i2.value, a),
		operation:        "product",
		parameter_exists: "yes",
		parameter_label:  "a",
		sep:              "no",
		source_value:     i2.value,
	}
	f.Append(&i3_0)

	current_node = current_node.next
	derivative_conditions(current_node)

	i3_1 := node{
		value:            product(i1.value, b),
		operation:        "product",
		parameter_exists: "yes",
		parameter_label:  "b",
		sep:              "yes",
		source_value:     i1.value,
	}
	f.Append(&i3_1)

	current_node = current_node.next
	derivative_conditions(current_node)

	i4 := node{
		value:            add(i3_0.value, i3_1.value),
		operation:        "add",
		parameter_exists: "no",
		sep:              "no",
	}
	f.Append(&i4)

	current_node = current_node.next
	derivative_conditions(current_node)

	i5 := node{
		value:            add(i4.value, c),
		operation:        "add",
		parameter_exists: "yes",
		parameter_label:  "c",
		sep:              "no",
		source_value:     i4.value,
	}
	f.Append(&i5)

	current_node = current_node.next
	derivative_conditions(current_node)

	i6 := node{
		value:            subtract(y, i5.value),
		operation:        "subtract",
		parameter_exists: "no",
		sep:              "no",
	}
	f.Append(&i6)

	current_node = current_node.next
	derivative_conditions(current_node)

	// this is the linked list node that accepts the float64 value

	i7 := node{
		value:            square(i6.value),
		operation:        "square",
		parameter_exists: "no",
		sep:              "last",
	}
	f.Append(&i7)

	current_node = current_node.next

	derivative_conditions(current_node)

	node_for_back := f.end

	for i := 0; i < f.length-1; i++ {
		if node_for_back.sep == "no" {
			node_for_back.i_derivation = node_for_back.i_derivation * node_for_back.next.i_derivation
		} else if node_for_back.sep == "yes" {

			node_for_back.i_derivation = node_for_back.i_derivation * node_for_back.next.i_derivation
		}
		node_for_back = node_for_back.prev
	}

	tracker_for_parameter_update := f.end
	for i := 0; i < f.length-1; i++ {
		if tracker_for_parameter_update.parameter_exists == "yes" {
			calculate_parameter(tracker_for_parameter_update)
			fmt.Printf("Intermediate derivate for %v parameter: %v\n", tracker_for_parameter_update.parameter_label, tracker_for_parameter_update.parameter_value)
		}
		tracker_for_parameter_update = tracker_for_parameter_update.prev
	}

	for i := 0; i < f.length-1; i++ {

		f.start = f.start.next
		fmt.Printf("i%v: %v\n", i+2, f.start.value)
		fmt.Printf("i%v operation: %v\n", i+2, f.start.operation)
		fmt.Printf("i%v current derived %v \n", i+2, f.start.i_derivation)

	}

}
