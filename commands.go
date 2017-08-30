package main

var commands = map[string]string{
	"verboseupload:arduino:avr:chiwawa":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:mini:cpu=atmega168":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:mini:cpu=atmega328":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:nano:cpu=atmega328":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:nano:cpu=atmega168":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:fio":                            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:esplora":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:leonardo":                       `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:uno":                            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:circuitplay32u4cat":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:mega:cpu=atmega1280":            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega1280 -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:mega:cpu=atmega2560":            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega2560 -cwiring -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:one":                            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:yunmini":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:ethernet":                       `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:samd:arduino_zero_edbg":             `"{runtime.tools.openocd.path}/bin/openocd" -d2 -s "{runtime.tools.openocd.path}/share/openocd/scripts/" -f "{build.path}/arduino_zero.cfg" -c "telnet_port disabled; program {build.path}/{build.project_name}.bin verify reset 0x00002000; shutdown"`,
	"verboseupload:arduino:samd:mkr1000":                       `"{runtime.tools.bossac-1.7.0.path}/bossac" -i -d --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,
	"verboseupload:arduino:samd:arduino_zero_native":           `"{runtime.tools.bossac-1.7.0.path}/bossac" -i -d --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,
	"verboseupload:arduino:avr:LilyPadUSB":                     `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:pro:cpu=8MHzatmega168":          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:pro:cpu=16MHzatmega168":         `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:pro:cpu=16MHzatmega328":         `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:pro:cpu=8MHzatmega328":          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:samd:mkrzero":                       `"{runtime.tools.bossac-1.7.0.path}/bossac" -i -d --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,
	"verboseupload:arduino:avr:lilypad:cpu=atmega328":          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:lilypad:cpu=atmega168":          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:robotControl":                   `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:samd:adafruit_circuitplayground_m0": `"{runtime.tools.bossac-1.7.0.path}/bossac" -i -d --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,
	"verboseupload:arduino:avr:atmegang:cpu=atmega168":         `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:atmegang:cpu=atmega8":           `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega8 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:yun":                            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:bt:cpu=atmega328":               `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:bt:cpu=atmega168":               `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:diecimila:cpu=atmega168":        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:diecimila:cpu=atmega328":        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:gemma":                          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -pattiny85 -c{upload.protocol} -P{serial.port} -b{upload.speed} -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:unowifi":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega328p -carduino -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:nrf52:primo":                        `"{runtime.tools.openocd-0.10.0-arduino1-static.path}/bin/openocd" -s "{runtime.tools.openocd-0.10.0-arduino1-static.path}/share/openocd/scripts/" -f "{runtime.platform.path}/arduino_primo.cfg" -c "program {{{build.path}/{build.project_name}-merged.hex}} verify reset exit"`,
	"verboseupload:arduino:avr:robotMotor":                     `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:leonardoeth":                    `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:megaADK":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega2560 -cwiring -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:avr:micro":                          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -v -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"verboseupload:arduino:samd:mkrfox1200":                    `"{runtime.tools.bossac-1.7.0.path}/bossac" -i -d --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,

	"quietupload:arduino:avr:chiwawa":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:mini:cpu=atmega168":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:mini:cpu=atmega328":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:nano:cpu=atmega328":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:nano:cpu=atmega168":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:fio":                            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:esplora":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:leonardo":                       `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:uno":                            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:circuitplay32u4cat":             `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:mega:cpu=atmega1280":            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega1280 -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:mega:cpu=atmega2560":            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega2560 -cwiring -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:one":                            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:yunmini":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:ethernet":                       `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:samd:arduino_zero_edbg":             `"{runtime.tools.openocd.path}/bin/openocd" -d0 -s "{runtime.tools.openocd.path}/share/openocd/scripts/" -f "{build.path}/arduino_zero.cfg" -c "telnet_port disabled; program {build.path}/{build.project_name}.bin verify reset 0x00002000; shutdown"`,
	"quietupload:arduino:samd:mkr1000":                       `"{runtime.tools.bossac-1.7.0.path}/bossac" --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,
	"quietupload:arduino:samd:arduino_zero_native":           `"{runtime.tools.bossac-1.7.0.path}/bossac" --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,
	"quietupload:arduino:avr:LilyPadUSB":                     `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:pro:cpu=8MHzatmega168":          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:pro:cpu=16MHzatmega168":         `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:pro:cpu=16MHzatmega328":         `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:pro:cpu=8MHzatmega328":          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:samd:mkrzero":                       `"{runtime.tools.bossac-1.7.0.path}/bossac" --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,
	"quietupload:arduino:avr:lilypad:cpu=atmega328":          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:lilypad:cpu=atmega168":          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:robotControl":                   `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:samd:adafruit_circuitplayground_m0": `"{runtime.tools.bossac-1.7.0.path}/bossac" --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,
	"quietupload:arduino:avr:atmegang:cpu=atmega168":         `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:atmegang:cpu=atmega8":           `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega8 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:yun":                            `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:bt:cpu=atmega328":               `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:bt:cpu=atmega168":               `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:diecimila:cpu=atmega168":        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega168 -carduino -P{serial.port} -b19200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:diecimila:cpu=atmega328":        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:gemma":                          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -pattiny85 -c{upload.protocol} -P{serial.port} -b{upload.speed} -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:unowifi":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega328p -carduino -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:nrf52:primo":                        `"{runtime.tools.openocd-0.10.0-arduino1-static.path}/bin/openocd" -s "{runtime.tools.openocd-0.10.0-arduino1-static.path}/share/openocd/scripts/" -f "{runtime.platform.path}/arduino_primo.cfg" -c "program {{{build.path}/{build.project_name}-merged.hex}} verify reset exit"`,
	"quietupload:arduino:avr:robotMotor":                     `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:leonardoeth":                    `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:megaADK":                        `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega2560 -cwiring -P{serial.port} -b115200 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:avr:micro":                          `"{runtime.tools.avrdude.path}/bin/avrdude" "-C{runtime.tools.avrdude.path}/etc/avrdude.conf" -q -q -patmega32u4 -cavr109 -P{serial.port} -b57600 -D "-Uflash:w:{build.path}/{build.project_name}.hex:i"`,
	"quietupload:arduino:samd:mkrfox1200":                    `"{runtime.tools.bossac-1.7.0.path}/bossac" --port={serial.port.file} -U true -i -e -w -v "{build.path}/{build.project_name}.bin" -R`,

	"install:arduino-connector": "source /tmp/env.sh && wget -O - https://downloads.arduino.cc/tools/arduino-connector.sh | bash"}