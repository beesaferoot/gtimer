package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Timer struct {
	hh binding.String
	mm binding.String
	ss binding.String
	state string
	start_btn *widget.Button
}

func (t *Timer) update() {
	if t.state != "stop" {
		ssStr, _ := t.ss.Get()
		ss, _ := strconv.Atoi(ssStr)
		if ss + 1 < 60 {
			t.ss.Set(fmt.Sprintf("%02d", ss + 1))
		}else {
			t.ss.Set(fmt.Sprintf("%02d", 0))
			mmStr, _ := t.mm.Get()
			mm, _ := strconv.Atoi(mmStr)
			if mm + 1 < 60 {
				t.mm.Set(fmt.Sprintf("%02d", mm + 1))
			}else{
				t.mm.Set(fmt.Sprintf("%02d", 0))

				hhStr, _ := t.hh.Get()
				hh, _ := strconv.Atoi(hhStr)
				if hh + 1 < 24 {
					t.hh.Set(fmt.Sprintf("%02d", hh + 1))
				}else{
					hh = 0
					t.hh.Set(fmt.Sprintf("%02d", hh + 1))
				}
			}

		}
	}else {
		t.hh.Set("00")
		t.mm.Set("00")
		t.ss.Set("00")
	}
}

func NewTimer () * Timer {
	var timer *Timer = &Timer{}
	timer.hh = binding.NewString()
	timer.mm = binding.NewString()
	timer.ss = binding.NewString()
	timer.hh.Set("00")
	timer.mm.Set("00")
	timer.ss.Set("00")
	return timer
}


func NewTimerDisplay(timer *Timer) *fyne.Container {
	// container to display timer 
	hText := widget.NewLabel("")
	mText := widget.NewLabel("")
	sText := widget.NewLabel("")
	// bind label data 
	hText.Bind(timer.hh)
	mText.Bind(timer.mm)
	sText.Bind(timer.ss)

	timeContainer := container.New(layout.NewGridLayout(5), 
							hText, 
							widget.NewLabel(":"), 
							mText, 
							widget.NewLabel(":"), 
							sText,
						)
	return timeContainer
}

func NewTimerConTrolDisplay(timer *Timer) *fyne.Container {

	startBtn := widget.NewButton("START", func() {
		timer.state = "start"
		timer.start_btn.Disable()
		go updateTimer(timer)
	})
	timer.start_btn = startBtn
	stopBtn := widget.NewButton("STOP", func() {
		timer.state = "stop"
		timer.start_btn.Enable()
	})
	btnContainer := container.NewGridWithColumns(2, startBtn, stopBtn)
	return btnContainer
}

func updateTimer (timer *Timer) {

	for range time.Tick(time.Second) {
		// update timer
		timer.update()
		if timer.state == "stop"{
			break
		}
	}
	
	
}

func main() {
	a := app.New()
	w := a.NewWindow("Timer")
	timer := NewTimer()
	timerDisplay := NewTimerDisplay(timer)
	btnDisplay := NewTimerConTrolDisplay(timer)
	w.SetContent(container.NewGridWithRows(2, timerDisplay, btnDisplay))
	w.SetFixedSize(true)
	w.Resize(fyne.Size{Width: 200, Height: 200})

	w.ShowAndRun()

}