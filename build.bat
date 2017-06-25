@echo off
setlocal EnableExtensions EnableDelayedExpansion

goto skip

go get
set CGO_ENABLED=0

pushd plugins
for /D %%d in (*) do (
    FOR /f "delims=" %%a in ('type %%d\modtime') do (
	set MODTIME=%%a
	set MODTIME=!modtime:~0,16!
    )
    
    if "!MODTIME!" == "%%~td" (
       echo %%d skipped
       
    ) else (
       echo.
       rem echo %%~td > %%d\modtime
       rem echo %%d %%~td !modtime!
       
       pushd ..
       ibt -begin ibt\%%d.ibt
       popd
       
       copy /v %%d\*.json *.json
       go build -o %%d.exe %%d\%%d.go
       set LastError=%ERRORLEVEL1%
       IF %ERRORLEVEL% EQU 0 (
       	   echo %%~td > %%d\modtime
       	   echo %%d %%~td !modtime!
       )

       pushd ..
       ibt -end ibt\%%d.ibt %LastError1%
       popd
    )
)
popd

FOR /f "delims=" %%a in ('type modtime') do (
    set MODTIME=%%a
    set MODTIME=!modtime:~0,16!
)
for /d %%f in (.) do (
    if "!MODTIME!" == "%%~tf" (
       echo Marcy skipped
       goto exit
    )
)
:skip

ibt -begin ibt\marcy.ibt

set /p build=<build
set /a build=build
set /a build=build+1
echo %build% > build

set /p majorver=<major.ver
set /a majorver=majorver
echo %majorver% > major.ver

set /p minorver=<minor.ver
set /a minorver=minorver
echo %minorver% > minor.ver

echo V%majorver%.%minorver% b%build%
go build -ldflags "-s -w -X main.Version=%majorver%.%minorver% -X main.Build=%build%"
set LastError=%ERRORLEVEL%
IF %ERRORLEVEL% EQU 0 (
   for /d %%f in (.) do (
       rem echo %%~tf > modtime
   )
)

ibt -end ibt\marcy.ibt %LastError%

:exit

