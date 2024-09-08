// Copyright 2022 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Orange Pi pin out.

package orangepi

import (
	"errors"
	"fmt"
	"strings"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/pin"
	"periph.io/x/conn/v3/pin/pinreg"
	"periph.io/x/host/v3/allwinner"
	"periph.io/x/host/v3/distro"
)

// Present return true if a Orange Pi board is detected.
func Present() bool {
	if isArm {
		// This works for the Orange Pi Zero, not sure if other Orange Pi boards
		// match the same DTModel prefix.
		return strings.HasPrefix(distro.DTModel(), "OrangePi")
	}

	return false
}

const (
	boardZero   string = "Orange Pi Zero"    // + LTS (H2/H3 have identical pinouts)
	boardZero2W string = "Orange Pi Zero 2W" //
)

var (
	PA1_1  pin.Pin    = pin.DC_IN // VCC 3v3 Ext
	PA1_2  pin.Pin    = pin.V5
	PA1_3  gpio.PinIO = allwinner.PA12
	PA1_4  pin.Pin    = pin.V5
	PA1_5  gpio.PinIO = allwinner.PA11
	PA1_6  pin.Pin    = pin.GROUND
	PA1_7  gpio.PinIO = allwinner.PA6
	PA1_8  gpio.PinIO = allwinner.PG6
	PA1_9  pin.Pin    = pin.GROUND
	PA1_10 gpio.PinIO = allwinner.PG7
	PA1_11 gpio.PinIO = allwinner.PA1
	PA1_12 gpio.PinIO = allwinner.PA7
	PA1_13 gpio.PinIO = allwinner.PA0
	PA1_14 pin.Pin    = pin.GROUND
	PA1_15 gpio.PinIO = allwinner.PA3
	PA1_16 gpio.PinIO = allwinner.PA19
	PA1_17 pin.Pin    = pin.DC_IN // VCC 3v3 Ext
	PA1_18 gpio.PinIO = allwinner.PA18
	PA1_19 gpio.PinIO = allwinner.PA15
	PA1_20 pin.Pin    = pin.GROUND
	PA1_21 gpio.PinIO = allwinner.PA16
	PA1_22 gpio.PinIO = allwinner.PA2
	PA1_23 gpio.PinIO = allwinner.PA14
	PA1_24 gpio.PinIO = allwinner.PA13
	PA1_25 pin.Pin    = pin.GROUND
	PA1_26 gpio.PinIO = allwinner.PA10

	FUN1_1  pin.Pin = pin.V5
	FUN1_2  pin.Pin = pin.GROUND
	FUN1_3  pin.Pin = gpio.INVALID       // USB-DM2
	FUN1_4  pin.Pin = gpio.INVALID       // USB-DP2
	FUN1_5  pin.Pin = gpio.INVALID       // USB-DM3
	FUN1_6  pin.Pin = gpio.INVALID       // USB-DP3
	FUN1_7  pin.Pin = allwinner.HP_RIGHT // LINEOUTR
	FUN1_8  pin.Pin = allwinner.HP_LEFT  // LINEOUTL
	FUN1_9  pin.Pin = gpio.INVALID       // TVOUT
	FUN1_10 pin.Pin = gpio.INVALID       // MBIAS, Bias Voltage output for mic
	FUN1_11 pin.Pin = allwinner.MIC_IN   // INPUT Analog Microphone pin (+)
	FUN1_12 pin.Pin = allwinner.MIC_GND  // INPUT Analog Microphone pin (-)
	FUN1_13 pin.Pin = allwinner.PL11     // IR-RX

	// Orange Pi Zero 2W
	P1_1  pin.Pin    = pin.V3_3       // max 30mA
	P1_2  pin.Pin    = pin.V5         // (filtered)
	P1_3  gpio.PinIO = allwinner.PI8  // High, TWI1_SDA
	P1_4  pin.Pin    = pin.V5         //
	P1_5  gpio.PinIO = allwinner.PI7  // High, TWI1_SCL
	P1_6  pin.Pin    = pin.GROUND     //
	P1_7  gpio.PinIO = allwinner.PI13 // High, PWM3
	P1_8  gpio.PinIO = allwinner.PH0  // Low, UART0_TX
	P1_9  pin.Pin    = pin.GROUND     //
	P1_10 gpio.PinIO = allwinner.PH1  // Low, UART0_RX
	P1_11 gpio.PinIO = allwinner.PH2  // Low, UART5_TX
	P1_12 gpio.PinIO = allwinner.PI1  // Low,
	P1_13 gpio.PinIO = allwinner.PH3  // Low, UART5_RX
	P1_14 pin.Pin    = pin.GROUND     //
	P1_15 gpio.PinIO = allwinner.PI5  // Low, UART2_TX
	P1_16 gpio.PinIO = allwinner.PI4  // Low, PWM4
	P1_17 pin.Pin    = pin.V3_3       //
	P1_18 gpio.PinIO = allwinner.PH4  // Low,
	P1_19 gpio.PinIO = allwinner.PH7  // Low, SPI1_MOSI
	P1_20 pin.Pin    = pin.GROUND     //
	P1_21 gpio.PinIO = allwinner.PH8  // Low, SPI1_MISO
	P1_22 gpio.PinIO = allwinner.PI6  // Low,
	P1_23 gpio.PinIO = allwinner.PH6  // Low, SPI1_CLK
	P1_24 gpio.PinIO = allwinner.PH5  // High, SPI1_CS0
	P1_25 pin.Pin    = pin.GROUND     //
	P1_26 gpio.PinIO = allwinner.PH9  // High, SPI1_CS1
	P1_27 gpio.PinIO = allwinner.PI10 // High, TWI2_SDA used to probe for HAT EEPROM, see https://github.com/raspberrypi/hats
	P1_28 gpio.PinIO = allwinner.PI9  // High, TWI2_SCL
	P1_29 gpio.PinIO = allwinner.PI0  //
	P1_30 pin.Pin    = pin.GROUND     //
	P1_31 gpio.PinIO = allwinner.PI15 //
	P1_32 gpio.PinIO = allwinner.PI11 // Low,  PWM1
	P1_33 gpio.PinIO = allwinner.PI12 // Low,  PWM2
	P1_34 pin.Pin    = pin.GROUND     //
	P1_35 gpio.PinIO = allwinner.PI2  // Low,  I2S_WS, SPI1_MISO, PWM1
	P1_36 gpio.PinIO = allwinner.PC12 // Low,  UART0_CTS, SPI1_CS2, UART1_CTS
	P1_37 gpio.PinIO = allwinner.PI16 //
	P1_38 gpio.PinIO = allwinner.PI4  // Low,  I2S_DIN, SPI1_MOSI, CLK0
	P1_39 pin.Pin    = pin.GROUND     //
	P1_40 gpio.PinIO = allwinner.PI3  //
)

// registerHeaders registers the headers for various Orange Pi boards. Currently
// only Orange Pi Zero is supported.
func registerHeaders(model string) error {
	// http://www.orangepi.org/html/hardWare/computerAndMicrocontrollers/details/Orange-Pi-Zero.html
	if strings.Contains(model, boardZero) {
		// 26pin expansion port
		if err := pinreg.Register("PA", [][]pin.Pin{
			{PA1_1, PA1_2},
			{PA1_3, PA1_4},
			{PA1_5, PA1_6},
			{PA1_7, PA1_8},
			{PA1_9, PA1_10},
			{PA1_11, PA1_12},
			{PA1_13, PA1_14},
			{PA1_15, PA1_16},
			{PA1_17, PA1_18},
			{PA1_19, PA1_20},
			{PA1_21, PA1_22},
			{PA1_23, PA1_24},
			{PA1_25, PA1_26},
		}); err != nil {
			return err
		}

		// 13pin function interface
		if err := pinreg.Register("FUN", [][]pin.Pin{
			{FUN1_1},
			{FUN1_2},
			{FUN1_3},
			{FUN1_4},
			{FUN1_5},
			{FUN1_6},
			{FUN1_7},
			{FUN1_8},
			{FUN1_9},
			{FUN1_10},
			{FUN1_11},
			{FUN1_12},
			{FUN1_13},
		}); err != nil {
			return err
		}
	} else if strings.Contains(model, boardZero2W) {
		if err := pinreg.Register("P1", [][]pin.Pin{
			{P1_1, P1_2},
			{P1_3, P1_4},
			{P1_5, P1_6},
			{P1_7, P1_8},
			{P1_9, P1_10},
			{P1_11, P1_12},
			{P1_13, P1_14},
			{P1_15, P1_16},
			{P1_17, P1_18},
			{P1_19, P1_20},
			{P1_21, P1_22},
			{P1_23, P1_24},
			{P1_25, P1_26},
			{P1_27, P1_28},
			{P1_29, P1_30},
			{P1_31, P1_32},
			{P1_33, P1_34},
			{P1_35, P1_36},
			{P1_37, P1_38},
			{P1_39, P1_40},
		}); err != nil {
			return err
		}

	}
	return nil
}

// driver implements periph.Driver.
type driver struct {
}

// String is the text representation of the board.
func (d *driver) String() string {
	return "orangepi"
}

// Prerequisites load drivers before the actual driver is loaded. For
// these boards, we do not need any prerequisites.
func (d *driver) Prerequisites() []string {
	return nil
}

// After this driver is loaded, we need to load generic Allwinner drivers
// for the GPIO pins which are identical on all Allwinner CPUs.
func (d *driver) After() []string {
	return []string{"allwinner-gpio", "allwinner-gpio-pl"}
}

// Init initializes the driver by checking its presence and if found, the
// driver will be registered.
func (d *driver) Init() (bool, error) {
	if !Present() {
		return false, errors.New("board Orange Pi not detected")
	}

	model := distro.DTModel()
	if model == "<unknown>" {
		return true, fmt.Errorf("orangepi: failed to obtain model")
	}

	err := registerHeaders(model)
	return true, err
}

// init register the driver.
func init() {
	if isArm {
		driverreg.MustRegister(&drv)
	}
}

var drv driver
