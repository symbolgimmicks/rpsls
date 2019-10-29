@echo off
for %%i in (./features/*) do (
	echo [%%i TEST]>"%%i.log"
	godog --no-colors "./features/%%i">>"%%i.log"
)