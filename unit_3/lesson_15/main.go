package main

import "fmt"

type celsius float64
type fahrenheint float64

func (c celsius) fahrenheint() fahrenheint {
	return fahrenheint((9.0 / 5.0 * c) + 32.0)
}

func (f fahrenheint) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func draw_line() {
	fmt.Println("=============")
}

func draw_row(val1, val2 float64) {
	fmt.Printf("|%5v|%5v|\n", val1, val2)
}

func c_to_f_datagen(row_num int) (string, string) {
	c := celsius(row_num*5.0 - 40.0)
	f := c.fahrenheint()
	return fmt.Sprintf("%.1f", c), fmt.Sprintf("%.1f", f)
}
func f_to_c_datagen(row_num int) (string, string) {
	f := fahrenheint(row_num*5.0 - 40.0)
	c := f.celsius()
	return fmt.Sprintf("%.1f", f), fmt.Sprintf("%.1f", c)
}

func draw_table(hdr1, hdr2 string, num_rows int, getRow func(int) (string, string)) {
	draw_line()
	fmt.Printf("|%5s|%5s|\n", hdr1, hdr2)
	draw_line()
	for i := 0; i < num_rows; i++ {
		col1, col2 := getRow(i)
		fmt.Printf("|%5s|%5s|\n", col1, col2)
	}

}

func main() {
	draw_table("C", "F", 29, c_to_f_datagen)
	draw_table("F", "C", 29, f_to_c_datagen)
}
