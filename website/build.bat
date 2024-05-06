@echo off
goto :init

:header
    echo %__NAME% v%__VERSION%
    echo Simple batch file to build my Portfolio Website project.
    echo Optionally, once the build completes you can either
    echo run or deploy the project.
    echo.
    goto :eof

:usage
    echo USAGE:
    echo.
    echo.  -h, --help           shows this help
    echo.  -v, --version        shows the version
    echo.  -r, --run            build and run locally
    echo.  -d, --deploy         build and deploy to gcloud
    echo.  -b, --build          only build
    goto :eof

:version
    if "%~1"=="full" call :header & goto :eof
    echo v%__VERSION%
    goto :eof

:init
    set "__NAME=Portfolio"
    set "__VERSION=0.0.1"
    set "__YEAR=2024"
    set "__BUILD_CONFIG=production"
    set "__PROJECT_ID=rmulnick-web"

    set "__BAT_FILE=%~0"
    set "__BAT_PATH=%~dp0"
    set "__BAT_NAME=%~nx0"

    set "OptHelp="
    set "OptVersion="
    set "OptVerbose="

:parse
    if /i "%~1"=="-h"         call :header & call :usage "%~2" & goto :end
    if /i "%~1"=="--help"     call :header & call :usage "%~2" & goto :end

    if /i "%~1"=="-v"         call :version & goto :end
    if /i "%~1"=="--version"  call :version & goto :end

    if /i "%~1"=="-r"         call :build & call :run & goto :end
    if /i "%~1"=="--run"      call :build & call :run & goto :end

    if /i "%~1"=="-d"         call :build & call :deploy & goto :end
    if /i "%~1"=="--deploy"   call :build & call :deploy & goto :end

    if /i "%~1"=="-b"         call :build & goto :end
    if /i "%~1"=="--build"    call :build & goto :end

    shift
    goto :parse

:build
    echo Building Portfolio project %__VERSION%...
    cmd /c "ng build -c %__BUILD_CONFIG%"
    echo Project built to ./dist directory
    goto :eof

:run
    cmd /c "node server.js"
    goto :eof

:deploy
    echo Deploying build to Google Cloud: App Engine
    echo.
    cmd /c "gcloud config set project %__PROJECT_ID%"
    cmd /c "gcloud app deploy"

:main
    echo Initializing Portfolio build.
    echo.

:end
    call :cleanup
    exit /B

:cleanup
    REM The cleanup function is only really necessary if you
    REM are _not_ using SETLOCAL.
    set "__NAME="
    set "__VERSION="
    set "__YEAR="
    set "__BUILD_CONFIG="
    set "__PROJECT_ID="

    set "__BAT_FILE="
    set "__BAT_PATH="
    set "__BAT_NAME="

    set "OptHelp="
    set "OptVersion="
    set "OptVerbose="

    goto :eof