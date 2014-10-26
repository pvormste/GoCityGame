@echo off
echo Installing: vormstein.eu/gocitygame/app/controller ...
go install vormstein.eu/gocitygame/app/controller

::echo Installing: vormstein.eu/gocitygame/app/models ...
::go install vormstein.eu/gocitygame/app/models

echo Installing: vormstein.eu/gocitygame/app/views ...
go install vormstein.eu/gocitygame/app/views

echo Installing: vormstein.eu/gocitygame/util ...
go install vormstein.eu/gocitygame/util

echo Ready!