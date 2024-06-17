package bds44

import (
	"errors"
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds44/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// MeteorologicalRoutineAirReportV1 is a message at the format BDS 4,4
//
// Specified in Doc 9871 / Table E-2-68
type MeteorologicalRoutineAirReportV1 struct {
	WindSpeedStatus             bool
	WindSpeed                   uint32
	WindDirectionStatus         bool
	WindDirection               float32
	StaticAirTemperatureStatus  bool
	StaticAirTemperature        float32
	AverageStaticPressureStatus bool
	AverageStaticPressure       uint32
	TurbulenceFlag              fields.TurbulenceFlag
	HumidityStatus              bool
	Humidity                    float32
}

// GetRegister returns the Register the message
func (message MeteorologicalRoutineAirReportV1) GetRegister() register.Register {
	return register.BDS44
}

func (message MeteorologicalRoutineAirReportV1) GetSource() fields.Source {
	return fields.SourceInvalid
}

// ToString returns a basic, but readable, representation of the message
func (message MeteorologicalRoutineAirReportV1) ToString() string {
	return fmt.Sprintf(""+
		"Message:                            %v\n"+
		"Wind Speed Status:                  %v\n"+
		"Wind Speed (knot):                  %v\n"+
		"Wind Direction Status:              %v\n"+
		"Wind Direction (degrees):           %v\n"+
		"Static Air Temperature Status:      %v\n"+
		"Static Air Temperature (degrees C): %v\n"+
		"Average Static Pressure Status:     %v\n"+
		"Average Static Pressure (hPa):      %v\n"+
		"Turbulence Flag:                    %v\n"+
		"Humidity Status:                    %v\n"+
		"Humidity (%%):                       %v",
		message.GetRegister().ToString(),
		message.WindSpeedStatus,
		message.WindSpeed,
		message.WindDirectionStatus,
		message.WindDirection,
		message.StaticAirTemperatureStatus,
		message.StaticAirTemperature,
		message.AverageStaticPressureStatus,
		message.AverageStaticPressure,
		message.TurbulenceFlag.ToString(),
		message.HumidityStatus,
		message.Humidity)
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message MeteorologicalRoutineAirReportV1) CheckCoherency() error {
	// If no data available, it is probably not coherent
	if !message.WindSpeedStatus && !message.WindDirectionStatus && !message.StaticAirTemperatureStatus && !message.AverageStaticPressureStatus && !message.HumidityStatus {
		return errors.New("the message does not convey any information")
	}

	if !message.WindSpeedStatus && message.WindSpeed != 0 {
		return errors.New("the wind speed status is set to false, but a wind speed value is given")
	}

	if !message.WindDirectionStatus && message.WindDirection != 0 {
		return errors.New("the wind direction status is set to false, but a wind direction value is given")
	}

	if !message.StaticAirTemperatureStatus && message.StaticAirTemperature != 0 {
		return errors.New("the static air temperature status is set to false, but a static air temperature value is given")
	}

	if !message.AverageStaticPressureStatus && message.AverageStaticPressure != 0 {
		return errors.New("the average static pressure status is set to false, but a average static pressure value is given")
	}

	if !message.HumidityStatus && message.Humidity != 0 {
		return errors.New("the humidity status is set to false, but a humidity value is given")
	}

	if message.WindSpeed > 250 {
		return errors.New("the wind speed is too high (above 250 knots)")
	}

	if (message.StaticAirTemperature < -80) || (message.StaticAirTemperature > 60) {
		return errors.New("the static air temperature is to high or to low (-80 <= temp <= 60)")
	}

	return nil
}

// ReadMeteorologicalRoutineAirReportV1 reads a message as a MeteorologicalRoutineAirReportV1
func ReadMeteorologicalRoutineAirReportV1(data []byte) (*MeteorologicalRoutineAirReportV1, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B MeteorologicalRoutineAirReport message must be 7 bytes long")
	}

	if data[0]&0xF0 != 0 {
		return nil, errors.New("the bits 1 to 4 must be zero")
	}

	windSpeedStatus, windSpeed := fields.ReadWindSpeed(data)
	windDirectionStatus, windDirection := fields.ReadWindDirectionV1AndV2(data)
	staticAirTemperatureStatus, staticAirTemperature := fields.ReadStaticAirTemperatureV1AndV2(data)
	averageStaticPressureStatus, averageStaticPressure := fields.ReadAverageStaticPressureV1AndV2(data)
	turbulenceFlag := fields.ReadTurbulenceFlag(data)
	humidityStatus, humidity := fields.ReadHumidityV1AndV2(data)

	return &MeteorologicalRoutineAirReportV1{
		WindSpeedStatus:             windSpeedStatus,
		WindSpeed:                   windSpeed,
		WindDirectionStatus:         windDirectionStatus,
		WindDirection:               windDirection,
		StaticAirTemperatureStatus:  staticAirTemperatureStatus,
		StaticAirTemperature:        staticAirTemperature,
		AverageStaticPressureStatus: averageStaticPressureStatus,
		AverageStaticPressure:       averageStaticPressure,
		TurbulenceFlag:              turbulenceFlag,
		HumidityStatus:              humidityStatus,
		Humidity:                    humidity,
	}, nil
}
