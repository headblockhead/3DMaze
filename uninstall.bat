@echo off
echo Uninstalling, This might take a few seconds...
del C:\3DMAZE\ARCADE~1.exe 
del C:\3DMAZE\textures.png  
del C:\3DMAZE\Manual.txt  
del C:\3DMAZE\UNINST~1.exe
del "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3D Maze\Launch.lnk"
del "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3D Maze\Manual.lnk"
del "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3D Maze\Uninstall.lnk"
rmdir "C:\3DMAZE\"
rmdir "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3D Maze\"
