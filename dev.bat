@echo off

set target_name=goctl-php

for /f "delims=" %%t in ('where goctl') do set goctl_path=%%t
for %%F in (%goctl_path%) do set goctl_dir=%%~dpF

go build -o %goctl_dir%/%target_name%.exe .
echo install %target_name% to: %goctl_dir%
