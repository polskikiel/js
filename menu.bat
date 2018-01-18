@ECHO OFF
REM.-- Prepare the Command Processor
title Przyklej plakaty
@SETLOCAL ENABLEEXTENSIONS
@SETLOCAL ENABLEDELAYEDEXPANSION

:menuLOOP
echo ------------------------------
echo 	     Menu             
echo ------------------------------
echo.
for /f "tokens=1,2,* delims=_ " %%A in ('"findstr /b /c:":menu_" "%~f0""') do echo.  %%B  %%C
set choice=
echo ------------------------------
echo.&set /p choice= Wybor:  ||GOTO:EOF
echo.&call:menu_%choice%
GOTO:menuLOOP

:menu_1  Przyklej plakaty

set wejscie=".\in"
set wyjscie=".\out"
if not exist "%wejscie%" (
	echo Podana sciezka do plikow z liczbami nie istnieje!
	GOTO:EOF
)
if exist "%wejscie%" (
	::Sprawdzenie poprawnosci sciezki zrodlowej-----
	for /f "delims=" %%a in ("%wejscie%") do (
		set variable=%%a
		if "!variable:~-1!"=="\" (
			set variable=!variable:~0,-1!
		)
    		set wejscie=!variable!
	)
	::----------------------------------------------
	if not exist "%wyjscie%" (
		echo Podana sciezka do zapisania plikow nie istnieje!
		GOTO:EOF
	)
	if exist "%wyjscie%" (
		::Sprawdzenie poprawnosci sciezki docelowej-----
		for /f "delims=" %%a in ("%wyjscie%") do (
			set variable=%%a
			if "!variable:~-1!"=="\" (
				set variable=!variable:~0,-1!
			)
    			set wyjscie=!variable!
		)
		::----------------------------------------------
		set curpath=%cd%
		pushd "!wejscie!"
		dir /a:-d /b *.txt>lista_plikow.tmp
		popd
		FOR /f "tokens=* delims=," %%a in ('type "!wejscie!"\lista_plikow.tmp') do (
			"Projekt.exe" "%wejscie%" "%wyjscie%" "%%a"
		)
		::---------------------------------------------------------------------------
		del "!wejscie!"\lista_plikow.tmp
	)
)
echo.
echo Wszystkie pliki zostaly przeanalizowane.
echo.
if exist !wejscie! (
	if exist !wyjscie! (
		"sudoku.py" !wejscie! !wyjscie!
		echo.
		echo Strona z wynikami zostala stworzona.
		echo.
	)
)
pause
cls
GOTO:menuLOOP


:menu_2 Backup
xcopy ".\in\*.*" ".\backup\" /d /s
xcopy ".\*.html" ".\backup\" 
xcopy ".\out\*.*" ".\backup\" /d /s
echo Backup zostal wykonany pomyslnie
pause
cls

:menu_


:menu_ENTER Wyjdz
pause