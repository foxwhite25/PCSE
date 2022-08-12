Potion Craft Save Editor
========================

## Introduction
This is a Save Editor for Game [Potion Craft](https://www.potioncraft.com/)

This Go Application will work on Windows, Linux and maybe MAC.
Go to the release page and download the executable you need, or build the project by yourself.

I want to add I am not responsible if you break your save file, I try to mantain this editor in a way, 
in which it can't break save files (or not by mistake) but it is still possible that this editor produces
a bugged save file (if you for example add an item in which the item doesn't exist).

That is why you always should **backup your save files**.

Since Potion Craft is currently in early access, 
no promises that this will work forever, but I will try to update when the game update.

## Currently Editable
* Gold
* Karma
* Popularity
* Statistics
* Add, Remove, Edit Items

## Installation

If you have go installed, you can directly clone the master branch and build it.

If not you can also download a executable for your platform in the [release](https://github.com/foxwhite25/PCSE/releases) page.

## Save File Locations

Linux / Ubuntu:

* `Placeholder I dont know`

MacOS:

* `Placeholder I dont know`

Windows:

* `C:\Users\%username%\AppData\LocalLow\niceplay games\Potion Craft\Saves`

## Additional Thanks go to

* Potion Craft Discord Server for data-mining the encryption key.

## The application is not working?

Send an issue with screenshot of the console output (or the copied text).
If you are using the compiled version and a black window appears and then disappears it means the application crashes because of some error.

To view the error code to be able to send an issue:
In the folder where you have the .exe file, `Shift + Rightclick` in a free space and in the context menu there should be an option like `Open Command Prompt here` or `Open Powershell here`, 
click that, begin writing `potion` and then press tab to autocomplete and enter to execute - now you should start the application using that console window. This time the window won't close after execution, meaning you have time to make a screenshot of the error.

## Notice

This repository contains content which I do not own.
Notably all the image files in the `/res` folder. These are by [Niceplay Games](http://niceplay-games.com/).

If you find any bugs / mistakes, feel free to open issues, or if you know how to fix it yourself, feel free to create a pull request.