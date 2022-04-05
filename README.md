# gtimer
### Simple timer gui app built with fyne

## Layout (UI)

There are two parts making up the visible UI:

- timer display: created with a GridLayout using a fyne.Container Object 
- control buttons display: made of 2 buttons placed in a GridLayout using a fyne.Container object

## Data binding (UI update)

A Timer `struct` with respective hour, minute, and second binding.String Objects binded to the timer display. This allows any update to the respective binded objects to update the timer display.

## Control Buttons

There are two buttons - start and stop. These buttons control the timer display by updating the `state` property of the Timer `struct`.  The start button starts out by setting the state to `"start"` and then sending off a update goroutine `updateTimer` which runs continuously in the background. 

Once the timer is started, the start button is disabled to prevent creating multiple goroutines to handle the same task. The stop button stops the timer by setting the state to `"stop"` and causing the background task to be terminated. Once the timer is stopped, the start button is enabled again allowing the timer to be started again.

## Screen Shot 

![gtimer](/gtimer/screenshot.png)