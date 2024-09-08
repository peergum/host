// Copyright 2022 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// This file contains pin mapping information that is specific to the Allwinner
// H616 model.

package allwinner

import (
	"strings"

	"periph.io/x/conn/v3/pin"
	"periph.io/x/host/v3/sysfs"
)

// mappingH616 describes the mapping of the H616 processor gpios to their
// alternate functions.
//
// It omits the in & out functions which are available on all gpio.
//
// The mapping comes from the datasheet page 55:
// https://linux-sunxi.org/images/b/b9/H616_Datasheet_V1.0_cleaned.pdf
//
//   - The datasheet uses TWI instead of I2C but it is renamed here for
//     consistency.
//   - RGMII means Reduced gigabit media-independent interface.
//   - SDC means SDCard?
//   - NAND connects to a NAND flash controller.
//   - CSI and CCI are for video capture.
var mappingH616 = map[string][5]pin.Func{
	"PC0":  {"NAND_WE", "SDC2_DS", "SPI0_CLK", "", "PC_EINT0"},
	"PC1":  {"NAND_ALE", "SDC2_RST", "", "", "PC_EINT1"},
	"PC2":  {"NAND_CLE", "", "SPI0_MOSI", "", "PC_EINT2"},
	"PC3":  {"NAND_CE1", "", "SPI0_CS0", "BOOT_SEL1", "PC_EINT3"},
	"PC4":  {"NAND_CE0", "", "SPI0_MISO", "BOOT_SEL2", "PC_EINT4"},
	"PC5":  {"NAND_RE", "SDC2_CLK", "", "BOOT_SEL3", "PC_EINT5"},
	"PC6":  {"NAND_RB0", "SDC2_CMD", "", "BOOT_SEL4", "PC_EINT6"},
	"PC7":  {"NAND_RB1", "", "SPI0_CS1", "", "PC_EINT7"},
	"PC8":  {"NAND_DQ7", "SDC2_D3", "", "", "PC_EINT8"},
	"PC9":  {"NAND_DQ6", "SDC2_D4", "", "", "PC_EINT9"},
	"PC10": {"NAND_DQ5", "SDC2_D0", "", "", "PC_EINT10"},
	"PC11": {"NAND_DQ4", "SDC2_D5", "", "", "PC_EINT11"},
	"PC12": {"NAND_DQS", "", "", "", "PC_EINT12"},
	"PC13": {"NAND_DQ3", "SDC2_D1", "", "", "PC_EINT13"},
	"PC14": {"NAND_DQ2", "SDC2_D6", "", "", "PC_EINT14"},
	"PC15": {"NAND_DQ1", "SDC2_D2", "SPI0_WP", "", "PC_EINT15"},
	"PC16": {"NAND_DQ0", "SDC2_D7", "SPI0_HOLD", "", "PC_EINT16"},

	"PF0": {"SDC0_D1", "JTAG_MS", "", "", "PF_EINT0"},
	"PF1": {"SDC0_D0", "JTAG_DI", "", "", "PF_EINT1"},
	"PF2": {"SDC0_CLK", "UART0_TX", "", "", "PF_EINT2"},
	"PF3": {"SDC0_CMD", "JTAG_DO", "", "", "PF_EINT3"},
	"PF4": {"SDC0_D3", "UART0_RX", "", "", "PF_EINT4"},
	"PF5": {"SDC0_D2", "JTAG_CK", "", "", "PF_EINT5"},
	"PF6": {"", "", "", "", "PF_EINT6"},

	"PG0":  {"SDC1_CLK", "", "", "", "PG_EINT0"},
	"PG1":  {"SDC1_CMD", "", "", "", "PG_EINT1"},
	"PG2":  {"SDC1_D0", "", "", "", "PG_EINT2"},
	"PG3":  {"SDC1_D1", "", "", "", "PG_EINT3"},
	"PG4":  {"SDC1_D2", "", "", "", "PG_EINT4"},
	"PG5":  {"SDC1_D3", "", "", "", "PG_EINT5"},
	"PG6":  {"UART1_TX", "", "JTAG_MS", "", "PG_EINT6"},
	"PG7":  {"UART1_RX", "", "JTAG_CK", "", "PG_EINT7"},
	"PG8":  {"UART1_RTS", "", "JTAG_DO", "", "PG_EINT8"},
	"PG9":  {"UART1_CTS", "", "JTAG_DI", "", "PG_EINT9"},
	"PG10": {"H_I2S2_MCLK", "X32KFOUT", "", "", "PG_EINT10"},
	"PG11": {"H_I2S2_BCLK", "", "BIST_RESULT0", "", "PG_EINT11"},
	"PG12": {"H_I2S2_LRCK", "", "BIST_RESULT1", "", "PG_EINT12"},
	"PG13": {"H_I2S2_DOUT0", "H_I2S2_DIN1", "BIST_RESULT2", "", "PG_EINT13"},
	"PG14": {"H_I2S2_DIN0", "H_I2S2_DOUT1", "BIST_RESULT3", "", "PG_EINT14"},
	"PG15": {"UART2_TX", "", "", "TWI4_SCK", "PG_EINT15"},
	"PG16": {"UART2_RX", "", "", "TWI4_SDA", "PG_EINT16"},
	"PG17": {"UART2_RTS", "", "", "TWI3_SCK", "PG_EINT17"},
	"PG18": {"UART2_CTS", "", "", "TWI3_SDA", "PG_EINT18"},
	"PG19": {"", "", "PWM1", "", "PG_EINT19"},

	"PH0":  {"UART0_TX", "", "PWM3", "TWI1_SCK", "PH_EINT0"},
	"PH1":  {"UART0_RX", "", "PWM3", "TWI1_SDA", "PH_EINT1"},
	"PH2":  {"UART5_TX", "OWA_MCLK", "PWM2", "TWI2_SCK", "PH_EINT2"},
	"PH3":  {"UART5_RX", "", "PWM1", "TWI2_SDA", "PH_EINT3"},
	"PH4":  {"", "OWA_OUT", "", "TWI3_SCK", "PH_EINT4"},
	"PH5":  {"UART2_TX", "H_I2S3_MCLK", "SPI1_CS0", "TWI3_SDA", "PH_EINT5"},
	"PH6":  {"UART2_RX", "H_I2S3_BCLK", "SPI1_CLK", "TWI4_SCK", "PH_EINT6"},
	"PH7":  {"UART2_RTS", "H_I2S3_LRCK", "SPI1_MOSI", "TWI4_SDA", "PH_EINT7"},
	"PH8":  {"UART2_CTS", "H_I2S3_DOUT0", "SPI1_MISO", "H_I2S3_DIN1", "PH_EINT8"},
	"PH9":  {"", "H_I2S3_DIN0", "SPI1_CS1", "H_I2S3_DOUT1", "PH_EINT9"},
	"PH10": {"", "IR_RX", "TCON_TRIG1", "", "PH_EINT10"},

	"PI0":  {"RGMII_RXD3/RMII_NULL", "DMIC_CLK", "H_I2S0_MCLK", "HDMI_SCL", "PI_EINT0"},
	"PI1":  {"RGMII_RXD2/RMII_NULL", "DMIC_DATA0", "H_I2S0_BCLK", "HDMI_SDA", "PI_EINT1"},
	"PI2":  {"RGMII_RXD1/RMII_RXD1", "DMIC_DATA1", "H_I2S0_LRCK", "HDMI_CEC", "PI_EINT2"},
	"PI3":  {"RGMII_RXD0/RMII_RXD0", "DMIC_DATA2", "H_I2S0_DOUT0", "H_I2S0_DIN1", "PI_EINT3"},
	"PI4":  {"RGMII_RXCK/RMII_NULL", "DMIC_DATA3", "H_I2S0_DIN0", "H_I2S0_DOUT1", "PI_EINT4"},
	"PI5":  {"RGMII_RXCTL/RMII_CRS_DV", "UART2_TX", "TS0_CLK", "TWI0_SCK", "PI_EINT5"},
	"PI6":  {"RGMII_NULL/RMII_RXER", "UART2_RX", "TS0_ERR", "TWI0_SDA", "PI_EINT6"},
	"PI7":  {"RGMII_TXD3/RMII_NULL", "UART2_RTS", "TS0_SYNC", "TWI1_SCK", "PI_EINT7"},
	"PI8":  {"RGMII_TXD2/RMII_NULL", "UART2_CTS", "TS0_DVLD", "TWI1_SDA", "PI_EINT8"},
	"PI9":  {"RGMII_TXD1/RMII_TXD1", "UART3_TX", "TS0_D0", "TWI2_SCK", "PI_EINT9"},
	"PI10": {"RGMII_TXD0/RMII_TXD0", "UART3_RX", "TS0_D1", "TWI2_SDA", "PI_EINT10"},
	"PI11": {"RGMII_TXCK/RMII_TXCK", "UART3_RTS", "TS0_D2", "PWM1", "PI_EINT11"},
	"PI12": {"RGMII_TXCTL/RMII_TXEN", "UART3_CTS", "TS0_D3", "PWM2", "PI_EINT12"},
	"PI13": {"RGMII_CLKIN/RMII_NULL", "UART4_TX", "TS0_D4", "PWM3", "PI_EINT13"},
	"PI14": {"MDC", "UART4_RX", "TS0_D5", "PWM4", "PI_EINT14"},
	"PI15": {"MDIO", "UART4_RTS", "TS0_D6", "CLK_FANOUT0", "PI_EINT15"},
	"PI16": {"EPHY_25M", "UART4_CTS", "TS0_D7", "CLK_FANOUT1", "PI_EINT16"},

	"PL0": {"", "S_TWI0_SCK", "", "", ""},
	"PL1": {"", "S_TWI0_SDA", "", "", ""},
}

// mapH5Pins uses mappingH5 to actually set the altFunc fields of all gpio
// and mark them as available.
//
// It is called by the generic allwinner processor code if an H5 is detected.
func mapH616Pins() error {
	for name, altFuncs := range mappingH5 {
		pin := cpupins[name]
		pin.altFunc = altFuncs
		pin.available = true
		if strings.Contains(string(altFuncs[4]), "_EINT") {
			pin.supportEdge = true
		}

		// Initializes the sysfs corresponding pin right away.
		pin.sysfsPin = sysfs.Pins[pin.Number()]
	}
	return nil
}
