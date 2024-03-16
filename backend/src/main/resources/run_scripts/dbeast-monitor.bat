@echo off

set HOME_DIR=%~dp0
cd %HOME_DIR%

echo ###################################################################
echo Dbeast toolkit home folder: %HOME_DIR%
echo ###################################################################
echo #

where java >nul 2>nul
if %errorlevel%==1 (
    @echo Java not found in path.
    exit
)

java -Dlog4j2.configurationFile=%HOME_DIR%\config\log4j2.xml -cp %HOME_DIR%\lib\*;%HOME_DIR%\bin\* co.dbeast.dbeast_toolkit.runner.DbeastToolkit -c %HOME_DIR%

PAUSE