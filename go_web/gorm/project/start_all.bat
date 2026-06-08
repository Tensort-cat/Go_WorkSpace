@echo off

echo Starting backend...

start cmd /k "cd /d %~dp0back && back.exe"

echo Starting frontend...

start cmd /k "cd /d %~dp0bubble_frontend && npm run serve"

echo All services started.
pause