package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")
	led := gpio.NewLedDriver(beagleboneAdaptor, "led", "P9_14")
	btnled := gpio.NewLedDriver(beagleboneAdaptor, "buttonLed", "P9_12")
	button := gpio.NewButtonDriver(beagleboneAdaptor, "button", "P9_16")

	sensor := gpio.NewAnalogSensorDriver(beagleboneAdaptor, "lightSensor", "P9_39")

	sensorAvg := NewAverager("sensorAvg", 100)

	work := func() {
		gobot.On(sensor.Event("data"), func(data interface{}) {
			sensorAvg.Add(data.(int))
			avg := sensorAvg.Compute()
			brightness := uint8(
				gobot.ToScale(gobot.FromScale(float64(avg), 0, 1024), 0, 255),
			)
			fmt.Println("sensor", data, "brightness", brightness, "average", avg)
			//			fmt.Println("brightness", brightness)
			led.Brightness(brightness)
		})

		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})

		gobot.On(button.Event("push"), func(data interface{}) {
			btnled.On()
		})

		gobot.On(button.Event("release"), func(data interface{}) {
			btnled.Off()
		})

	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{led, sensor, btnled, button},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
