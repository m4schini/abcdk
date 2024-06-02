package log

type PrintFunc func(v ...interface{})
type PrintfFunc func(format string, v ...interface{})

type MqttAdapter struct {
	PrintF  PrintFunc
	PrintfF PrintfFunc
}

func (m *MqttAdapter) Println(v ...interface{}) {
	m.PrintF(v...)
}

func (m *MqttAdapter) Printf(format string, v ...interface{}) {
	m.PrintfF(format, v...)
}
