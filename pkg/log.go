package pkg

import "log"

func (m *MongoDB) Log(message string) {
	if m.SettingLog {
		log.Println(message)
	}
}

func (m *MongoDB) PanicLog(messageErr error) {
	if m.Panic {
		panic(messageErr.Error())
	}

	if m.SettingLog {
		log.Println(messageErr.Error())
	}
}
