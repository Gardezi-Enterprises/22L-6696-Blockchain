@echo off
echo ============================================
echo   Blockchain App — Startup Script
echo   22L-6696 A1
echo ============================================
echo.

:: ---- Check Go ----
where go >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Go is not installed or not in PATH.
    echo  Download from: https://golang.org/dl/
    echo  Restart this script after installing.
    pause
    exit /b 1
)

:: ---- Check Node ----
where node >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Node.js is not installed or not in PATH.
    echo  Download from: https://nodejs.org/
    pause
    exit /b 1
)

echo [OK] Go and Node.js found.
echo.

:: ---- Start Go Backend ----
echo [1/2] Starting Go backend on http://localhost:8080 ...
start "Go Blockchain API" cmd /k "cd /d %~dp0backend && go run main.go"

:: Wait a couple seconds for the server to start
timeout /t 3 /nobreak >nul

:: ---- Install React deps if needed ----
if not exist "%~dp0frontend\node_modules" (
    echo [2/2] Installing React dependencies (first run only)...
    cd /d "%~dp0frontend"
    call npm install
) else (
    echo [2/2] React dependencies already installed.
)

:: ---- Start React Frontend ----
echo [2/2] Starting React frontend on http://localhost:3000 ...
start "React Blockchain UI" cmd /k "cd /d %~dp0frontend && npm start"

echo.
echo ============================================
echo  Both servers are starting...
echo  Backend : http://localhost:8080
echo  Frontend: http://localhost:3000
echo ============================================
pause
