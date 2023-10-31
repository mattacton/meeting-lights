# meeting-lights
A small program to handle key press events and map them to functions.

This application was designed and built around the tools I had available. 
It currently runs on a Raspberry Pi 3 model B as the main foreground application and takes input
from a Stack Overflow macropad https://drop.com/buy/stack-overflow-the-key-macropad. When the correct
button(s) or button combinations are pressed, a function is run to change the color of a
Philips Hue smart light which sits outside my office. 

This is a handy way to signal those in my household the state I'm currently "in". Think of it as
an "On Air" light with a few more states. e.g. Free, In a meeting, Focusing, etc. 

## Running
The application requires the following environment variables for operation:
* LIGHT_HOST - The host of the Philips Hue hub
* LIGHT_SECRET - The secret key used for AuthN / AuthZ
* LIGHT_IDS - A comma separated list of light Ids to act upon

The project includes binaries for MacOS and Linux ARM architectures.

## Notes
On my Raspberry Pi, it can take the WiFi module a good 30 seconds to start up, so I have added a delay in
the startup script of my Pi so that the meeting-lights binary executes after WiFi has started
