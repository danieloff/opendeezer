# opendeezer
A method to open a deezer website page in the desktop app

Steps to setup the system of opening pages with deezer app
-

## currently Windows only

- install the external application button extension to the browser

"External Application Button Offered by: yokris.dev"

https://chrome.google.com/webstore/detail/external-application-butt/bifmfjgpgndemajpeeoiopbeilbaifdo?hl=en

follow the instructions to setup the plugin

go to a deezer website page, log in to deezer, and enable the plugin for the site.

- at the stage of adding an app, you may use these example settings

```
Display Name	Open in Deezer
Executable Name	opendeezer.exe
If the application is known by the PATH environment variable, there is no need to enter its absolute path.
Arguments	[PATHNAME]

check the box "All Contexts"

save/update application near the bottom when finished
```

- for this to work the opendeezer relay needs to be installed

download or build the opendeezer.exe and *add it to the path*

a pre-built windows 64 exe is available in the releases section of github

https://github.com/danieloff/opendeezer/releases/download/0.1/opendeezer.exe

- Close and restart the browser. If everything is configured correctly the system should now work

right click on a deezer webpage with the extension installed and there should be an option to "Open in Deezer"

After using the option the deezer app should open and show the matching page inside the app, the same as the page in the browser.

Done
