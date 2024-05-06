@echo off
goto :deploy

:deploy
    echo Building Docker container...
    echo.
    cmd /c "docker build -t portfolio:alpha ."
    cmd /c "docker tag portfolio:alpha gcr.io/rmulnick-web/portfolio:alpha"
    cmd /c "docker push gcr.io/rmulnick-web/portfolio:alpha"

    echo.
    echo Docker build complete. Update Cloud Run instance.