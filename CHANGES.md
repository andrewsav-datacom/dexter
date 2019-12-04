# Changes to the original project

A change is made comparing to upstream as of the time of writing:

* Fixed a crash when dexter is run on Windows from an executable compiled with [ps2exe](https://gallery.technet.microsoft.com/scriptcenter/PS2EXE-Convert-PowerShell-9e4e07f1).


## Building for Windows

It's best done on a linux machine. Clone the repo to a directory and then run a `make` command from it, e.g.:

```bash
DEFAULT_PROVIDER=google CLIENT_ID=clientid.apps.googleusercontent.com CLIENT_SECRET=clientsecret OS=windows make
```

The executable will be in `build` sub-directory.
