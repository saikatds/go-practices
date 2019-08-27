package main

import "fmt"

type car struct {
	gas_pedal uint16
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}

func main(){
	a_car := car{12,33 , 34 , 34.5}

	fmt.Println(a_car.brake_pedal)
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.steering_wheel)
	fmt.Println(a_car.top_speed_kmh)

}